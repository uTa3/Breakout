package breakout

import (
	_ "image/png"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bar struct {
	x, y int
	speed int
	width, height int
	image *ebiten.Image
}

func (bar *Bar) Update()  {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		bar.x -= bar.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		bar.x += bar.speed
	}

	// bounce off
	// when bar reaches the edge
	if bar.x - bar.width/2 < 0 {
		bar.x = bar.width/2
	}
	if ScreenWidth < bar.x + bar.width/2 {
		bar.x = ScreenWidth - bar.width/2
	}
}

func (bar *Bar) Draw(screen *ebiten.Image)  {
	op := &ebiten.DrawImageOptions{}
	// set Bar.X, Y as the center of the image
	op.GeoM.Translate(-float64(bar.width)/2, -float64(bar.height)/2)
	op.GeoM.Translate(float64(bar.x), float64(bar.y))
	screen.DrawImage(bar.image, op)
	ebitenutil.DebugPrintAt(screen, "x: " + strconv.Itoa(bar.x), 0, ScreenHeight-26)
	ebitenutil.DebugPrintAt(screen, "y: " + strconv.Itoa(bar.y), 0, ScreenHeight-14)
}