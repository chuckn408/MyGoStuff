package main

import "math/rand"

type Part struct {
	X int
	Y int
}

type SnakeBody struct {
	Parts  []Part
	Xspeed int
	Yspeed int
	newColor string
}

func (sb *SnakeBody) ChangeDir(vertical int, horizontal int) {
	sb.Yspeed = vertical
	sb.Xspeed = horizontal
}

func (sb *SnakeBody) Update(width int, height int, longerSnake bool) {
	sb.Parts = append(sb.Parts, sb.Parts[len(sb.Parts)-1].GetUpdatedPart(sb, width, height))
	if !longerSnake {
		sb.Parts = sb.Parts[1:]
	}
}

func (sb *SnakeBody) RndClr(g.Score) {

	colorPikr := (g.Score / (rand.Intn))
/*	switch chgColor := runtime.GOOS; {
	default:
		sb.newColor := "Black"
	case colorPikr = 1: 
		sb.newColor = "Red"
	case colorPikr = 2:
		sb.newColor := "Blue"
	case colorPikr = 3:
		sb.newColor := "Green"
	case colorPikr = 4: 
		sb.newColor := "White"
	sb.newColor 
	}*/
	Println("Color is ", colorPikr)
}

func (sb *SnakeBody) ResetPos(width int, height int) {
	snakeParts := []Part{
		{
			X: int(width / 2),
			Y: int(height / 2),
		},
		{
			X: int(width/2) + 1,
			Y: int(height / 2),
		},
		{
			X: int(width/2) + 2,
			Y: int(height / 2),
		},
	}
	sb.Parts = snakeParts
	sb.Xspeed = 1
	sb.Yspeed = 0
}

func (sp *Part) GetUpdatedPart(sb *SnakeBody, width int, height int) Part {
	newPart := *sp
	newPart.X = (newPart.X + sb.Xspeed) % width
	if newPart.X < 0 {
		newPart.X += width
	}
	newPart.Y = (newPart.Y + sb.Yspeed) % height
	if newPart.Y < 0 {
		newPart.Y += height
	}
	RndClr()
	return newPart
}