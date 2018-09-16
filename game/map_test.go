package game

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/image/colornames"
	"testing"
)

func TestCreate(t *testing.T) {
	players := []Player{
		NewPlayer("Player 1", colornames.Green),
		NewPlayer("Player 2", colornames.Red),
	}

	m := NewMap(players)


	assert.Equal(t, players[0], m.Cell(Position{10, 4}).Winner)
	assert.Equal(t, players[1], m.Cell(Position{0, 4}).Winner)

	p := Position{3, 0}
	cell := m.Cell(p)

	assert.Equal(t, p, cell.Position)

	assert.Equal(t, 3, len(cell.links))
	assert.True(t, cell.Linked(m.Cell(Position{3, 1})))
	assert.True(t, cell.Linked(m.Cell(Position{4, 0})))
	assert.True(t, cell.Linked(m.Cell(Position{2, 0})))
}
