package server

import (
	dto "game/internals/DTO"
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
		gameService = services.NewGameService(services.Scoreboard{Score: []dto.PlayResponse{}})
		gameHandler = handlers.NewGameHandler(gameService, "Game")
	)

	group := ginRoutes.GROUP("")
	game := group.Group("/")
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
