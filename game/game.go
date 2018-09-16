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
	ball Position
}

type Player interface {
	Name() string
	Color() color.RGBA
	Wins() int
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
	Cell(Position) Cell
}

func NewGame(gameMap Map, players []Player) *Game {
	// TODO: Добавить ошибку, если игроков < 2
	return &Game{
		GameMap: gameMap,
		Players: players,
		activePlayer: players[0],
		turns: 0,
		ball: Position{5, 4},
	}
}

func (g Game) ActivePlayer() Player {
	return g.activePlayer
}

// TODO: Покрыть тестами
func (g Game) CanMove(p Position) bool {
	m := g.GameMap
	ball := g.ball
	ballCell := m.Cell(ball)

	return (math.Abs(float64(p.X - ball.X)) <= 1) && (math.Abs(float64(p.Y - ball.Y)) <= 1) &&
		!ballCell.Linked(m.Cell(p))
}

func (g Game) Ball() Position {
	return g.ball
}