package controllers_test

import (
	"github.com/gin-gonic/gin"
	"github.com/pantrif/s2-geojson/internal/app/server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

const (
	root = "../../../website"
)

var validJSON = []byte(`
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
    },
    {
      "type": "Feature",
      "properties": {},
      "geometry": {
        "type": "Point",
        "coordinates": [
          -97.86610642636226,
          21.24842223562702
        ]
      }
    }
  ]
}
`)

func TestCover(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := server.NewRouter(root)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/cover", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Result().StatusCode)

	data := url.Values{}
	data.Set("max_level_geojson", "5")
	data.Set("min_level_geojson", "2")
	data.Set("geojson", string(validJSON))

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/cover", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)
}

func TestCheckIntersection(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := server.NewRouter(root)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/check_intersection", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Result().StatusCode)

	data := url.Values{}
	data.Set("radius", "1000")
	data.Set("max_level_circle", "12")
	data.Set("lat", "35.5666")
	data.Set("lng", "23.4444")
	data.Set("tokens", "48761ac,48761b4")
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/check_intersection", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)
}
