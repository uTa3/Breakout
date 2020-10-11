package breakout

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


const (
	bollSize = 5
)

var (
	bollImage *ebiten.Image
	bollSpeed = 2
)

type Boll struct {
	x, y int
	dx, dy int
}

func NewBoll(posX, posY int) *Boll {
	return &Boll{x: posX, y: posY,dx: -bollSpeed,dy: -bollSpeed}
}

func init() {
	bollImage = ebiten.NewImage(bollSize, bollSize)
	bollImage.Fill(color.White)
}

func (b *Boll) move()  {
	// 衝突したら反転
	if (b.x < bollSize/2) || (b.x > ScreenWidth - bollSize/2) {
		b.dx *= -1
	}
	if (b.y < bollSize/2) || (b.y > ScreenHeight - bollSize/2) {
		b.dy *= -1
	}
	b.x += b.dx
	b.y += b.dy
}

func (b *Boll) Update()  {
	b.move()
}

func (b *Boll) Draw(screen *ebiten.Image)  {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-bollSize/2, -bollSize/2)
	op.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(bollImage, op)
	ebitenutil.DebugPrintAt(screen, "x: " + strconv.Itoa(b.x), 0, 0)
	ebitenutil.DebugPrintAt(screen, "y: " + strconv.Itoa(b.y), 0, 12)
}
