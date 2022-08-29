package ports

import (
	dto "game/internals/DTO"
	"github.com/gin-gonic/gin"
)

// IGameService defines the interface for game service
type IGameService interface {
	GetChoices() (dto.GetChoices, error)
	GetRandomChoice() (string, error)
	Play(choice string) (*string, error)
	InitializeScoreboard(playerId []string)
	GetScoreboard(playerId string) (*dto.GetScoreboard, error)
	ResetScoreboard(playerId string) error
}

// IGameHandler defines the interface for game handler
type IGameHandler interface {
	GetChoices(c *gin.Context)
	GetRandomChoice(c *gin.Context)
	Play(c *gin.Context)
	GetScoreboard(c *gin.Context)
	ResetScoreboard(c *gin.Context)
}
