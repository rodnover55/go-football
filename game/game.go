package game

import "image/color"

type Game struct {
	GameMap Map
	Players []Player
	activePlayer Player
	turns int
}

type Player interface {
	Name() string
	Color() color.RGBA
	Wins() int
}

type Cell struct {
	Filled bool
	Winner Player
}

const (
	WIDTH = 11
	HEIGHT = 9
)
type Field [HEIGHT][WIDTH]Cell

type Map interface {
	Field() Field
	Position() (int, int)
}

func NewGame(gameMap Map, players []Player) *Game {
	// TODO: Добавить ошибку, если игроков < 2
	return &Game{
		GameMap: gameMap,
		Players: players,
		activePlayer: players[0],
		turns: 0,
	}
}

func (g Game) ActivePlayer() Player {
	return g.activePlayer
}