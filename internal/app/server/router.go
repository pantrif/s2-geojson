package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"s2-geojson/internal/app/controllers"
)

func NewRouter(root string) *gin.Engine {
	health := new(controllers.HealthController)
	p := new(controllers.PolygonController)

	r := gin.Default()
	r.GET("/health", health.Status)
	r.LoadHTMLGlob(root + "/*.html")
	r.Static("/js", root+"/js")
	r.Static("/css", root+"/css")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/cover", p.Cover)
	r.POST("/check_intersection", p.CheckIntersection)

	return r
}
