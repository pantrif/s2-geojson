package server

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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
