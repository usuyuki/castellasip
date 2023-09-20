package main

import (
	"fmt"
	"image/color"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type game struct {
	playerX, playerY float64
}

func (g *game) Update() error {
	// プレイヤーの移動
	const moveSpeed = 4
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.playerY -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.playerY += moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.playerX -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerX += moveSpeed
	}

	// 画面外に出ないように制限
	g.playerX = math.Max(0, math.Min(screenWidth-16, g.playerX))
	g.playerY = math.Max(0, math.Min(screenHeight-16, g.playerY))

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	// プレイヤーを描画（単純な四角形として）
	playerImg := ebiten.NewImage(16, 16)
	playerImg.Fill(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerX, g.playerY)
	screen.DrawImage(playerImg, op)

	// FPSを描画
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Simple RPG")
	if err := ebiten.RunGame(&game{}); err != nil {
		fmt.Fprintf(os.Stderr, "RunGame failed: %v\n", err)
		os.Exit(1)
	}
}
