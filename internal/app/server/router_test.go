package server

import (
	"github.com/gin-gonic/gin"
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
    }
  ]
}
`)

func TestRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := NewRouter(root)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"status\":\"ok\"}\n", w.Body.String())
}

func TestCover(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := NewRouter(root)

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

	r := NewRouter(root)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/check_intersection", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Result().StatusCode)

	data := url.Values{}
	data.Set("max_level_geojson", "5")
	data.Set("min_level_geojson", "2")
	data.Set("radius", "1000")
	data.Set("lat", "35.5666")
	data.Set("lng", "23.4444")
	data.Set("geojson", string(validJSON))
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/check_intersection", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)
}
