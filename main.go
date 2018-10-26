/*
Purpose:
- GPX (GPS Exchange Format) Distance Calculator

Description:
- Shortest, longest distance between a given location and the GPX waypoints.

Releases:
- 1.0.0 - 2018/10/26 : initial release

Author:
- Klaus Tockloth

Copyright and license:
- Copyright (c) 2018 Klaus Tockloth
- MIT license

Permission is hereby granted, free of charge, to any person obtaining a copy of this software
and associated documentation files (the Software), to deal in the Software without restriction,
including without limitation the rights to use, copy, modify, merge, publish, distribute,
sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or
substantial portions of the Software.

The software is provided 'as is', without warranty of any kind, express or implied, including
but not limited to the warranties of merchantability, fitness for a particular purpose and
noninfringement. In no event shall the authors or copyright holders be liable for any claim,
damages or other liability, whether in an action of contract, tort or otherwise, arising from,
out of or in connection with the software or the use or other dealings in the software.

Contact (eMail):
- freizeitkarte@googlemail.com

Remarks:
- The haversine formula determines the great-circle distance between two points on a sphere given
  their longitudes and latitudes. Important in navigation, it is a special case of a more general
  formula in spherical trigonometry, the law of haversines, that relates the sides and angles of
  spherical triangles. (Wikipedia)

Links:
- https://github.com/tkrajina/gpxgo
- https://gist.github.com/cdipaolo/d3f8db3848278b49db68
- https://en.wikipedia.org/wiki/Haversine_formula
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/tkrajina/gpxgo/gpx"
)

// general program info
var (
	progName    = os.Args[0]
	progVersion = "1.0.0"
	progDate    = "2018/10/26"
	progPurpose = "GPX (GPS Exchange Format) Distance Calculator (distances in meters)"
	progInfo    = "Shortest, longest distance between a given point and the GPX points."
)

/*
init initializes this program
*/
func init() {

	// initialize logger
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

/*
main starts this program
*/
func main() {

	gpxfile := flag.String("gpxfile", "", "GPX file to parse")
	lat := flag.String("lat", "", "latitude (decimal degrees) of given point")
	lon := flag.String("lon", "", "longitude (decimal degrees) of given point")

	flag.Usage = printUsage
	flag.Parse()

	if *gpxfile == "" || *lat == "" || *lon == "" {
		printUsage()
	}

	latitude, err := strconv.ParseFloat(*lat, 64)
	if err != nil {
		log.Fatalf("error <%v> at strconv.ParseFloat()", err)
	}
	longitude, err := strconv.ParseFloat(*lon, 64)
	if err != nil {
		log.Fatalf("error <%v> at strconv.ParseFloat()", err)
	}

	filename := *gpxfile
	gpxBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error <%v> at ioutil.ReadFile(); filename = <%v>", err, filename)
	}

	gpxFile, err := gpx.ParseBytes(gpxBytes)
	if err != nil {
		log.Fatalf("error <%v> at gpx.ParseBytes()", err)
	}

	shortestDistance := math.MaxInt32
	longestDistance := math.MinInt32

	// track data
	for _, track := range gpxFile.Tracks {
		for _, segment := range track.Segments {
			for _, point := range segment.Points {
				tmpDistance := int(Distance(latitude, longitude, point.Latitude, point.Longitude))
				if tmpDistance > longestDistance {
					longestDistance = tmpDistance
				}
				if tmpDistance < shortestDistance {
					shortestDistance = tmpDistance
				}
			}
		}
	}

	// print results (csv format)
	fmt.Printf("\"%s\",%s,%s,%d,%d\n", *gpxfile, *lat, *lon, shortestDistance, longestDistance)

	os.Exit(0)
}

/*
printUsage prints the usage of this program
*/
func printUsage() {

	fmt.Printf("\nProgram:\n")
	fmt.Printf("  Name    : %s\n", progName)
	fmt.Printf("  Release : %s - %s\n", progVersion, progDate)
	fmt.Printf("  Purpose : %s\n", progPurpose)
	fmt.Printf("  Info    : %s\n", progInfo)

	fmt.Printf("\nUsage:\n")
	fmt.Printf("  %s -gpxfile=filename -lat=latitude -lon=longitude\n", progName)

	fmt.Printf("\nExample:\n")
	fmt.Printf("  %s -gpxfile=ellenbogen.gpx -lat=55.05 -lon=8.41\n", progName)

	fmt.Printf("\nOptions:\n")
	flag.PrintDefaults()

	fmt.Printf("\nOutput:\n")
	fmt.Printf("  gpxfile,lat,lon,shortest,longest\n")

	fmt.Printf("\nExample:\n")
	fmt.Printf("  \"ellenbogen.gpx\",55.05,8.41,1066,3425\n")

	fmt.Printf("\n")
	os.Exit(1)
}

/*
Distance returns the distance (in meters) between two points of a given longitude and latitude relatively
accurately (using a spherical approximation of the Earth) through the Haversin Distance Formula for great
arc distance on a sphere with accuracy for small distances.
*/
func Distance(lat1, lon1, lat2, lon2 float64) float64 {

	// convert to radians
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // earth radius in meters

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

/*
hsin implements the haversin(θ) function
*/
func hsin(theta float64) float64 {

	return math.Pow(math.Sin(theta/2), 2)
}
