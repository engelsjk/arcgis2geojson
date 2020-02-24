package arcgis2geojson

import (
	"encoding/json"
	"errors"

	"github.com/paulmach/orb"
	geojson "github.com/paulmach/orb/geojson"
)

func Convert(data []byte, idAttribute string) ([]byte, error) {

	arcgisJSON := ArcGISJSON{}

	err := json.Unmarshal(data, &arcgisJSON)
	if err != nil {
		return nil, err
	}

	if arcgisJSON.SpatialReference.WKID != 4326 {
		return nil, errors.New("error: arc gis features must be in wkid 4326 for valid conversion to geojson")
	}

	fc := geojson.NewFeatureCollection()

	if len(arcgisJSON.Features) != 0 {
		for i := 0; i < len(arcgisJSON.Features); i++ {
			f := arcgisJSON.Features[i]
			feature := featureToFeature(f, idAttribute)
			fc.Features = append(fc.Features, feature)
		}
	}
	return fc.MarshalJSON()
}

func featureToFeature(f ArcGISFeature, idAttribute string) *geojson.Feature {

	var feature = new(geojson.Feature)

	// x,y >> point
	if f.X != 0 && f.Y != 0 {
		point := [][]float64{
			[]float64{f.X, f.Y},
		}
		feature = pointsToFeature(point)
		// support for f.Z != 0?
	}

	// points >> point/multipoint
	if len(f.Points) != 0 {
		feature = pointsToFeature(f.Points)
	}
	if len(f.Geometry.Points) != 0 {
		feature = pointsToFeature(f.Geometry.Points)
	}

	// paths >> linestring/multilinestring
	if len(f.Paths) != 0 {
		feature = pathsToFeature(f.Paths)
	}
	if len(f.Geometry.Paths) != 0 {
		feature = pathsToFeature(f.Geometry.Paths)
	}

	// rings >> polygon/multipolygon
	if len(f.Rings) != 0 {
		feature = ringsToFeature(f.Rings)
	}
	if len(f.Geometry.Rings) != 0 {
		feature = ringsToFeature(f.Geometry.Rings)
	}

	// xmin/xmax/ymin/ymax >> bounding box (polygon)
	bbox := []float64{f.Xmin, f.Ymin, f.Xmax, f.Ymax}
	if bbox[0] != 0 && bbox[1] != 0 && bbox[2] != 0 && bbox[3] != 0 {
		feature = boundingBoxToFeature(bbox)
	}

	// add properties
	feature.Properties = make(map[string]interface{})
	for k, v := range f.Attributes {
		feature.Properties[k] = v
	}
	// add id
	id, err := getId(f.Attributes, idAttribute)
	if err == nil {
		feature.ID = id
	}

	return feature
}

// structs

type ArcGISFeature struct {
	Attributes map[string]interface{} `json:"attributes"`
	Geometry   struct {
		Points [][]float64   `json:"points"`
		Paths  [][][]float64 `json:"paths"`
		Rings  [][][]float64 `json:"rings"`
	} `json:"geometry"`
	X      float64       `json:"x"`
	Y      float64       `json:"y"`
	Z      float64       `json:"z"`
	Xmin   float64       `json:"xmin"`
	Xmax   float64       `json:"xmax"`
	Ymin   float64       `json:"ymin"`
	Ymax   float64       `json:"ymax"`
	Paths  [][][]float64 `json:"paths"`
	Points [][]float64   `json:"points"`
	Rings  [][][]float64 `json:"rings"`
}

type ArcGISJSON struct {
	DisplayFieldName string            `json:"displayFieldName"`
	FieldAliases     map[string]string `json:"fieldAliases"`
	GeometryType     string
	SpatialReference struct {
		WKID       int
		LatestWKID int
	} `json:"spatialReference"`
	Fields []struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		Alias  string `json:"alias"`
		Length int    `json:"length"`
	} `json:"fields"`
	Features []ArcGISFeature `json:"features"`
}

// feature conversions

func pointsToFeature(points [][]float64) *geojson.Feature {
	var feature = new(geojson.Feature)
	if len(points) == 0 {
		return feature
	}
	if len(points) == 1 {
		p := orb.Point{points[0][0], points[0][1]}
		feature = geojson.NewFeature(p)
	} else {
		mp := orb.MultiPoint{}
		for _, p := range points {
			mp = append(mp, orb.Point{p[0], p[1]})
		}
		feature = geojson.NewFeature(mp)
	}
	return feature
}

func pathsToFeature(paths [][][]float64) *geojson.Feature {
	var feature = new(geojson.Feature)
	if len(paths) == 0 {
		return feature
	}
	if len(paths) == 1 {
		ls := orb.LineString{}
		for _, pt := range paths[0] {
			ls = append(ls, orb.Point{pt[0], pt[1]})
		}
		feature = geojson.NewFeature(ls)
	} else {
		mls := orb.MultiLineString{}
		for _, path := range paths {
			ls := orb.LineString{}
			for _, pt := range path {
				ls = append(ls, orb.Point{pt[0], pt[1]})
			}
			mls = append(mls, ls)
		}
		feature = geojson.NewFeature(mls)
	}
	return feature
}

func ringsToFeature(rings [][][]float64) *geojson.Feature {
	var feature = new(geojson.Feature)
	if len(rings) == 0 {
		return feature
	}
	outerRings := convertRingsToGeoJSON(rings)
	newRings := []orb.Ring{}
	for _, r := range outerRings {
		ring := orb.Ring{}
		for _, pt := range r {
			ring = append(ring, orb.Point{pt[0], pt[1]})
		}
		newRings = append(newRings, ring)
	}
	// TODO: if len(outerRings) == 0?
	if len(outerRings) == 1 {
		p := orb.Polygon{newRings[0]}
		feature = geojson.NewFeature(p)
	} else {
		mp := orb.MultiPolygon{newRings}
		feature = geojson.NewFeature(mp)
	}
	return feature
}

func boundingBoxToFeature(bbox []float64) *geojson.Feature {
	ring := orb.Ring{}
	ring = append(ring, orb.Point{bbox[2], bbox[3]})
	ring = append(ring, orb.Point{bbox[0], bbox[3]})
	ring = append(ring, orb.Point{bbox[0], bbox[1]})
	ring = append(ring, orb.Point{bbox[2], bbox[1]})
	ring = append(ring, orb.Point{bbox[2], bbox[3]})
	p := orb.Polygon{ring}
	return geojson.NewFeature(p)
}
