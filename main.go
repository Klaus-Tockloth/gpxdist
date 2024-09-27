/*
Purpose:
- GPX (GPS Exchange Format) Distance Calculator

Description:
- Shortest, longest distance between a given location and the GPX waypoints.

Releases:
- v1.0.0 - 2018/10/26 : initial release
- v1.1.0 - 2024/09/27 : gpxgo lib v1.4, go v1.23.1, Haversine func from gpxgo lib

Author:
- Klaus Tockloth

Copyright and license:
- Copyright (c) 2018-2024 Klaus Tockloth
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
	"log"
	"math"
	"os"
	"strconv"

	"github.com/tkrajina/gpxgo/gpx"
)

// general program info
var (
	progName    = os.Args[0]
	progVersion = "v1.1.0"
	progDate    = "2024/09/27"
	progPurpose = "GPX (GPS Exchange Format) Distance Calculator (distances in meters)"
	progInfo    = "Shortest, longest distance between a given point and the GPX points."
)

/*
main starts this program.
*/
func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

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
	gpxBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error <%v> at os.ReadFile(); filename = <%v>", err, filename)
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
				tmpDistance := int(math.Round(gpx.HaversineDistance(latitude, longitude, point.Latitude, point.Longitude)))
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
}

/*
printUsage prints the usage of this program.
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
