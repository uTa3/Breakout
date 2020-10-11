package breakout

import "github.com/hajimehoshi/ebiten/v2"

const (
	ScreenWidth  = 480
	ScreenHeight = 360
)

var (
	barFirstPosX = ScreenWidth/2
	barFirstPosY = 4*ScreenHeight/5

	bollFirstPosX = ScreenWidth/2
	bollFirstPosY = 3*ScreenHeight/5
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
	g.bar = NewBar(barFirstPosX, barFirstPosY)
	g.boll = NewBoll(bollFirstPosX, bollFirstPosY)
	return g, err
}

func (g *Game) Update() error {
	g.bar.Update()
	g.boll.Update()
	// g.block.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ここでscreenをDrawに渡すのと
	// barImageを他でいじってスクリーンに描画,つまりg.bar.Draw(barImage)してからscreen.Draw(varImage, op)
	// このDrawがインターフェースを満たすための大本のDrawだからscreenをいろんなところに渡したくないかも
	// 後者はgame.goでbarImageに相当するものをもつか,barImageにアクセスできるようにするかどちらかしなければならない
	g.boll.Draw(screen)
	g.bar.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}