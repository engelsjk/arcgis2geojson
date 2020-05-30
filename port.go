package arcgis2geojson

import (
	"errors"
)

type Ring [][]float64
type Polygon []Ring

///////////////////////////////////////////////////////////////////////////////////////
// CODE PORTED FROM https://github.com/Esri/arcgis-to-geojson-utils/blob/master/index.js)

// checks if 2 x,y points are equal
func pointsEqual(a, b []float64) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// checks if the first and last points of a ring are equal and closes the ring
func closeRing(coordinates [][]float64) [][]float64 {
	if !pointsEqual(coordinates[0], coordinates[len(coordinates)-1]) {
		coordinates = append(coordinates, coordinates[0])
	}
	return coordinates
}

func reverse(ring [][]float64) [][]float64 {
	newRing := make([][]float64, len(ring))
	copy(newRing, ring)
	for i := len(newRing)/2 - 1; i >= 0; i-- {
		opp := len(newRing) - 1 - i
		newRing[i], newRing[opp] = newRing[opp], newRing[i]
	}
	return newRing
}

// determine if polygon ring coordinates are clockwise. clockwise signifies outer ring, counter-clockwise an inner ring
// or hole. this logic was found at http://stackoverflow.com/questions/1165647/how-to-determine-if-a-list-of-polygon-
// points-are-in-clockwise-order
func ringIsClockwise(ringToTest [][]float64) bool {
	total := 0.0
	rLength := len(ringToTest)
	pt1 := ringToTest[0]
	var pt2 []float64
	for i := 0; i < (rLength - 1); i++ {
		pt2 = ringToTest[i+1]
		total = total + (pt2[0]-pt1[0])*(pt2[1]+pt1[1])
		pt1 = pt2
	}
	return total >= 0
}

// ported from terraformer.js https://github.com/Esri/Terraformer/blob/master/terraformer.js#L504-L519
func vertexIntersectsVertex(a1, a2, b1, b2 []float64) bool {
	var uaT = ((b2[0] - b1[0]) * (a1[1] - b1[1])) - ((b2[1] - b1[1]) * (a1[0] - b1[0]))
	var ubT = ((a2[0] - a1[0]) * (a1[1] - b1[1])) - ((a2[1] - a1[1]) * (a1[0] - b1[0]))
	var uB = ((b2[1] - b1[1]) * (a2[0] - a1[0])) - ((b2[0] - b1[0]) * (a2[1] - a1[1]))

	if uB != 0 {
		var ua = uaT / uB
		var ub = ubT / uB
		if ua >= 0 && ua <= 1 && ub >= 0 && ub <= 1 {
			return true
		}
	}
	return false
}

// ported from terraformer.js https://github.com/Esri/Terraformer/blob/master/terraformer.js#L521-L531
func arrayIntersectsArray(a, b [][]float64) bool {
	for i := 0; i < (len(a) - 1); i++ {
		for j := 0; j < (len(b) - 1); j++ {
			if vertexIntersectsVertex(a[i], a[i+1], b[j], b[j+1]) {
				return true
			}
		}
	}
	return false
}

// ported from terraformer.js https://github.com/Esri/Terraformer/blob/master/terraformer.js#L470-L480
func coordinatesContainPoint(coordinates [][]float64, point []float64) bool {
	contains := false
	l := len(coordinates)
	j := l - 1
	for i := 0; i < l; i++ {
		if ((coordinates[i][1] <= point[1] && point[1] < coordinates[j][1]) ||
			(coordinates[j][1] <= point[1] && point[1] < coordinates[i][1])) &&
			(point[0] < (((coordinates[j][0]-coordinates[i][0])*(point[1]-coordinates[i][1]))/(coordinates[j][1]-coordinates[i][1]))+coordinates[i][0]) {
			contains = !contains
		}
		j = i
	}
	return contains
}

// ported from terraformer-arcgis-parser.js https://github.com/Esri/terraformer-arcgis-parser/blob/master/terraformer-arcgis-parser.js#L106-L113
func coordinatesContainCoordinates(outer, inner [][]float64) bool {
	var intersects = arrayIntersectsArray(outer, inner)
	var contains = coordinatesContainPoint(outer, inner[0])
	if !intersects && contains {
		return true
	}
	return false
}

// do any polygons in this array contain any other polygons in this array?
// used for checking for holes in arcgis rings
// ported from terraformer-arcgis-parser.js https://github.com/Esri/terraformer-arcgis-parser/blob/master/terraformer-arcgis-parser.js#L117-L172
func convertRingsToGeoJSON(rings []Ring) []Polygon {

	polygons := []Polygon{}
	holes := []Ring{}

	// for each ring
	for r := 0; r < len(rings); r++ {
		ring := make(Ring, len(rings[r]))
		copy(ring, rings[r])
		ring = closeRing(ring)
		if len(ring) < 4 {
			continue
		}
		// is this ring an outer ring? is it clockwise?
		if ringIsClockwise(ring) {
			outerRing := Ring(reverse(ring))
			polygon := Polygon{}
			polygon = append(polygon, outerRing)
			polygons = append(polygons, polygon) // push to outer rings
		} else {
			holes = append(holes, reverse(ring)) // wind inner rings clockwise for RFC 7946 compliance
		}
	}

	uncontainedHoles := []Ring{}

	// while there are holes left...
	var hole Ring
	for len(holes) > 0 {
		// pop a hole off out stack
		hole, holes = holes[len(holes)-1], holes[:len(holes)-1]

		// loop over all outer rings and see if they contain our hole.
		contained := false
		for x := (len(polygons) - 1); x >= 0; x-- {
			polygon := polygons[x]
			outerRing := polygon[0]
			if coordinatesContainCoordinates(outerRing, hole) {
				// the hole is contained push it into our polygon
				polygon = append(polygon, hole)
				contained = true
				break
			}
		}

		// ring is not contained in any outer ring
		// sometimes this happens https://github.com/Esri/esri-leaflet/issues/320
		if !contained {
			uncontainedHoles = append(uncontainedHoles, hole)
		}
	}

	// if we couldn't match any holes using contains we can try intersects...
	for len(uncontainedHoles) != 0 {

		// pop a hole off out stack
		hole, uncontainedHoles = uncontainedHoles[len(uncontainedHoles)-1], uncontainedHoles[:len(uncontainedHoles)-1]

		// loop over all outer rings and see if any intersect our hole.
		intersects := false

		for x := len(polygons) - 1; x >= 0; x-- {
			polygon := polygons[x]
			outerRing := polygon[0]
			if arrayIntersectsArray(outerRing, hole) {
				// the hole is contained push it into our polygon
				polygon = append(polygon, hole)
				intersects = true
				break
			}
		}

		if !intersects {
			polygon := Polygon{}
			polygon = append(polygon, Ring(reverse(hole)))
			polygons = append(polygons, polygon)
		}
	}

	return polygons
}

// This function ensures that rings are oriented in the right directions
// outer rings are clockwise, holes are counterclockwise
// used for converting GeoJSON Polygons to ArcGIS Polygons
func orientRings(poly [][][]float64) [][][]float64 {
	output := [][][]float64{}
	polygon := [][][]float64{}
	copy(polygon, poly)
	ring, polygon := polygon[0], polygon[1:]
	polygon2 := [][][]float64{}
	copy(polygon2, polygon)
	outerRing := closeRing(ring)
	if len(outerRing) >= 4 {
		if !ringIsClockwise(outerRing) {
			for i := len(outerRing)/2 - 1; i >= 0; i-- {
				opp := len(outerRing) - 1 - i
				outerRing[i], outerRing[opp] = outerRing[opp], outerRing[i]
			}
		}
		output = append(output, outerRing)
		for i := 0; i < len(polygon); i++ {
			polygon2 := [][][]float64{}
			copy(polygon2, polygon)
			hole := closeRing(polygon2[i])
			if len(hole) >= 4 {
				if ringIsClockwise(hole) {
					for i := len(hole)/2 - 1; i >= 0; i-- {
						opp := len(hole) - 1 - i
						hole[i], hole[opp] = hole[opp], hole[i]
					}
				}
				output = append(output, hole)
			}
		}
	}
	return output
}

// This function flattens holes in multipolygons to one array of polygons
// used for converting GeoJSON Polygons to ArcGIS Polygons
func flattenMultiPolygonRings(rings [][][][]float64) [][][]float64 {
	output := [][][]float64{}
	for i := 0; i < len(rings); i++ {
		polygon := orientRings(rings[i])
		for x := (len(polygon) - 1); x >= 0; x-- {
			ring := [][]float64{}
			copy(ring, polygon[x])
			output = append(output, ring)
		}
	}
	return output
}

func getId(attributes map[string]interface{}, idAttribute string) (interface{}, error) {
	for k, v := range attributes {
		if k == idAttribute {
			return v, nil
		}
	}
	keys := []string{"OBJECTID", "FID"}
	for _, key := range keys {
		for k, v := range attributes {
			if k == key {
				return v, nil
			}
		}

	}
	return nil, errors.New("no valid id attribute found")
}
