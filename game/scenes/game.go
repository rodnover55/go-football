package scenes

import (
	"errors"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"github.com/rodnover55/go-football/game"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/gofont/goregular"
	"math"
)

type Labels map[game.Player]*text.Text

type GameScene struct {
	game *game.Game
	win  *pixelgl.Window

	playersLabels Labels
	canMove bool
	cursorPosition game.Position
}

func (scene *GameScene) Update() {
	position, err := scene.findNearestPoint(scene.win.MousePosition())

	scene.cursorPosition = position
	scene.canMove = err == nil

	scene.paintMap()
	scene.paintPlayers()

	if !scene.canMove {
		scene.paintPath()
	}
}

func MakeGame(win *pixelgl.Window) *GameScene {
	ttf, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}

	face := truetype.NewFace(ttf, &truetype.Options{
		Size: 14,
	})

	players := []game.Player{
		game.NewPlayer("Player 1", colornames.Green),
		game.NewPlayer("Player 2", colornames.Red),
	}

	g := game.NewGame(
		game.NewMap(players),
		players,
	)

	labels := Labels{}

	labelHeight := 560.0

	for _, player := range players {
		label := text.New(pixel.V(450, labelHeight), text.NewAtlas(face, text.ASCII))

		fmt.Fprintf(label, "%s (Wins: %d)", player.Name(), player.Wins())
		label.Color = player.Color()

		labels[player] = label
		labelHeight -= 25
	}

	return &GameScene{
		win:  win,
		game: g,
		playersLabels: labels,
	}
}

func (scene GameScene) paintPlayers() {
	for _, label := range scene.playersLabels {
		label.Draw(scene.win, pixel.IM)
	}
}

func (scene GameScene) paintMap() {
	field := scene.game.GameMap.Field()

	drawer := imdraw.New(nil)
	drawer.Color = colornames.White

	for x := 0; x < game.WIDTH; x += 1 {
		for y := 0; y < game.HEIGHT; y += 1 {
			startCell := field[y][x]

			if y < game.HEIGHT-1 {
				drawLine(
					drawer,
					getCoord(game.Position{x, y}),
					getCoord(game.Position{x, y+1}),
					startCell.Filled && field[y+1][x].Filled,
				)
			}

			if x < game.WIDTH-1 {
				drawLine(
					drawer,
					getCoord(game.Position{x, y}),
					getCoord(game.Position{x+1, y}),
					startCell.Filled && field[y][x+1].Filled,
				)
			}
		}
	}

	drawer.Color = colornames.White
	drawer.Push(getCoord(scene.game.GameMap.Position()))
	drawer.Circle(2, 4.0)

	drawer.Draw(scene.win)
}

func (scene GameScene) paintPath() {
	ballPosition := getCoord(scene.game.GameMap.Position())
	position := getCoord(scene.cursorPosition)

	drawer := imdraw.New(nil)
	drawer.Color = colornames.Gray

	drawer.Push(ballPosition, position)
	drawer.Line(1)

	drawer.Color = scene.game.ActivePlayer().Color()
	drawer.Push(position)
	drawer.Circle(2, 4.0)

	drawer.Draw(scene.win)
}

// TODO: Сделать методом GameScene и использовать координаты окна
func getCoord(p game.Position) pixel.Vec {
	return pixel.V(float64(30+40*p.X), float64(570 - 40*p.Y))
}

func (scene GameScene) findNearestPoint(position pixel.Vec) (p game.Position, err error) {
	p = game.Position{
		X: int(math.Round((position.X - 30.0) / 40.0)),
		Y: int(math.Round((570.0 - position.Y) / 40.0)),
	}

	if scene.game.CanMove(p) {
		err = errors.New("can't move")
	}

	return p, err
}

func drawLine(drawer *imdraw.IMDraw, from pixel.Vec, to pixel.Vec, filled bool) {
	thickness := 1.0
	drawer.Color = colornames.Gray

	if filled {
		drawer.Color = colornames.White
		thickness = 2.0
	}

	drawer.Push(from, to)

	drawer.Line(thickness)
}