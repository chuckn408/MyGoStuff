package main

import (
		"image"
		 _ "image/png"
		 "math"
		"github.com/hajimehoshi/ebiten/v2"
)

// although comments start the same as in this one, the below comment actually is read and executed
import "embed"
//go:embed assets/*
var assets embed.FS
var PlayerSprite = mustLoadImage("assets/player.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
		}
		defer f.Close()
		
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	
	return ebiten.NewImageFromImage(img)
}

type Vector struct {
	X float64
	Y float64
}

type Game struct {
	playerPosition Vector
}

func (g *Game) Update() error {

	speed := 5.0
	// g.playerPosition.X += speed
	
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.playerPosition.Y += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.playerPosition.Y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.playerPosition.X -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerPosition.X += speed
	}

	return nil
	}

// draw the ship here ; undo '_' in math import to play with stuff
func (g *Game) Draw(screen *ebiten.Image) {

	width := PlayerSprite.Bounds().Dx()
	height := PlayerSprite.Bounds().Dy()
	halfW := float64(width / 2)
	halfH := float64(height / 2)

	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(150, 200)   // coord location
	// op.GeoM.Scale(1, -1)
	op.GeoM.Scale(1, 1)  // scaling ; neg vals dont work it seems
	op.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)
	// op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(45.0 * math.Pi / 180.0)  // angles
	op.GeoM.Translate(halfW, halfH)
	
	// playing with colours
	//op := &colorm.DrawImageOptions{}
	//cm := colorm.ColorM{}
	//cm.Scale(1.0, 1.0, 1.0, 0.5)
	//colorm.DrawImage(screen, PlayerSprite, cm, op)
	
	screen.DrawImage(PlayerSprite, op)
}

// draw the sccreen here- pretty simple
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{
		playerPosition: Vector{X: 100, Y: 100},
	}
	
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
		}
	}
	