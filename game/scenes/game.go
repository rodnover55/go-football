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

	playersLabels  Labels
	canMove        bool
	cursorPosition game.Position
	path []game.Position
}

func (scene *GameScene) Update() {
	position, err := scene.findNearestPoint(scene.win.MousePosition())
	scene.cursorPosition = position
	scene.canMove = err == nil

	scene.checkActions()

	scene.paintMap()
	scene.paintFilled()
	scene.paintPlayers()
	scene.paintPath()
	scene.paintBall()
}

func (scene *GameScene) checkActions() {
	leftPressed := scene.win.JustPressed(pixelgl.MouseButtonLeft)

	countPasses := len(scene.path)

	switch {
	case leftPressed && scene.canMove:
		scene.move()
		// TODO: Сделать оповещение о невозможности хода
	case leftPressed && (countPasses > 0) && (scene.cursorPosition == scene.path[countPasses - 1]):
		scene.game.Move(scene.path)

		if winner := scene.game.Winner(); winner != nil {
			fmt.Printf("Wins: %s", winner.Name())
		}

		scene.path = []game.Position{}
	}
}

func NewGame(win *pixelgl.Window) *GameScene {
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
		win:           win,
		game:          g,
		playersLabels: labels,
	}
}

func (scene GameScene) paintPlayers() {
	for _, label := range scene.playersLabels {
		label.Draw(scene.win, pixel.IM)
	}
}

func (scene GameScene) paintMap() {
	drawer := imdraw.New(nil)
	drawer.Color = colornames.Gray

	for x := 0; x < game.WIDTH; x += 1 {
		drawLine(
			drawer,
			getCoord(game.Position{X: x, Y: 0}),
			getCoord(game.Position{X: x, Y: game.HEIGHT - 1}),
			false,
		)
	}

	for y := 0; y < game.HEIGHT; y += 1 {
		drawLine(
			drawer,
			getCoord(game.Position{X: 0, Y: y}),
			getCoord(game.Position{X: game.WIDTH - 1, Y: y}),
			false,
		)
	}

	drawer.Draw(scene.win)
}

func (scene GameScene) paintFilled() {
	m := scene.game.GameMap

	drawer := imdraw.New(nil)
	drawer.Color = colornames.White

	for x := 0; x < game.WIDTH; x += 1 {
		for y := 0; y < game.HEIGHT; y += 1 {
			startCell := m.Cell(game.Position{X: x, Y: y})

			for _, cell := range startCell.Links() {
				if (startCell.Position.X < cell.Position.X) ||
					(startCell.Position.Y < cell.Position.Y) {

					drawLine(
						drawer,
						getCoord(startCell.Position),
						getCoord(cell.Position),
						true,
					)
				}
			}
		}
	}

	drawer.Draw(scene.win)
}

func (scene GameScene) paintPath() {
	drawer := imdraw.New(nil)
	drawer.Color = scene.game.ActivePlayer().Color()

	var path []game.Position

	if scene.canMove {
		path = append(scene.path, scene.cursorPosition)
	} else {
		path = scene.path
	}

	from := getCoord(scene.game.Ball())

	for _, p := range path {
		to := getCoord(p)
		drawer.Push(from, to)
		from = to
	}

	drawer.Line(2.0)

	drawer.Draw(scene.win)
}

func (scene GameScene) paintBall() {
	var position game.Position

	countPasses := len(scene.path)

	switch {
	case scene.canMove:
		position = scene.cursorPosition
	case countPasses > 0:
		position = scene.path[countPasses - 1]
	default:
		position = scene.game.Ball()
	}

	drawer := imdraw.New(nil)
	drawer.Color = colornames.White
	drawer.Push(getCoord(position))
	drawer.Circle(2, 4.0)
	drawer.Draw(scene.win)
}

// TODO: Сделать методом GameScene и использовать координаты окна
func getCoord(p game.Position) pixel.Vec {
	return pixel.V(float64(30+40*p.X), float64(570-40*p.Y))
}

func (scene GameScene) findNearestPoint(position pixel.Vec) (p game.Position, err error) {
	p = game.Position{
		X: int(math.Round((position.X - 30.0) / 40.0)),
		Y: int(math.Round((570.0 - position.Y) / 40.0)),
	}

	if (p.X < 0) || (p.Y < 0) ||
		(p.X >= game.WIDTH) || (p.Y >= game.HEIGHT) ||
		!scene.game.CanMove(append(scene.path, p)) {
		err = errors.New("can't step")
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

func (scene *GameScene) move() error {
	path := append(scene.path, scene.cursorPosition)

	if !scene.game.CanMove(path) {
		return errors.New("can't move")
	}

	scene.path = path

	return nil
}
