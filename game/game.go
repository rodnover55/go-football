package game

import (
	"image/color"
	"math"
)

type Game struct {
	GameMap Map
	Players []Player
	activePlayer int
	turns int
	ball Position
	winner Player
}

type Player interface {
	Name() string
	Color() color.RGBA
	Wins() int
	Win()
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
	AddLink(from Position, to Position)
	Copy() Map
}

func NewGame(gameMap Map, players []Player) *Game {
	// TODO: Добавить ошибку, если игроков < 2
	return &Game{
		GameMap: gameMap,
		Players: players,
		activePlayer: 0,
		turns: 0,
		ball: Position{5, 4},
	}
}

func (g Game) ActivePlayer() Player {
	return g.Players[g.activePlayer]
}

func (g Game) canStep(to Position) bool {
	from := g.Ball()

	m := g.GameMap
	fromCell := m.Cell(from)

	return (math.Abs(float64(to.X - from.X)) <= 1) && (math.Abs(float64(to.Y - from.Y)) <= 1) &&
		!fromCell.Linked(m.Cell(to))
}

func (g Game) CanMove(path []Position) bool {
	if g.winner != nil {
		return false
	}

	m := g.GameMap.Copy()

	lastPass := false

	for _, to := range path {
		if !g.canStep(to) || lastPass {
			return false
		}


		m.AddLink(g.ball, to)

		lastPass = len(m.Cell(to).Links()) < 2
		g.ball = to
	}

	return true
}

func (g Game) Ball() Position {
	return g.ball
}

func (g *Game) Move(path []Position) {
	from := g.ball
	m := g.GameMap

	// TODO: Добавить проверку на правильность пути
	for _, to := range path {
		m.AddLink(from, to)
		from = to
	}

	g.activePlayer = (g.activePlayer + 1) % len(g.Players)
	g.ball = path[len(path) - 1]

	g.winner = m.Cell(g.ball).Winner

	if g.winner != nil {
		g.winner.Win()
	}
}

func (g Game) Finished() bool {
	return g.winner != nil
}

func (g Game) Winner() Player {
	return g.winner
}