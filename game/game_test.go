package game

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/image/colornames"
	"testing"
)

func TestMovingSuccess(t *testing.T) {
	players := []Player{
		NewPlayer("Player 1", colornames.Green),
		NewPlayer("Player 2", colornames.Red),
	}

	g := NewGame(
		NewMap(players),
		players,
	)

	g.ball = Position{2, 0}

	assert.True(t, g.CanMove(Position{3, 1}))
}