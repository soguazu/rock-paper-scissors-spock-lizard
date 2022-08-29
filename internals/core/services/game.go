package services

import (
	"encoding/json"
	"errors"
	dto "game/internals/DTO"
	"game/internals/core/ports"
	"game/pkg/utils"
	uuid "github.com/satori/go.uuid"
	"strings"
)

var choices = []string{"rock", "paper", "scissors", "spock", "lizard"}

// Player datastore
type Player struct {
	ID    *string
	Score int
}

// Computer datastore
type Computer struct {
	Score int
}

// Scoreboard datastore
type Scoreboard struct {
	Player
	Computer
}

// GameService service
type GameService struct {
	Scoreboard Scoreboard
}

// NewGameService function create a new instance for service
func NewGameService(scoreboard Scoreboard) ports.IGameService {
	return &GameService{Scoreboard: scoreboard}
}

// GetChoices returns all the choices
func (g *GameService) GetChoices() (dto.GetChoices, error) {
	return dto.GetChoices{
		Choices: choices,
		Token:   *g.Scoreboard.Player.ID,
	}, nil
}

// GetRandomChoice returns random selected choice
func (g *GameService) GetRandomChoice() (string, error) {
	randomNumber, err := g.getRandomNumber()
	if err != nil {
		return "", err
	}
	return choices[randomNumber], nil
}

// Play starts the game
func (g *GameService) Play(playerChoice string) (*string, error) {
	computerChoice, err := g.GetRandomChoice()
	if err != nil {
		return nil, err
	}

	result, err := g.getWinner(playerChoice, computerChoice)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// InitializeScoreboard starts the game
func (g *GameService) InitializeScoreboard(Id []string) {
	playerId := g.validatePlayerId(Id)
	if g.Scoreboard.Player.ID != nil && *g.Scoreboard.Player.ID == playerId {
		return
	}

	g.Scoreboard = Scoreboard{
		Player{
			ID: &playerId,
		},
		Computer{},
	}
}

func (g *GameService) validatePlayerId(playerId []string) string {
	if len(playerId) > 0 && g.isUUIDValid(playerId[0]) {
		return playerId[0]
	}
	return uuid.NewV4().String()
}

func (g *GameService) isUUIDValid(id string) bool {
	_, err := uuid.FromString(id)
	if err != nil {
		return false
	}
	return true
}

// getRandomNumber a function that returns random number
func (g *GameService) getRandomNumber() (int, error) {
	Headers := map[string]string{
		"Accept":       "application/json; charset=utf-8",
		"Content-Type": "application/json",
	}
	client := utils.Client{BaseURL: "https://codechallenge.boohma.com", Headers: Headers}

	response, err := client.GET("GET", "random", nil)
	if err != nil {
		return 0, err
	}

	var randomNumber dto.GetRandomNumber
	err = json.Unmarshal(response, &randomNumber)
	if err != nil {
		return 0, err
	}

	for randomNumber.RandomNumber >= 5 {
		randomNumber.RandomNumber = randomNumber.RandomNumber / 2
	}
	return randomNumber.RandomNumber, nil
}

// getRules a function that returns the rules of the game
func (g *GameService) getRules() map[string][]string {
	rules := make(map[string][]string)
	rules["rock"] = []string{"scissors", "lizard"}
	rules["paper"] = []string{"rock", "spock"}
	rules["scissors"] = []string{"paper", "lizard"}
	rules["lizard"] = []string{"paper", "spock"}
	rules["spock"] = []string{"scissors", "rock"}
	return rules
}

// getRules a function that returns the rules of the game
func (g *GameService) getWinner(playerChoice, computerChoice string) (string, error) {

	if playerChoice == "" || !g.contains(choices, strings.ToLower(playerChoice)) {
		return "", errors.New("invalid choice")
	}

	if playerChoice == computerChoice {
		return "It's a tie", nil
	}

	rules := g.getRules()
	if choices, ok := rules[strings.ToLower(playerChoice)]; ok && g.contains(choices, computerChoice) {
		g.updateScoreboard(g.Scoreboard.Player.ID)
		return "player wins", nil
	}

	g.updateScoreboard(nil)
	return "computer wins", nil
}

func (g *GameService) contains(choices []string, searchTerm string) bool {
	for _, choice := range choices {
		if choice == searchTerm {
			return true
		}
	}
	return false
}

func (g *GameService) updateScoreboard(player *string) {
	if player != nil {
		g.Scoreboard.Player.Score++
	} else {
		g.Scoreboard.Computer.Score++
	}
}

// GetScoreboard persist the score of the game
func (g *GameService) GetScoreboard(playerId string) (*dto.GetScoreboard, error) {
	if g.Scoreboard.Player.ID == nil || *g.Scoreboard.Player.ID != playerId {
		return nil, errors.New("invalid player id")
	}
	return &dto.GetScoreboard{
		PlayerScore:   g.Scoreboard.Player.Score,
		ComputerScore: g.Scoreboard.Computer.Score,
	}, nil
}

// ResetScoreboard persist the score of the game
func (g *GameService) ResetScoreboard(playerId string) error {
	if g.Scoreboard.Player.ID == nil || *g.Scoreboard.Player.ID != playerId {
		return errors.New("invalid player id")
	}

	g.Scoreboard.Player.Score = 0
	g.Scoreboard.Computer.Score = 0
	return nil
}
