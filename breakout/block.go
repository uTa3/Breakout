package breakout

import "github.com/hajimehoshi/ebiten/v2"

type Block struct {
	x, y  int
	width, height int
	isDead bool
	image *ebiten.Image
}

func (block *Block) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(block.width), float64(block.height))
	op.GeoM.Translate(float64(block.x), float64(block.y))
	screen.DrawImage(block.image, op)
}