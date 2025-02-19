package home

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rowi/arewevryet/globals"
	mi "github.com/rowi/arewevryet/middlewares"
	"github.com/rowi/arewevryet/models"
	"github.com/rowi/arewevryet/utils"
)

func HomeHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		var user models.Account
		var userconfig models.UserConfig
		c.ShouldBind(&user)

		if len(user.Name) < 1 || len(user.Password) < 1 {
			globals.SetError("Aucun nom d'utilisateur ou mot de passe n'a été rentré")
			mi.ShowPage(c, "index", gin.H{})
			return
		}

		// Create the user with password
		err, resp := utils.APICall("POST", os.Getenv("JELLYFIN_URL") + "Users/New", os.Getenv("JELLYFIN_API"), &userconfig, user)

		// Updates user policy to hide Sev folder
		enabledfolders := []string{
			"ca0de50d2c11073f53df7c82dc3fe2a4",
			"bdf38141c3a366eb1a2a8240d2e65e68",
			"9d7ad6afe9afa2dab1a2f6e00ad28fa6",
			"db4c1708cbb5dd1676284a40f2950aba",
			"bebdce85c5b682ddbce0412f41cff060",
			"d565273fd114d77bdf349a2896867069",
		}
		userconfig.EnableAllFolders = false
		userconfig.EnabledFolders = enabledfolders
		userconfig.AuthenticationProviderId = "Jellyfin.Server.Implementations.Users.DefaultAuthenticationProvider"
		userconfig.PasswordResetProviderId = "Jellyfin.Server.Implementations.Users.DefaultPasswordResetProvider"

		err2, resp2 := utils.APICall("POST", os.Getenv("JELLYFIN_URL") + "Users/" + userconfig.Id + "/Policy", os.Getenv("JELLYFIN_API"), nil, userconfig)

		if err || err2 {
			globals.SetError("Erreur lors de la création de votre compte : " + resp + resp2)
		} else {
			globals.SetSuccess("Votre compte " + user.Name + " a été créé avec succès !")
			globals.SetInfo("Accédez à Jellyfin à partir du lien stream.row-it.net, ou téléchargez l'application sur vos appareils")
		}
	}

	data := gin.H{}

	mi.ShowPage(c, "index", data)
}