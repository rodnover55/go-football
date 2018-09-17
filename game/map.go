package game

import "math"

type SimpleMap struct {
	field Field
}

func (m *SimpleMap) Copy() Map {
	return &SimpleMap{
		field: m.field,
	}
}

func (m *SimpleMap) AddLink(from Position, to Position) {
	m.cell(from).addLink(m.cell(to))
}

func (m SimpleMap) Cell(p Position) Cell {
	return m.field[p.Y][p.X]
}

func (m *SimpleMap) cell(p Position) *Cell {
	return &m.field[p.Y][p.X]
}

func NewMap(players []Player) Map {
	f := Cell{Filled: true}
	o := Cell{Filled: false}

	field := Field{
		{f, f, f, f, f, f, f, f, f, f, f},
		{f, o, o, f, o, f, o, f, o, o, f},
		{f, o, o, f, o, f, o, f, o, o, f},
		{f, o, o, f, o, f, o, f, o, o, f},
		{o, o, o, f, o, f, o, f, o, o, o},
		{f, o, o, f, o, f, o, f, o, o, f},
		{f, o, o, f, o, f, o, f, o, o, f},
		{f, o, o, f, o, f, o, f, o, o, f},
		{f, f, f, f, f, f, f, f, f, f, f},
	}

	m := &SimpleMap{
		field: field,
	}

	middleHeight := int(math.Trunc(HEIGHT / 2))
	m.cell(Position{WIDTH - 1, middleHeight}).Winner = players[0]
	m.cell(Position{0, middleHeight}).Winner = players[1]

	for x := 0; x < WIDTH; x += 1 {
		for y := 0; y < HEIGHT; y += 1 {
			position := Position{X: x, Y: y}
			cell := m.cell(position)

			cell.Position = position

			if y < HEIGHT-1 {
				toCell := m.cell(Position{X: x, Y: y + 1})

				if cell.Filled && toCell.Filled {
					cell.addLink(toCell)
				}
			}

			if x < WIDTH-1 {
				toCell := m.cell(Position{X: x + 1, Y: y})

				if cell.Filled && toCell.Filled {
					cell.addLink(toCell)
				}
			}
		}
	}

	return m
}
