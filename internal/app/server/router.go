package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"s2-geojson/internal/app/controllers"
)

func NewRouter() *gin.Engine {
	health := new(controllers.HealthController)
	p := new(controllers.PolygonController)

	r := gin.Default()
	r.GET("/health", health.Status)
	r.LoadHTMLGlob("website/*.html")
	r.Static("/js", "./website/js")
	r.Static("/css", "./website/css")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/cover", p.Cover)
	r.POST("/check_intersection", p.CheckIntersection)

	return r
}
