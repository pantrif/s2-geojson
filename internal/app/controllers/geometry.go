package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/geo/s1"
	"github.com/golang/geo/s2"
	"s2_geojson/pkg/geo"
	"strconv"
	"strings"
)

// GeometryController struct
type GeometryController struct{}

// Cover uses s2 region coverer to cover geometries of geojson (only points and polygons supported)
func (u GeometryController) Cover(c *gin.Context) {
	gJSON := []byte(c.PostForm("geojson"))
	maxLevel, err := strconv.Atoi(c.PostForm("max_level_geojson"))
	minLevel, err := strconv.Atoi(c.PostForm("min_level_geojson"))

	fs, err := geo.DecodeGeoJSON(gJSON)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var tokens []string
	var s2cells [][][]float64

	for _, f := range fs {

		if f.Geometry.IsPolygon() {
			for _, p := range f.Geometry.Polygon {
				p := geo.PointsToPolygon(p)
				_, t, c := geo.CoverPolygon(p, maxLevel, minLevel)
				s2cells = append(s2cells, c...)
				tokens = append(tokens, t...)
			}
		}
		if f.Geometry.IsPoint() {
			point := geo.Point{Lat: f.Geometry.Point[1], Lng: f.Geometry.Point[0]}
			_, t, c := geo.CoverPoint(point, maxLevel)
			s2cells = append(s2cells, c...)
			tokens = append(tokens, t)
		}
	}

	c.JSON(200, gin.H{
		"max_level_geojson": maxLevel,
		"cell_tokens":       strings.Join(tokens, ","),
		"cells":             s2cells,
	})
}

// CheckIntersection checks intersection of geoJSON geometries with a point and with a circle
func (u GeometryController) CheckIntersection(c *gin.Context) {
	lat, err := strconv.ParseFloat(c.PostForm("lat"), 64)
	lng, err := strconv.ParseFloat(c.PostForm("lng"), 64)
	radius, err := strconv.ParseFloat(c.PostForm("radius"), 64)

	gJSON := []byte(c.PostForm("geojson"))
	maxLevel, err := strconv.Atoi(c.PostForm("max_level_geojson"))
	minLevel, err := strconv.Atoi(c.PostForm("min_level_geojson"))
	maxLevelCircle, err := strconv.Atoi(c.PostForm("max_level_circle"))

	fs, err := geo.DecodeGeoJSON(gJSON)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	angle := s1.Angle((radius / 1000) / geo.EarthRadius)
	ca := s2.CapFromCenterAngle(s2.PointFromLatLng(s2.LatLngFromDegrees(lat, lng)), angle)
	circeCov := &s2.RegionCoverer{MaxLevel: maxLevelCircle, MaxCells: 300}
	circleRegion := s2.Region(ca)
	circleCovering := circeCov.Covering(circleRegion)

	var values []string
	var s2cells [][][]float64

	for _, c := range circleCovering {
		c1 := s2.CellFromCellID(s2.CellIDFromToken(c.ToToken()))

		var s2cell [][]float64
		for i := 0; i < 4; i++ {
			latlng := s2.LatLngFromPoint(c1.Vertex(i))
			s2cell = append(s2cell, []float64{latlng.Lat.Degrees(), latlng.Lng.Degrees()})
		}

		s2cells = append(s2cells, s2cell)

		values = append(values, c.ToToken())
	}

	ll := s2.LatLngFromDegrees(lat, lng)
	cell := s2.CellFromLatLng(ll)

	intersectsPoint, intersectsCircle := false, false

	for _, f := range fs {

		if f.Geometry.IsPolygon() {
			for _, p := range f.Geometry.Polygon {
				p := geo.PointsToPolygon(p)
				covering, _, _ := geo.CoverPolygon(p, maxLevel, minLevel)

				if covering.IntersectsCell(cell) {
					intersectsPoint = true
				}
				if covering.Intersects(circleCovering) {
					intersectsCircle = true
				}
			}
		}

		if f.Geometry.IsPoint() {
			point := geo.Point{Lat: f.Geometry.Point[1], Lng: f.Geometry.Point[0]}
			cc, _, _ := geo.CoverPoint(point, maxLevel)

			if cell.IntersectsCell(cc) {
				intersectsPoint = true
			}
			if circleCovering.IntersectsCell(cc) {
				intersectsCircle = true
			}
		}
	}

	c.JSON(200, gin.H{
		"intersects_with_point":  intersectsPoint,
		"intersects_with_circle": intersectsCircle,
		"radius":                 radius,
		"cells":                  s2cells,
	})
}
