package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rowi/arewevryet/home"
	mi "github.com/rowi/arewevryet/middlewares"
)

func SetRoutes(r *gin.Engine) {
	// Routes declaration

	// 404
	r.NoRoute(mi.NoRouteHandler)

	// GET Routes
	r.GET("/", home.HomeHandler)
	r.POST("/", home.HomeHandler)
}
