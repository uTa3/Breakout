package breakout

import (
	"image/color"
	_ "image/png"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	barWidth = 100
	barHeight = 10
)

var (
	barImage *ebiten.Image
	barSpeed = 8
)

type Bar struct {
	x, y int
}

func NewBar(posX, posY int) *Bar {
	return &Bar{x: posX, y: posY}
}

func init() {
	barImage = ebiten.NewImage(barWidth, barHeight)
	barImage.Fill(color.White)
}

func (b *Bar) move() {
	var dx int
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dx = -barSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		dx = barSpeed
	}
	// はみ出していなかったら足す
	if !(b.x + dx < barWidth/2 || b.x + dx > ScreenWidth - barWidth/2) {
		b.x += dx
	}
}

func (b *Bar) Update()  {
	b.move()
}

func (b *Bar) Draw(screen *ebiten.Image)  {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-barWidth/2, -barHeight/2)
	op.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(barImage, op)
	ebitenutil.DebugPrintAt(screen, "x: " + strconv.Itoa(b.x), 0, ScreenHeight-26)
	ebitenutil.DebugPrintAt(screen, "y: " + strconv.Itoa(b.y), 0, ScreenHeight-14)
}