# GPX (GPS Exchange Format) Distance Calculator

gpxdist is an utility to calculate the shortest/longest distance between a given point (in decimal degrees) and all points
of a GPX file. The calulation is done by using the Haversin Distance Formula for great
arc distance on a sphere (approximation with good accuracy for small distances).

## Usage

``` text
Program:
  Name    : ./gpxdist
  Release : 1.0.0 - 2018/10/26
  Purpose : GPX (GPS Exchange Format) Distance Calculator (distances in meters)
  Info    : Shortest, longest distance between a given point and the GPX points.

Usage:
  ./gpxdist -gpxfile=filename -lat=latitude -lon=longitude

Example:
  ./gpxdist -gpxfile=ellenbogen.gpx -lat=55.05 -lon=8.41

Options:
  -gpxfile string
        GPX file to parse
  -lat string
        latitude (decimal degrees) of given point
  -lon string
        longitude (decimal degrees) of given point

Output:
  gpxfile,lat,lon,shortest,longest

Example:
  "ellenbogen.gpx",55.05,8.41,1066,3425
```
