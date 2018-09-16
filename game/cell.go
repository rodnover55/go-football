package game

type Cell struct {
	Filled bool
	Winner Player
	Position Position
	links []*Cell
}

func (cell Cell) Linked(l Cell) bool {
	for _, link := range cell.links {
		if link.Position == l.Position {
			return true
		}
	}

	return false
}

func (cell *Cell) addLink(link *Cell) {
	cell.links = append(cell.links, link)
}

func (cell Cell) Links() []Cell {
	var cells []Cell

	for _, link := range cell.links {
		cells = append(cells, *link)
	}

	return cells
}