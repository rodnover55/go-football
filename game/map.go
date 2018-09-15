package game


type SimpleMap struct {
	field Field
	ballPosition Position
}

func (m SimpleMap) Position() Position {
	return m.ballPosition
}

func NewMap(players []Player) Map {
	f := Cell{true, nil}
	o := Cell{false, nil}

	l := Cell{false, players[0]}
	r := Cell{false, players[1]}

	field := Field{
		{f, f, f, f, f, f, f, f, f, f, f},
		{f, o, o, f, o, f, o, f, o, o, f},
		{f, o, o, f, o, f, o, f, o, o, f},
		{f, o, o, f, o, f, o, f, o, o, f},
		{r, o, o, f, o, f, o, f, o, o, l},
		{f, o, o, f, o, f, o, f, o, o, f},
		{f, o, o, f, o, f, o, f, o, o, f},
		{f, o, o, f, o, f, o, f, o, o, f},
		{f, f, f, f, f, f, f, f, f, f, f},
	}

	return &SimpleMap{
		field: field,
		ballPosition: Position{5, 4},
	}
}


func (m SimpleMap) Field() Field {
	return m.field
}

