package main

import (
	"context"
	"fmt"
	"log"

	"googlemaps.github.io/maps"

	geo "github.com/kellydunn/golang-geo"
	"github.com/kr/pretty"
)

func GetDistMatrix(origins []string, dest []string) *maps.DistanceMatrixResponse {
	c, err := maps.NewClient(maps.WithAPIKey(Key))
	if err != nil {
		log.Fatalf("fatal client creation error: %s", err)
	}

	r := &maps.DistanceMatrixRequest{
		Origins:       origins,
		Destinations:  dest,
		Mode:          "ModeTransit",
		DepartureTime: "1501452000",
		Units:         "UnitsMetric",
	}
	matrix, err := c.DistanceMatrix(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal directions error: %s", err)
	}

	return matrix
}

func DistMatrixExample() {
	pretty.Println(GetDistMatrix([]string{"UNSW, Sydney", "International House, Wollongong"}, []string{"Sydney CBD"}))
}

func BuildPolyline(p1, p2 *geo.Point, dist float64) []string {
	// We assume p1 is the top left and p2 is the bottom right.
	// dist is in km

	points := make([]maps.LatLng, 0)
	// Make a temp copy of our point
	p_start := geo.NewPoint(p1.Lat(), p1.Lng())

	// while p is still "above" p2
	for p_start.BearingTo(p2) > 90 {
		// reset p to the start of the "row"
		p := geo.NewPoint(p_start.Lat(), p_start.Lng())
		// while p is still "left" of p2
		for p.BearingTo(p2) < 180 && p.BearingTo(p2) > 0 {
			// add p to points
			points = append(points, maps.LatLng{p.Lat(), p.Lng()})

			// move p left dist kms (left 1 column)
			p = p.PointAtDistanceAndBearing(dist, 90)
		}
		// move p_start down a "row"
		p_start = p_start.PointAtDistanceAndBearing(dist, 180)
	}
	fmt.Println(len(points), " coords to be encoded.")

	results := make([]string, 0)
	i, prev := 25, 0
	for ; i < len(points); i += 25 {
		poly := maps.Encode(points[prev:i])
		fmt.Printf("%s\n", poly)
		results = append(results, poly)
		prev = i
	}
	poly := maps.Encode(points[prev:i])
	fmt.Printf("%s\n", poly)
	results = append(results, poly)

	return results
}
