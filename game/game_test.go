package game

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/image/colornames"
	"testing"
)

func TestStepSuccess(t *testing.T) {
	players := []Player{
		NewPlayer("Player 1", colornames.Green),
		NewPlayer("Player 2", colornames.Red),
	}

	g := NewGame(
		NewMap(players),
		players,
	)

	g.ball = Position{2, 0}
	assert.True(t, g.canStep(Position{3, 1}))
}

func TestStepFail(t *testing.T) {
	examples := map[string]struct{
		from Position
		to Position
	} {
		"hasLink": {
			from: Position{2, 0},
			to: Position{3, 0},
		},
		"far": {
			from: Position{2, 0},
			to: Position{2, 2},
		},
	}

	for name, example := range examples {
		t.Run(name, func (t *testing.T) {
			players := []Player{
				NewPlayer("Player 1", colornames.Green),
				NewPlayer("Player 2", colornames.Red),
			}

			g := NewGame(
				NewMap(players),
				players,
			)

			g.ball = example.from
			assert.False(t, g.canStep(example.to))
		})
	}
}

func TestMapRewrite(t *testing.T) {
	players := []Player{
		NewPlayer("Player 1", colornames.Green),
		NewPlayer("Player 2", colornames.Red),
	}

	g := NewGame(
		NewMap(players),
		players,
	)

	g.ball = Position{5, 4}

	to := Position{6, 4}
	g.CanMove([]Position{to})

	ball := g.GameMap.Cell(g.ball)
	assert.False(t, ball.Linked(g.GameMap.Cell(to)))
}

func TestMoveSinglePass(t *testing.T) {
	players := []Player{
		NewPlayer("Player 1", colornames.Green),
		NewPlayer("Player 2", colornames.Red),
	}

	g := NewGame(
		NewMap(players),
		players,
	)
	ball := Position{6, 4}
	g.GameMap.AddLink(g.ball, ball)

	g.ball = ball
	assert.True(t, g.CanMove([]Position{{6, 3}}))
}

func TestMoveMultiplePasses(t *testing.T) {
	players := []Player{
		NewPlayer("Player 1", colornames.Green),
		NewPlayer("Player 2", colornames.Red),
	}

	g := NewGame(
		NewMap(players),
		players,
	)
	ball := Position{6, 3}
	g.GameMap.AddLink(g.ball, ball)

	g.ball = ball
	g.Move([]Position{{7, 3}, {8, 2}})

	assert.Equal(t, 2, len(g.GameMap.Cell(ball).links))
}

func TestMoveFalse(t *testing.T) {
	examples := map[string]struct{
		from Position
		path []Position
	} {
		"excessMove": {
			from: Position{5, 4},
			path: []Position{
				{6, 4},
				{6, 3},
			},
		},
	}

	for name, example := range examples {
		t.Run(name, func (t *testing.T) {
			players := []Player{
				NewPlayer("Player 1", colornames.Green),
				NewPlayer("Player 2", colornames.Red),
			}

			g := NewGame(
				NewMap(players),
				players,
			)

			g.ball = example.from
			assert.False(t, g.CanMove(example.path))
		})
	}
}

