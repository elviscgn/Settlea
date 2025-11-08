package handlers

import (
	"net/http"

	base_game "settlea/internal/game/base"
	"settlea/internal/game/data"
	"settlea/pkg/bestagons/edge"

	"github.com/gin-gonic/gin"
)

type BoardResponse struct {
	Tiles      []*data.Tile             `json:"tiles"`
	Corners    []*data.Corner           `json:"corners"`
	Edges      []*edge.Edge             `json:"edges"`
	Ports      map[string]data.PortData `json:"ports"`
	Iterations int                      `json:"iterations"`
	Duration   string                   `json:"duration"`
}

func NewGame(c *gin.Context) {

	// init game
	new_game := base_game.Game{}
	initialised_game := new_game.InitGame(2, 2, 1, "base", map[string]data.PortData{})

	// get tiles, vertices, edges
	map_tiles := initialised_game.Board.Tiles
	map_corners := initialised_game.Board.Corners
	map_edges := initialised_game.Board.Edges

	// validate tiles & allocate
	_, iterations, duration := data.StartValidation(map_tiles)

	// gen random ports
	ports := data.GeneratePorts(9)

	// send everything to client
	response := BoardResponse{
		Tiles:      map_tiles, // assuming this is now populated correctly
		Corners:    map_corners,
		Edges:      map_edges,
		Ports:      ports,
		Iterations: iterations,
		Duration:   duration.String(),
	}

	c.JSON(http.StatusOK, response)
}
