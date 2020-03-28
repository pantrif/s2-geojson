package geo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var validJson = []byte(`
{
  "type": "FeatureCollection",
  "features": [
    {
      "type": "Feature",
      "properties": {},
      "geometry": {
        "type": "Polygon",
        "coordinates": [
          [
            [
              -71.98116712385131,
              40.58058466412764
            ],
            [
              -127.79249821831927,
              48.63290858589535
            ],
            [
              -90.17478214204795,
              25.64152637306577
            ],
            [
              -71.98116712385131,
              40.58058466412764
            ]
          ]
        ]
      }
    }
  ]
}
`)

func TestDecodeGeoJSON(t *testing.T) {
	r, err := DecodeGeoJSON(validJson)
	assert.NoError(t, err)
	assert.True(t, r[0].Geometry.IsPolygon())

	_, err = DecodeGeoJSON([]byte("foo"))
	assert.Error(t, err)
}

func TestPointsToPolygon(t *testing.T) {
	r, err := DecodeGeoJSON(validJson)
	assert.NoError(t, err)
	assert.True(t, r[0].Geometry.IsPolygon())

	p := PointsToPolygon(r[0].Geometry.Polygon[0])
	assert.Equal(t, 4, p.NumEdges())
}

func TestCoverPolygon(t *testing.T) {
	f, _ := DecodeGeoJSON(validJson)
	p := PointsToPolygon(f[0].Geometry.Polygon[0])

	u, tk, c := CoverPolygon(p, 4, 1)
	assert.True(t, u.IsValid())
	assert.Equal(t, 22, len(tk))
	assert.Equal(t, 4, len(c[0]))

}

func TestCoverPoint(t *testing.T) {
	cell, token, edges := CoverPoint(Point{Lat: 38.34, Lng: 34.34}, 1)
	assert.Equal(t, "14", cell.ID().ToToken())
	assert.Equal(t, "14", token)
	assert.Equal(t, 4, len(edges[0]))

}
