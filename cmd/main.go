package main

import (
	"game/cmd/server"
	_ "game/docs"
)

// @title Rock Paper Scissors Spock Lizard Swagger API
// @version 1.0
// @description Rock Paper Scissors Spock Lizard Swagger API.
// @termsOfService http://swagger.io/terms/

// @contact.name Team API Support
// @contact.email cuguazu@gmail.com

// @license.name MIT
// @license.url https://github.com/sguazu

// @BasePath /
func main() {
	server.Injection()
}
