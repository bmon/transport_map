package main

import (
	"context"
	"log"

	geo "github.com/kellydunn/golang-geo"
	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

var key string = "AIzaSyCTRKnGLrsBszA_Vs12CWsYYTFcEjQNSrk"

func main() {
	//syd := Destination{"Sydney", "sydney cbd", -33.870848, 151.207321}
	//createDB(syd)
	BuildPolyline(geo.NewPoint(-33.862510, 151.159305), geo.NewPoint(-33.947907, 151.270602), .5)
}

func distMatrixExample() {
	c, err := maps.NewClient(maps.WithAPIKey(key))
	if err != nil {
		log.Fatalf("fatal client creation error: %s", err)
	}

	r := &maps.DistanceMatrixRequest{
		Origins:       []string{"UNSW, Sydney", "International House, Wollongong"},
		Destinations:  []string{"Sydney CBD"},
		Mode:          "ModeTransit",
		DepartureTime: "1501452000",
		Units:         "UnitsMetric",
	}
	matrix, err := c.DistanceMatrix(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal directions error: %s", err)
	}

	pretty.Println(matrix)
}
