package main

import (
		"image"
		 _ "image/png"
		"github.com/hajimehoshi/ebiten/v2"
)

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

type Game struct{}

func (g *Game) Update() error {
	return nil
	}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(PlayerSprite, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{}
	
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
		}
	}