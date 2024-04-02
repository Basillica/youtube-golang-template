package main

import (
	"go-template/api"
	"go-template/config/log"
	_ "go-template/docs"
)

var (
	apiVersion = "v1"
)

// @title           Go-Template api documentation
// @version         0.1.0
// @description     The api documentation of the apis of our template.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Tony
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	log.SetLevel(log.LevelInfo)
	log.Infow("starting server...", log.M{
		"success": true, "api-version": apiVersion,
	})
	a := api.New(apiVersion)
	log.Infow("starting running ...", log.M{
		"success": true, "api-version": apiVersion,
	})
	a.Start(a.Engine)
}
