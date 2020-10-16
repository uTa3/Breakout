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

	initBollX = ScreenWidth/2
	initBollY = 3*ScreenHeight/5
	bollRadius = 5
)

var (
	barSpeed = 8
	bollVelocity = 2
)


type Game struct{
	bar *Bar
	boll *Boll
	block *Block
}

func NewGame() (*Game, error) {
	g := &Game{}
	var err error
	// generate a new obj, bar, boll, block
	g.bar = &Bar{
		x: initBarX,
		y: initBarY,
		speed: barSpeed,
		width: barWidth,
		height: barHeight,
		image: ebiten.NewImage(barWidth, barHeight),
	}
	g.bar.image.Fill(color.White)

	g.boll = &Boll{
		x: initBollX,
		y: initBollY,
		velocityX: bollVelocity,
		velocityY: bollVelocity,
		radius: bollRadius,
		image: ebiten.NewImage(bollRadius, bollRadius),
	}
	g.boll.image.Fill(color.White)

	return g, err
}

func (g *Game) Update() error {
	g.bar.Update()
	g.boll.Update(g.bar)
	// g.block.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.boll.Draw(screen)
	g.bar.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}