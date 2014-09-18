package main

import (
	"github.com/vanng822/gpsutil"
	"fmt"
	"flag"
)

func main() {
	var lat1, lat2, lng1, lng2 float64
	flag.Float64Var(&lat1, "lat1", 0.0, "Latitude point 1")
	flag.Float64Var(&lat2, "lat2", 0.0, "Latitude point 2")
	flag.Float64Var(&lng1, "lng1", 0.0, "Longitude point 1")
	flag.Float64Var(&lng2, "lng2", 0.0, "Longitude point 2")
	flag.Parse()
	fmt.Println(fmt.Sprintf("%f", gpsutil.GetDistance(lat1, lng1, lat2, lng2)))
}