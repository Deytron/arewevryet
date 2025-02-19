package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rowi/arewevryet/conf"
	"github.com/rowi/arewevryet/globals"
	"github.com/rowi/arewevryet/utils"
)

func ShowPage(c *gin.Context, page string, data gin.H) {
	var t = conf.GetConfig().Templates

	// Any message that is set in code will be displayed on the page
	successMessage, errorMessage, infoMessage := globals.GetMessages()

	if successMessage != "" || errorMessage != "" || infoMessage != "" {
		data["SuccessMessage"] = successMessage
		data["ErrorMessage"] = errorMessage
		data["InfoMessage"] = infoMessage
	}
	
	err := t.ExecuteTemplate(c.Writer, page+".html", data)
	utils.Fatal(err, "Executing requested template")

	globals.ClearMessages()
}
