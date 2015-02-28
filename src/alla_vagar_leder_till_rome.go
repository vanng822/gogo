package main

import (
	"github.com/vanng822/gpsutil"
	"fmt"
	"math/rand"
	"flag"
)

func main() {
	var k, p int
	flag.IntVar(&k, "k", 10, "Number of km")
	flag.IntVar(&p, "p", 4, "Number of points per km")
	flag.Parse()
	var point *gpsutil.LatLng
	for i := 1; i <= k; i++ {
		for j := 0; j < p; j++ {
			point = gpsutil.GetPointByBearing(41.8723889, 12.48018019999995, float64(rand.Intn(360)), float64(i) * 1000)
			fmt.Println(fmt.Sprintf("%dkm:%f,%f",i, point.Lat(), point.Lng()))
		}
	}
}