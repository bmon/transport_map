package main

import (
	geo "github.com/kellydunn/golang-geo"
	"googlemaps.github.io/maps"
)

func main() {
	syd := Destination{"Sydney", -33.870848, 151.207321}
	createDB(syd)
	polylines := BuildPolyline(geo.NewPoint(-33.862510, 151.159305), geo.NewPoint(-33.947907, 151.270602), .5)
	for _, poly := range polylines {
		matrix := GetDistMatrix([]string{"enc:" + poly + ":"}, []string{syd.Name})
		writeDistMatrix(((&maps.Polyline{poly}).Decode()), matrix)
	}
}
