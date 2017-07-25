package main

import (
	"fmt"

	geo "github.com/kellydunn/golang-geo"
	polyline "github.com/twpayne/go-polyline"
)

func BuildPolyline(p1, p2 *geo.Point, dist float64) {
	// We assume p1 is the top left and p2 is the bottom right.
	// dist is in km

	points := make([][]float64, 0)
	// Make a temp copy of our point
	p_start := geo.NewPoint(p1.Lat(), p1.Lng())

	// while p is still "above" p2
	for p_start.BearingTo(p2) > 90 {
		// reset p to the start of the "row"
		p := geo.NewPoint(p_start.Lat(), p_start.Lng())
		// while p is still "left" of p2
		for p.BearingTo(p2) < 180 && p.BearingTo(p2) > 0 {
			// add p to points
			points = append(points, []float64{p.Lat(), p.Lng()})

			// move p left dist kms (left 1 column)
			p = p.PointAtDistanceAndBearing(dist, 90)
		}
		// move p_start down a "row"
		p_start = p_start.PointAtDistanceAndBearing(dist, 180)
	}
	fmt.Println(len(points), " coords to be encoded.")
	fmt.Printf("%s\n", polyline.EncodeCoords(points))

}
