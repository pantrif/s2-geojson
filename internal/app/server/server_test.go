package server

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServer(t *testing.T) {
	gin.SetMode(gin.TestMode)
	err := Init(root, "3333")
	assert.Error(t, err)
}
