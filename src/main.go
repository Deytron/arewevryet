package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rowi/arewevryet/conf"
	mi "github.com/rowi/arewevryet/middlewares"
	"github.com/rowi/arewevryet/routes"
	"github.com/rowi/arewevryet/utils"
)

func main() {
	// Load dotenv
	err := godotenv.Load()
	utils.Fatal(err, "Loading .env")

	// Initialisation of router
	conf.InitConfig()

	// config.GetConfig().Router and config.GetConfig().Templates
	r := conf.GetConfig().Router

	// Middlewares
	r.Use(mi.CORSMiddleware(), mi.NormalizeCheckboxInput())

	// Set limit for file upload
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	//Set routes
	routes.SetRoutes(r)

	// Static files
	r.Static("/assets", "/app/assets")

	// Don't trust all proxies
	perr := r.SetTrustedProxies(nil)
	utils.Fatal(perr, "Setting trusted proxies")

	// Setup behind reverse proxy
	runerr := r.Run(":" + os.Getenv("PORT"))
	utils.Fatal(runerr, "Server on port"+os.Getenv("PORT"))

}
