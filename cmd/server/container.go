package server

import (
	"game/internals/core/services"
	"game/internals/handlers"
	"game/pkg/config"
	"github.com/gin-gonic/gin"
	"log"
)

// Injection inject all dependencies
func Injection() {

	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	var (
		ginRoutes   = NewGinRouter(gin.Default())
		gameService = services.NewGameService(services.Scoreboard{})
		gameHandler = handlers.NewGameHandler(gameService, "Game")
	)

	v1 := ginRoutes.GROUP("v1")
	game := v1.Group("/")
	game.GET("/choices", gameHandler.GetChoices)
	game.GET("/choice", gameHandler.GetRandomChoice)
	game.GET("/scoreboard", gameHandler.GetScoreboard)
	game.GET("/reset-scoreboard", gameHandler.ResetScoreboard)
	game.POST("/play", gameHandler.Play)

	err = ginRoutes.SERVE()

	if err != nil {
		return
	}

}
