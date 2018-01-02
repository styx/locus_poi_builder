package main

import "strconv"
import "github.com/paulmach/go.geo"

// GetPolygonCentroid - compute the centroid of a polygon set
// using a spherical co-ordinate system
func GetPolygonCentroid(ps *geo.PointSet) *geo.Point {
	// GeoCentroid function added in https://github.com/paulmach/go.geo/pull/24
	return ps.GeoCentroid()
}

// GetLineCentroid - compute the centroid of a line string
func GetLineCentroid(ps *geo.PointSet) *geo.Point {
	path := geo.NewPath()
	path.PointSet = *ps

	halfDistance := path.Distance() / 2
	travelled := 0.0

	for i := 0; i < len(path.PointSet)-1; i++ {
		segment := geo.NewLine(&path.PointSet[i], &path.PointSet[i+1])
		distance := segment.Distance()

		// middle line segment
		if (travelled + distance) > halfDistance {
			var remainder = halfDistance - travelled
			return segment.Interpolate(remainder / distance)
		}

		travelled += distance
	}

	return ps.GeoCentroid()
}


// compute the centroid of a way
func computeCentroid(latlons []map[string]string) map[string]string {
	// convert lat/lon map to geo.PointSet
	points := geo.NewPointSet()
	for _, each := range latlons {
		var lon, _ = strconv.ParseFloat(each["lon"], 64)
		var lat, _ = strconv.ParseFloat(each["lat"], 64)
		points.Push(geo.NewPoint(lon, lat))
	}

	// determine if the way is a closed centroid or a linestring
	// by comparing first and last coordinates.
	isClosed := false
	if points.Length() > 2 {
		isClosed = points.First().Equals(points.Last())
	}

	// compute the centroid using one of two different algorithms
	var compute *geo.Point
	if isClosed {
		compute = GetPolygonCentroid(points)
	} else {
		compute = GetLineCentroid(points)
	}

	// return point as lat/lon map
	var centroid = make(map[string]string)
	centroid["lat"] = strconv.FormatFloat(compute.Lat(), 'f', 6, 64)
	centroid["lon"] = strconv.FormatFloat(compute.Lng(), 'f', 6, 64)

	return centroid
}
