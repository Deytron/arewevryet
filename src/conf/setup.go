package conf

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/rowi/arewevryet/globals"
)

var appConfig *AppConfig

type AppConfig struct {
	Templates *template.Template
	Router    *gin.Engine
}

// Simply disable logging if is prod
func InitConfig() {
	if os.Getenv("HOST") == "bat.gbna-sante.fr" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	appConfig = &AppConfig{
		Templates: ParseTemplates(),
		Router:    gin.Default(),
	}

	appConfig.Router.Use(gin.Logger(), gin.Recovery())
}

func GetConfig() *AppConfig {
	return appConfig
}

func ParseTemplates() *template.Template {
	templ := template.New("")

	err := filepath.Walk(globals.HTMLPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	return templ
}
