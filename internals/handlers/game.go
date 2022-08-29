package handlers

import (
	"fmt"
	"game/internals/DTO"
	"net/http"

	"game/internals/DTO/types"
	"game/internals/core/ports"
	"game/pkg/utils"
	"github.com/gin-gonic/gin"
)

// GameHandler handler
type GameHandler struct {
	GameService ports.IGameService
	handlerName string
}

var (
	result  utils.Result
	message types.Messages
)

// NewGameHandler function creates a new instance for game handler
func NewGameHandler(cs ports.IGameService, n string) ports.IGameHandler {
	return &GameHandler{
		GameService: cs,
		handlerName: n,
	}
}

// GetChoices godoc
// @Summary      Get choices
// @Description  get all choices
// @Tags         game
// @Accept       json
// @Produce      json
// @Param PlayId header string false "Player ID"
// @Success      200  {object}  dto.GetChoices
// @Failure      400  {object}  dto.Error
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /choices [get]
func (g *GameHandler) GetChoices(c *gin.Context) {
	g.GameService.InitializeScoreboard(c.Request.Header["Playid"])
	choices, err := g.GameService.GetChoices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(choices, message.GetResponseMessage(g.handlerName, types.OKAY)))
}

// GetRandomChoice godoc
// @Summary      Get random choice
// @Description  gets random generated choice
// @Tags         game
// @Accept       json
// @Produce      json
// @Param PlayId header string false "Player ID"
// @Success      200  {object}  dto.GetChoice
// @Failure      500  {object}  dto.Error
// @Router       /choice [get]
func (g *GameHandler) GetRandomChoice(c *gin.Context) {
	choice, err := g.GameService.GetRandomChoice()
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(choice, message.GetResponseMessage(g.handlerName, types.OKAY)))
}

// Play godoc
// @Summary      Play
// @Description  this plays the game rock, paper, scissors, lizard, spour
// @Tags         game
// @Accept       json
// @Produce      json
// @Param PlayId header string false "Player ID"
// @Param game body dto.PlayRequest true "play"
// @Success      200  {object}  dto.PlayResponse
// @Failure      400  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /play [post]
func (g *GameHandler) Play(c *gin.Context) {
	g.GameService.InitializeScoreboard(c.Request.Header["Playid"])
	var body dto.PlayRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	winner, err := g.GameService.Play(body.Choice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.ReturnSuccessResult(winner, message.GetResponseMessage(g.handlerName, types.OKAY)))
}

// GetScoreboard godoc
// @Summary      GetScoreboard
// @Description  Returns the scoreboard of the current play
// @Tags         game
// @Accept       json
// @Produce      json
// @Param PlayId header string false "Player ID"
// @Success      200  {object}  dto.GetResponse
// @Failure      400  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /scoreboard [get]
func (g *GameHandler) GetScoreboard(c *gin.Context) {

	if len(c.Request.Header["Playid"]) < 1 {
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult("invalid player id"))
		return
	}

	scoreboard, err := g.GameService.GetScoreboard((c.Request.Header["Playid"])[0])
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.ReturnSuccessResult(scoreboard, message.GetResponseMessage(fmt.Sprintf("%v score", g.handlerName), types.OKAY)))
}

// ResetScoreboard godoc
// @Summary      ResetScoreboard
// @Description  Resets the scoreboard
// @Tags         game
// @Accept       json
// @Produce      json
// @Param PlayId header string false "Player ID"
// @Success      200  {object}  dto.GetResponse
// @Failure      400  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /reset-scoreboard [get]
func (g *GameHandler) ResetScoreboard(c *gin.Context) {
	if len(c.Request.Header["Playid"]) < 1 {
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult("invalid player id"))
		return
	}

	err := g.GameService.ResetScoreboard(c.Request.Header["Playid"][0])
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.ReturnBasicResult(message.GetResponseMessage(fmt.Sprintf("%v score", g.handlerName), types.UPDATED)))
}
