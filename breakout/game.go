package breakout

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)


const (
	ScreenWidth  = 480
	ScreenHeight = 360
)

const (
	initBarX = ScreenWidth/2
	initBarY = 4*ScreenHeight/5
	barWidth = 100
	barHeight = 20

	initBallX = ScreenWidth/2
	initBallY = 3*ScreenHeight/5
	ballRadius = 4
)

var (
	barSpeed = 8
	ballVelocity = 4
)


type Game struct{
	bar *Bar
	ball *Ball
	block *Block
}

func NewGame() (*Game, error) {
	g := &Game{}
	var err error
	// generate a new obj, bar, ball, block
	g.bar = &Bar{
		x: initBarX,
		y: initBarY,
		speed: barSpeed,
		width: barWidth,
		height: barHeight,
		image: ebiten.NewImage(barWidth, barHeight),
	}
	g.bar.image.Fill(color.White)

	g.ball = &Ball{
		x: initBallX,
		y: initBallY,
		velocityX: ballVelocity,
		velocityY: ballVelocity,
		radius: ballRadius,
		image: ebiten.NewImage(ballRadius, ballRadius),
	}
	g.ball.image.Fill(color.White)

	return g, err
}

func (g *Game) Update() error {
	g.bar.Update()
	g.ball.Update(g.bar)
	// g.block.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ball.Draw(screen)
	g.bar.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}