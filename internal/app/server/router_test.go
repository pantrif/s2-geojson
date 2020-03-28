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

func TestRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := NewRouter(root)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"status\":\"ok\"}\n", w.Body.String())
}
