package geo

import (
	"github.com/golang/geo/s2"
	"github.com/paulmach/go.geojson"
)

const (
	EarthRadius = 6371.01
	maxCells    = 100
)

type Point struct {
	Lat float64
	Lng float64
}

func DecodeGeoJSON(json []byte) ([]*geojson.Feature, error) {
	f, err := geojson.UnmarshalFeatureCollection(json)
	if err != nil {
		return nil, err
	}
	return f.Features, nil
}

func GeoJSONPointsToPolygon(points [][]float64) *s2.Polygon {
	var pts []s2.Point
	for _, pt := range points {
		pts = append(pts, s2.PointFromLatLng(s2.LatLngFromDegrees(pt[1], pt[0])))
	}
	loop := s2.LoopFromPoints(pts)

	return s2.PolygonFromLoops([]*s2.Loop{loop})
}

func CoverPolygon(p *s2.Polygon, maxLevel, minLevel int) (s2.CellUnion, []string, [][][]float64) {
	var tokens []string
	var s2cells [][][]float64

	rc := &s2.RegionCoverer{MaxLevel: maxLevel, MinLevel: minLevel, MaxCells: maxCells}
	r := s2.Region(p)
	covering := rc.Covering(r)

	for _, c := range covering {
		c1 := s2.CellFromCellID(s2.CellIDFromToken(c.ToToken()))

		var s2cell [][]float64
		for i := 0; i < 4; i++ {
			latlng := s2.LatLngFromPoint(c1.Vertex(i))
			s2cell = append(s2cell, []float64{latlng.Lat.Degrees(), latlng.Lng.Degrees()})
		}

		s2cells = append(s2cells, s2cell)

		tokens = append(tokens, c.ToToken())
	}
	return covering, tokens, s2cells
}

func CoverPoint(p Point, maxLevel int) (s2.Cell, string, [][][]float64) {
	var s2cells [][][]float64

	cid := s2.CellFromLatLng(s2.LatLngFromDegrees(p.Lat, p.Lng)).ID().Parent(maxLevel)
	cell := s2.CellFromCellID(cid)
	token := cid.ToToken()

	var s2cell [][]float64
	for i := 0; i < 4; i++ {
		latLng := s2.LatLngFromPoint(cell.Vertex(i))
		s2cell = append(s2cell, []float64{latLng.Lat.Degrees(), latLng.Lng.Degrees()})
	}
	s2cells = append(s2cells, s2cell)

	return cell, token, s2cells
}
