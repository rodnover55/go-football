package game

import "image/color"

type LocalPlayer struct {
	name string
	color color.RGBA
	wins int
}

func (p LocalPlayer) Name() string {
	return p.name
}

func (p LocalPlayer) Color() color.RGBA {
	return p.color
}

func (p LocalPlayer) Wins() int {
	return p.wins
}

func NewPlayer(name string, color color.RGBA) Player {
	return &LocalPlayer{name, color, 0}
}

