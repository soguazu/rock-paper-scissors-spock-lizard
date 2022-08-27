package main

import (
	"game/cmd/server"
	_ "game/docs"
)

// @title Evea Core Business Swagger API
// @version 1.0
// @description Evea Core Business Swagger API.
// @termsOfService http://swagger.io/terms/

// @contact.name Evea Team API Support
// @contact.email info@evea.com

// @license.name MIT
// @license.url https://github.com/sguazu

// @BasePath /v1
func main() {
	server.Injection()
}
