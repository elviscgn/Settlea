package base_game

import (
	"settlea/internal/game/data"
	"settlea/pkg/uid"
)

// catan game
type Game struct {
	GameID  string
	Players []data.Player
	// seed         int
	DiscardLimit int
	vpsToWin     int
	Board        *data.SettleaMap
	Ports        map[string]data.PortData
}

func (g *Game) InitGame(players_no int, discardLimit int, vpsToWin int, style string, ports map[string]data.PortData) *Game {
	players := make([]data.Player, players_no)
	settleaMap := &data.SettleaMap{}
	gameID := uid.GenerateUniqueID(13)

	return &Game{
		gameID,
		players,
		discardLimit,
		vpsToWin,
		settleaMap.NewMap(style),
		ports,
	}
}
