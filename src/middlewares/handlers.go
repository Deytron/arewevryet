package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
)

func NoRouteHandler(c *gin.Context) {
	data := gin.H{
		"Title": "Page non trouvée",
	}

	ShowPage(c, "404", data)
}

func UnauthorizedHandler(c *gin.Context) {
	data := gin.H{
		"Title": "Non autorisé",
	}

	ShowPage(c, "unauthorized", data)
}

func LogoutHandler(c *gin.Context) {
	c.SetCookie("sessiontoken", "XXX", -1000, "/", os.Getenv("HOST"), true, true)
	c.SetCookie("userid", "XXX", -1000, "/", os.Getenv("HOST"), true, true)
	c.Redirect(302, "/login")
}
