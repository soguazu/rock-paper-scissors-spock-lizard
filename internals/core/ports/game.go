package ports

import (
	"github.com/gin-gonic/gin"
)

// IGameService defines the interface for game service
type IGameService interface {
	GetChoices() ([]string, error)
	GetRandomChoice() (string, error)
	Play(choice string) (string, error)
}

// IGameHandler defines the interface for game handler
type IGameHandler interface {
	GetChoices(c *gin.Context)
	GetRandomChoice(c *gin.Context)
	Play(c *gin.Context)
}
