package arcgis2geojson

import "testing"

func TestConversion1(t *testing.T) {
	data := []byte(`{
		"displayFieldName": "FULLNAME",
		"fieldAliases": {
		 "PIN": "PIN",
		 "MINOR": "MINOR",
		 "OBJECTID": "ObjectID",
		 "Shape.area": "shape.area",
		 "Shape.len": "shape.len",
		 "MAJOR": "MAJOR",
		 "ADDR_HN": "HouseNum",
		 "ADDR_FULL": "FullAddr",
		 "ZIP5": "ZIP5",
		 "CTYNAME": "CTYNAME",
		 "POSTALCTYNAME": "POSTALCTYNAME",
		 "PROP_NAME": "PROP_NAME",
		 "PLAT_NAME": "PLAT_NAME",
		 "LOTSQFT": "LOTSQFT",
		 "APPRLNDVAL": "APPRLNDVAL",
		 "APPR_IMPR": "APPR_IMPR",
		 "ANNEXING_CITY": "ANNEXING_CITY",
		 "PAAUNIQUENAME": "PAAUNIQUENAME",
		 "PROPTYPE": "PROPTYPE",
		 "KCA_ZONING": "KCA_ZONING",
		 "KCA_ACRES": "KCA_ACRES",
		 "PREUSE_CODE": "Present Use Code",
		 "PREUSE_DESC": "Present Use Description"
		},
		"geometryType": "esriGeometryPolygon",
		"spatialReference": {
		 "wkid": 4326,
		 "latestWkid": 4326
		},
		"fields": [
		 {
		  "name": "PIN",
		  "type": "esriFieldTypeString",
		  "alias": "PIN",
		  "length": 10
		 },
		 {
		  "name": "MINOR",
		  "type": "esriFieldTypeString",
		  "alias": "MINOR",
		  "length": 4
		 },
		 {
		  "name": "OBJECTID",
		  "type": "esriFieldTypeOID",
		  "alias": "ObjectID"
		 },
		 {
		  "name": "Shape.area",
		  "type": "esriFieldTypeDouble",
		  "alias": "shape.area"
		 },
		 {
		  "name": "Shape.len",
		  "type": "esriFieldTypeDouble",
		  "alias": "shape.len"
		 },
		 {
		  "name": "MAJOR",
		  "type": "esriFieldTypeString",
		  "alias": "MAJOR",
		  "length": 6
		 },
		 {
		  "name": "ADDR_HN",
		  "type": "esriFieldTypeString",
		  "alias": "HouseNum",
		  "length": 10
		 },
		 {
		  "name": "ADDR_FULL",
		  "type": "esriFieldTypeString",
		  "alias": "FullAddr",
		  "length": 120
		 },
		 {
		  "name": "ZIP5",
		  "type": "esriFieldTypeString",
		  "alias": "ZIP5",
		  "length": 5
		 },
		 {
		  "name": "CTYNAME",
		  "type": "esriFieldTypeString",
		  "alias": "CTYNAME",
		  "length": 28
		 },
		 {
		  "name": "POSTALCTYNAME",
		  "type": "esriFieldTypeString",
		  "alias": "POSTALCTYNAME",
		  "length": 28
		 },
		 {
		  "name": "PROP_NAME",
		  "type": "esriFieldTypeString",
		  "alias": "PROP_NAME",
		  "length": 50
		 },
		 {
		  "name": "PLAT_NAME",
		  "type": "esriFieldTypeString",
		  "alias": "PLAT_NAME",
		  "length": 50
		 },
		 {
		  "name": "LOTSQFT",
		  "type": "esriFieldTypeInteger",
		  "alias": "LOTSQFT"
		 },
		 {
		  "name": "APPRLNDVAL",
		  "type": "esriFieldTypeDouble",
		  "alias": "APPRLNDVAL"
		 },
		 {
		  "name": "APPR_IMPR",
		  "type": "esriFieldTypeDouble",
		  "alias": "APPR_IMPR"
		 },
		 {
		  "name": "ANNEXING_CITY",
		  "type": "esriFieldTypeString",
		  "alias": "ANNEXING_CITY",
		  "length": 2
		 },
		 {
		  "name": "PAAUNIQUENAME",
		  "type": "esriFieldTypeString",
		  "alias": "PAAUNIQUENAME",
		  "length": 100
		 },
		 {
		  "name": "PROPTYPE",
		  "type": "esriFieldTypeString",
		  "alias": "PROPTYPE",
		  "length": 1
		 },
		 {
		  "name": "KCA_ZONING",
		  "type": "esriFieldTypeString",
		  "alias": "KCA_ZONING",
		  "length": 20
		 },
		 {
		  "name": "KCA_ACRES",
		  "type": "esriFieldTypeDouble",
		  "alias": "KCA_ACRES"
		 },
		 {
		  "name": "PREUSE_CODE",
		  "type": "esriFieldTypeSmallInteger",
		  "alias": "Present Use Code"
		 },
		 {
		  "name": "PREUSE_DESC",
		  "type": "esriFieldTypeString",
		  "alias": "Present Use Description",
		  "length": 50
		 }
		],
		"features": [
		 {
		  "attributes": {
		   "PIN": "0723059046",
		   "MINOR": "9046",
		   "OBJECTID": 48772,
		   "Shape.area": 101139.11680800001,
		   "Shape.len": 2125.1325428451037,
		   "MAJOR": "072305",
		   "ADDR_HN": "737",
		   "ADDR_FULL": "737 LOGAN AVE N",
		   "ZIP5": "98057",
		   "CTYNAME": "RENTON",
		   "POSTALCTYNAME": null,
		   "PROP_NAME": "BOEING VACANT LAND",
		   "PLAT_NAME": "",
		   "LOTSQFT": 101141,
		   "APPRLNDVAL": 2022800,
		   "APPR_IMPR": 0,
		   "ANNEXING_CITY": null,
		   "PAAUNIQUENAME": null,
		   "PROPTYPE": "C",
		   "KCA_ZONING": "UC",
		   "KCA_ACRES": 2.3218778599999998,
		   "PREUSE_CODE": 316,
		   "PREUSE_DESC": "Vacant(Industrial)                                "
		  },
		  "geometry": {
		   "rings": [
			[
			 [
			  -122.21207743860465,
			  47.491934854980641
			 ],
			 [
			  -122.21047981064409,
			  47.491946358462215
			 ],
			 [
			  -122.21047971473607,
			  47.491951842114545
			 ],
			 [
			  -122.21036544902873,
			  47.491952663274546
			 ],
			 [
			  -122.21033739477114,
			  47.491952865076485
			 ],
			 [
			  -122.21031346576503,
			  47.491217008832088
			 ],
			 [
			  -122.21032865848835,
			  47.491220995036706
			 ],
			 [
			  -122.21121456105593,
			  47.491453400905208
			 ],
			 [
			  -122.21198646780773,
			  47.491658174687636
			 ],
			 [
			  -122.21207743860465,
			  47.491934854980641
			 ]
			],
			[
			 [
			  -122.21153936410711,
			  47.491362810654692
			 ],
			 [
			  -122.21044247866021,
			  47.491074434755546
			 ],
			 [
			  -122.21179003416054,
			  47.491060650119863
			 ],
			 [
			  -122.21192307370254,
			  47.491465281914692
			 ],
			 [
			  -122.21153936410711,
			  47.491362810654692
			 ]
			]
		   ]
		  }
		 }
		]
	   }`)

	b, err := Convert(data, "PIN")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(b))
}

func TestConversion2(t *testing.T) {
	data := []byte(`{"displayFieldName":"FULLNAME","fieldAliases":{"PIN":"PIN","MINOR":"MINOR","OBJECTID":"ObjectID","Shape.area":"shape.area","Shape.len":"shape.len","MAJOR":"MAJOR","ADDR_HN":"HouseNum","ADDR_FULL":"FullAddr","ZIP5":"ZIP5","CTYNAME":"CTYNAME","POSTALCTYNAME":"POSTALCTYNAME","PROP_NAME":"PROP_NAME","PLAT_NAME":"PLAT_NAME","LOTSQFT":"LOTSQFT","APPRLNDVAL":"APPRLNDVAL","APPR_IMPR":"APPR_IMPR","ANNEXING_CITY":"ANNEXING_CITY","PAAUNIQUENAME":"PAAUNIQUENAME","PROPTYPE":"PROPTYPE","KCA_ZONING":"KCA_ZONING","KCA_ACRES":"KCA_ACRES","PREUSE_CODE":"Present Use Code","PREUSE_DESC":"Present Use Description"},"geometryType":"esriGeometryPolygon","spatialReference":{"wkid":4326,"latestWkid":4326},"fields":[{"name":"PIN","type":"esriFieldTypeString","alias":"PIN","length":10},{"name":"MINOR","type":"esriFieldTypeString","alias":"MINOR","length":4},{"name":"OBJECTID","type":"esriFieldTypeOID","alias":"ObjectID"},{"name":"Shape.area","type":"esriFieldTypeDouble","alias":"shape.area"},{"name":"Shape.len","type":"esriFieldTypeDouble","alias":"shape.len"},{"name":"MAJOR","type":"esriFieldTypeString","alias":"MAJOR","length":6},{"name":"ADDR_HN","type":"esriFieldTypeString","alias":"HouseNum","length":10},{"name":"ADDR_FULL","type":"esriFieldTypeString","alias":"FullAddr","length":120},{"name":"ZIP5","type":"esriFieldTypeString","alias":"ZIP5","length":5},{"name":"CTYNAME","type":"esriFieldTypeString","alias":"CTYNAME","length":28},{"name":"POSTALCTYNAME","type":"esriFieldTypeString","alias":"POSTALCTYNAME","length":28},{"name":"PROP_NAME","type":"esriFieldTypeString","alias":"PROP_NAME","length":50},{"name":"PLAT_NAME","type":"esriFieldTypeString","alias":"PLAT_NAME","length":50},{"name":"LOTSQFT","type":"esriFieldTypeInteger","alias":"LOTSQFT"},{"name":"APPRLNDVAL","type":"esriFieldTypeDouble","alias":"APPRLNDVAL"},{"name":"APPR_IMPR","type":"esriFieldTypeDouble","alias":"APPR_IMPR"},{"name":"ANNEXING_CITY","type":"esriFieldTypeString","alias":"ANNEXING_CITY","length":2},{"name":"PAAUNIQUENAME","type":"esriFieldTypeString","alias":"PAAUNIQUENAME","length":100},{"name":"PROPTYPE","type":"esriFieldTypeString","alias":"PROPTYPE","length":1},{"name":"KCA_ZONING","type":"esriFieldTypeString","alias":"KCA_ZONING","length":20},{"name":"KCA_ACRES","type":"esriFieldTypeDouble","alias":"KCA_ACRES"},{"name":"PREUSE_CODE","type":"esriFieldTypeSmallInteger","alias":"Present Use Code"},{"name":"PREUSE_DESC","type":"esriFieldTypeString","alias":"Present Use Description","length":50}],"features":[{"attributes":{"PIN":"1124059028","MINOR":"9028","OBJECTID":77729,"Shape.area":165625.41204,"Shape.len":1982.6789756972437,"MAJOR":"112405","ADDR_HN":"15531","ADDR_FULL":"15531 SE 37TH ST","ZIP5":"98006","CTYNAME":"BELLEVUE","POSTALCTYNAME":"BELLEVUE","PROP_NAME":"TRAILER INNS RV PARK (96 RV SITES)","PLAT_NAME":"","LOTSQFT":146797,"APPRLNDVAL":6018600,"APPR_IMPR":1000,"ANNEXING_CITY":null,"PAAUNIQUENAME":null,"PROPTYPE":"C","KCA_ZONING":"NMU","KCA_ACRES":3.3699954000000001,"PREUSE_CODE":38,"PREUSE_DESC":"Mobile Home Park                                  "},"geometry":{"rings":[[[-122.13301309996658,47.576130560080706],[-122.133622130384,47.576136210402545],[-122.13371157768161,47.576137024642129],[-122.13370690774072,47.576431912684384],[-122.13368919013941,47.577536435925438],[-122.13245272276269,47.57737012471086],[-122.13219560450202,47.57730780494429],[-122.13219792222468,47.577090459670188],[-122.13220619140972,47.576309255214163],[-122.13279761914873,47.576315769082825],[-122.13301309996658,47.576130560080706]],[[-122.13207695928254,47.576144597881537],[-122.13207725725498,47.576119179977262],[-122.13231170189988,47.576122342367427],[-122.13270880760213,47.576127731150962],[-122.13268442042489,47.576149067462055],[-122.13207695928254,47.576144597881537]]]}}]}`)

	b, err := Convert(data, "PIN")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(b))
}
