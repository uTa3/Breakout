package breakout

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ball struct {
	x, y                 int
	velocityX, velocityY int
	radius               int
	image                *ebiten.Image
}

func (ball *Ball) Update() {
	// bouncd off
	// when the ball reaches the edge
	if (ball.x-ball.radius < 0) || (ScreenWidth < ball.x+ball.radius) {
		ball.velocityX = -ball.velocityX
	}
	if (ball.y-ball.radius < 0) || (ScreenHeight < ball.y+ball.radius) {
		ball.velocityY = -ball.velocityY
	}
	ball.x += ball.velocityX
	ball.y += ball.velocityY
}

func (ball *Ball) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	// set Bar.X, Y as the center of the image
	op.GeoM.Translate(-float64(ball.radius), -float64(ball.radius))
	op.GeoM.Translate(float64(ball.x), float64(ball.y))
	screen.DrawImage(ball.image, op)
	ebitenutil.DebugPrintAt(screen, "x: "+strconv.Itoa(ball.x), 0, 0)
	ebitenutil.DebugPrintAt(screen, "y: "+strconv.Itoa(ball.y), 0, 12)
}
