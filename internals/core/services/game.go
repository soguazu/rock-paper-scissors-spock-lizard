package services

import (
	"game/internals/core/ports"
)

// GameService service
type GameService struct {
}

// NewGameService function create a new instance for service
func NewGameService() ports.IGameService {
	return &GameService{}
}

// GetChoices returns all the choices
func (g *GameService) GetChoices() ([]string, error) {
	return nil, nil
}

// GetRandomChoice returns random selected choice
func (g *GameService) GetRandomChoice() (string, error) {
	return "", nil
}

// Play starts the game
func (g *GameService) Play(choice string) (string, error) {
	return choice, nil
}
