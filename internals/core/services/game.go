package services

import (
	"encoding/json"
	"errors"
	"fmt"
	dto "game/internals/DTO"
	"game/internals/core/ports"
	"game/pkg/utils"
	uuid "github.com/satori/go.uuid"
)

var choices = []dto.Choices{
	{Name: "rock", Id: 1},
	{Name: "paper", Id: 2},
	{Name: "scissors", Id: 3},
	{Name: "spock", Id: 4},
	{Name: "lizard", Id: 5},
}

// Scoreboard datastore
type Scoreboard struct {
	Score []dto.PlayResponse
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
func (g *GameService) GetChoices() ([]dto.Choices, error) {
	return choices, nil
}

// GetRandomChoice returns random selected choice
func (g *GameService) GetRandomChoice() (*dto.Choices, error) {
	randomNumber, err := g.getRandomNumber()
	if err != nil {
		return nil, err
	}

	return &choices[randomNumber-1], nil
}

// Play starts the game
func (g *GameService) Play(playerChoice int) (*dto.PlayResponse, error) {
	computerChoice, err := g.GetRandomChoice()
	if err != nil {
		return nil, err
	}

	result, err := g.getWinner(playerChoice, computerChoice.Id)
	if err != nil {
		return nil, err
	}

	score := dto.PlayResponse{
		Results: result,
		Player:  playerChoice,
	}
	g.updateScoreboard(score)

	return &score, nil
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

	for randomNumber.RandomNumber > 5 {
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
func (g *GameService) getWinner(playerChoice, computerChoice int) (string, error) {
	fmt.Println(playerChoice, "****")
	if playerChoice < 1 || playerChoice > 5 {
		return "", errors.New("invalid choice")
	}

	if computerChoice < 0 {
		computerChoice = 1
	}

	if playerChoice == computerChoice {
		return "tie", nil
	}

	rules := g.getRules()
	if selected, ok := rules[choices[playerChoice-1].Name]; ok && g.containsSelected(selected, choices[computerChoice-1].Name) {
		return "wins", nil
	}

	return "lose", nil
}

func (g *GameService) contains(choices []dto.Choices, searchTerm int) bool {
	for _, choice := range choices {
		if choice.Id == searchTerm {
			return true
		}
	}
	return false
}

func (g *GameService) containsSelected(choices []string, searchTerm string) bool {
	for _, choice := range choices {
		if choice == searchTerm {
			return true
		}
	}
	return false
}

func (g *GameService) updateScoreboard(scoreboard dto.PlayResponse) {
	g.Scoreboard.Score = append(g.Scoreboard.Score, scoreboard)
}

// GetScoreboard persist the score of the game
func (g *GameService) GetScoreboard() []dto.PlayResponse {
	if len(g.Scoreboard.Score) > 10 {
		return g.Scoreboard.Score[0:10]
	}
	return g.Scoreboard.Score
}

// ResetScoreboard persist the score of the game
func (g *GameService) ResetScoreboard() {
	g.Scoreboard.Score = []dto.PlayResponse{}
}
