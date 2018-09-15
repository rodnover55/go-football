package game

import (
	"image/color"
	"math"
)

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

type Position struct {
	X int
	Y int
}

type Field [HEIGHT][WIDTH]Cell

type Map interface {
	Field() Field
	Position() Position
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

// TODO: Покрыть тестами
func (g Game) CanMove(p Position) bool {
	m := g.GameMap
	ball := m.Position()
	field := m.Field()

	return (math.Abs(float64(p.X - ball.X)) <= 1) && (math.Abs(float64(p.Y - ball.Y)) <= 1) &&
		!field[p.Y][p.X].Filled
}