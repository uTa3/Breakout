package breakout

import (
	"image/color"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/opentype"
)

const (
	ScreenWidth  = 480
	ScreenHeight = 360
)

const (
	initBarX  = ScreenWidth / 2
	initBarY  = 4 * ScreenHeight / 5
	barWidth  = 100
	barHeight = 20

	initBallX  = ScreenWidth / 2
	initBallY  = 3 * ScreenHeight / 5
	ballRadius = 4

	blockWidth  = 60
	blockHeight = 30
	row         = 2
	column      = 5
)

var (
	barSpeed     = 8
	ballVelocity = 4

	imageClear    = ebiten.NewImage(ScreenWidth, ScreenHeight)
	imageGameover = ebiten.NewImage(ScreenWidth, ScreenHeight)
)

type Game struct {
	bar       *Bar
	ball      *Ball
	blocks    [][]*Block
	gameclear bool
	gameover  bool
}

func init() {
	f, err := opentype.Parse(goitalic.TTF)
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    64,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		log.Fatalf("NewFace: %v", err)
	}
	imageGameover.Fill(color.White)
	imageClear.Fill(color.White)
	drawTextCenter(imageGameover, "GAME OVER", face, 0, 0, color.Black)
	drawTextCenter(imageClear, "GAME CLEAR", face, 0, 0, color.Black)
}

func drawTextCenter(image *ebiten.Image, str string, f font.Face, x, y int, clr color.Color) {
	w := textWidth(f, str)
	x += (ScreenWidth - w) / 2
	text.Draw(image, str, f, x, ScreenHeight/2, clr)
}

func textWidth(f font.Face, str string) int {
	maxW := 0
	for _, line := range strings.Split(str, "\n") {
		b, _ := font.BoundString(f, line)
		w := (b.Max.X - b.Min.X).Ceil()
		if maxW < w {
			maxW = w
		}
	}
	return maxW
}

func NewGame() (*Game, error) {
	g := &Game{}
	var err error
	g.init()
	return g, err
}

// generate a new obj, bar, ball, block
func (g *Game) init() {
	g.bar = &Bar{
		x:      initBarX,
		y:      initBarY,
		speed:  barSpeed,
		width:  barWidth,
		height: barHeight,
		image:  ebiten.NewImage(barWidth, barHeight),
	}
	g.bar.image.Fill(color.White)

	g.ball = &Ball{
		x:         initBallX,
		y:         initBallY,
		velocityX: ballVelocity,
		velocityY: -ballVelocity,
		radius:    ballRadius,
		image:     ebiten.NewImage(ballRadius, ballRadius),
	}
	g.ball.image.Fill(color.White)

	g.blocks = make([][]*Block, row)
	for r := 0; r < row; r++ {
		for c := 0; c < column; c++ {
			g.blocks[r] = append(g.blocks[r], &Block{
				x:       blockWidth + 3*blockWidth/2*c,
				y:       blockHeight + 2*blockHeight*r,
				width:   blockWidth,
				height:  blockHeight,
				isAlive: true,
				image:   ebiten.NewImage(blockWidth, blockHeight),
			})
			g.blocks[r][c].image.Fill(color.White)
		}
	}
}

func (g *Game) Update() error {
	if g.gameover || g.gameclear {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.init()
			g.gameover, g.gameclear = false, false
		}
	}
	aliveBlocks := 0
	for r := 0; r < row; r++ {
		for c := 0; c < column; c++ {
			if g.blocks[r][c].isAlive {
				aliveBlocks++
			}
		}
	}
	if aliveBlocks == 0 {
		g.gameclear = true
	}
	if g.bar.y < g.ball.y {
		g.gameover = true
	}
	g.bar.Update()
	g.ball.Update()
	// when the ball hits the bar
	//  	 -------------
	// here->|           |<- here
	//  	 -------------
	if g.bar.x-g.bar.width/2 < g.ball.x+g.ball.radius && // left
		g.ball.x-g.ball.radius < g.bar.x+g.bar.width/2 && // right
		g.bar.y-g.bar.height/2 < g.ball.y && g.ball.y < g.bar.y+g.bar.height/2 {
		g.ball.velocityX = -g.ball.velocityX
		// Avoid overlapping with the bar (more accurate)
		if g.ball.x < g.bar.x {
			g.ball.x = g.bar.x - g.bar.width/2 - g.ball.radius
		} else if g.bar.x < g.ball.y {
			g.ball.x = g.bar.x + g.bar.width/2 + g.ball.radius
		}
	}
	//		  ↓  here  ↓
	//  	 -------------
	//       |           |
	//  	 -------------
	//		  ↑  here  ↑
	if g.bar.y-g.bar.height/2 < g.ball.y+g.ball.radius && // upper
		g.ball.y-g.ball.radius < g.bar.y+g.bar.height/2 && // lower
		g.bar.x-g.bar.width/2 < g.ball.x && g.ball.x < g.bar.x+g.bar.width/2 {
		g.ball.velocityY = -g.ball.velocityY
		if g.ball.y < g.bar.y {
			g.ball.y = g.bar.y - g.bar.height/2 - g.ball.radius
		} else if g.bar.y < g.ball.y {
			g.ball.y = g.bar.y + g.bar.height/2 + g.ball.radius
		}
	}

	for r := 0; r < row; r++ {
		for c := 0; c < column; c++ {
			if g.blocks[r][c].x-g.blocks[r][c].width/2 < g.ball.x+g.ball.radius && // left
				g.ball.x-g.ball.radius < g.blocks[r][c].x+g.blocks[r][c].width/2 && // right
				g.blocks[r][c].y-g.blocks[r][c].height/2 < g.ball.y && g.ball.y < g.blocks[r][c].y+g.blocks[r][c].height/2 {
				if g.blocks[r][c].isAlive {
					g.ball.velocityX = -g.ball.velocityX
					// Avoid overlapping with the bar (more accurate)
					if g.ball.x < g.blocks[r][c].x {
						g.ball.x = g.blocks[r][c].x - g.blocks[r][c].width/2 - g.ball.radius
					} else if g.blocks[r][c].x < g.ball.y {
						g.ball.x = g.blocks[r][c].x + g.blocks[r][c].width/2 + g.ball.radius
					}
				}
				g.blocks[r][c].isAlive = false
			}
			if g.blocks[r][c].y-g.blocks[r][c].height/2 < g.ball.y+g.ball.radius && // upper
				g.ball.y-g.ball.radius < g.blocks[r][c].y+g.blocks[r][c].height/2 && // lower
				g.blocks[r][c].x-g.blocks[r][c].width/2 < g.ball.x && g.ball.x < g.blocks[r][c].x+g.blocks[r][c].width/2 {
				if g.blocks[r][c].isAlive {
					g.ball.velocityY = -g.ball.velocityY
					if g.ball.y < g.blocks[r][c].y {
						g.ball.y = g.blocks[r][c].y - g.blocks[r][c].height/2 - g.ball.radius
					} else if g.blocks[r][c].y < g.ball.y {
						g.ball.y = g.blocks[r][c].y + g.blocks[r][c].height/2 + g.ball.radius
					}
				}
				g.blocks[r][c].isAlive = false
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.gameover {
		screen.DrawImage(imageGameover, nil)
	} else if g.gameclear {
		screen.DrawImage(imageClear, nil)
	} else {
		g.ball.Draw(screen)
		g.bar.Draw(screen)
		for r := 0; r < row; r++ {
			for c := 0; c < column; c++ {
				if g.blocks[r][c].isAlive {
					g.blocks[r][c].Draw(screen)
				}
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
