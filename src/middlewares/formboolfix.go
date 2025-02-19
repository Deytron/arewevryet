package middlewares

import "github.com/gin-gonic/gin"

func NormalizeCheckboxInput() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.ParseForm()
		for key, values := range c.Request.Form {
			if values[0] == "on" {
				c.Request.Form.Set(key, "true")
			}
		}
		c.Next()
	}
}
