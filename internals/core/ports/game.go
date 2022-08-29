package ports

import (
	dto "game/internals/DTO"
	"github.com/gin-gonic/gin"
)

// IGameService defines the interface for game service
type IGameService interface {
	GetChoices() ([]dto.Choices, error)
	GetRandomChoice() (*dto.Choices, error)
	Play(choice int) (*dto.PlayResponse, error)
	GetScoreboard() []dto.PlayResponse
	ResetScoreboard()
}

// IGameHandler defines the interface for game handler
type IGameHandler interface {
	GetChoices(c *gin.Context)
	GetRandomChoice(c *gin.Context)
	Play(c *gin.Context)
	GetScoreboard(c *gin.Context)
	ResetScoreboard(c *gin.Context)
}
