package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := HealthController{}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	h.Status(c)
	assert.Equal(t, 200, w.Result().StatusCode)
}
