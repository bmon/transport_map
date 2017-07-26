package main

import geo "github.com/kellydunn/golang-geo"

func main() {
	syd := Destination{"Sydney", "sydney cbd", -33.870848, 151.207321}
	createDB(syd)
	polylines := BuildPolyline(geo.NewPoint(-33.862510, 151.159305), geo.NewPoint(-33.947907, 151.270602), .5)
	matrx := GetDistMatrix(polylines[0], []string{syd.
}
