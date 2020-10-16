package breakout

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Boll struct {
	x, y int
	velocityX, velocityY int
	radius int
	image *ebiten.Image
}

func NewBoll(posX, posY, vX, vY, r int, img *ebiten.Image) *Boll {
	return &Boll{x: posX, y: posY, velocityX: vX, velocityY: vY, radius: r, image: img}
}

func (boll *Boll) Update(bar *Bar)  {
	/* bounce off */
	// when the ball reaches the edge
	if (boll.x - boll.radius < 0) || (ScreenWidth < boll.x + boll.radius) {
		boll.velocityX = -boll.velocityX
	}
	if (boll.y - boll.radius < 0) || (ScreenHeight < boll.y + boll.radius) {
		boll.velocityY = -boll.velocityY
	}	
	boll.x += boll.velocityX
	boll.y += boll.velocityY
}

func (boll *Boll) Draw(screen *ebiten.Image)  {
	op := &ebiten.DrawImageOptions{}
	// set Bar.X, Y as the center of the image
	op.GeoM.Translate(-float64(boll.radius), -float64(boll.radius))
	op.GeoM.Translate(float64(boll.x), float64(boll.y))
	screen.DrawImage(boll.image, op)
	ebitenutil.DebugPrintAt(screen, "x: " + strconv.Itoa(boll.x), 0, 0)
	ebitenutil.DebugPrintAt(screen, "y: " + strconv.Itoa(boll.y), 0, 12)
}
