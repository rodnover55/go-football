package window

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func New() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Football!!!",
		Bounds: pixel.R(0, 0, 800, 600),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	return win
}