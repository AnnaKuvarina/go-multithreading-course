package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"sync"
)

const (
	screenWidth  = 640
	screenHeight = 360
	boidsCount   = 500
	viewRadius   = 13
	adjRate      = 0.015
)

var (
	green   = color.RGBA{19, 255, 50, 255}
	boids   [boidsCount]*Boid
	boidMap [screenWidth + 1][screenHeight + 1]int
	rWLock  = sync.RWMutex{}
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y), green)
		screen.Set(int(boid.position.x-1), int(boid.position.y), green)
		screen.Set(int(boid.position.x+1), int(boid.position.y+1), green)
		screen.Set(int(boid.position.x+1), int(boid.position.y-1), green)
	}
}

func main() {
	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}

	for i := 0; i < boidsCount; i++ {
		boid := createBoid(i)
		boids[i] = boid
		updateBoidMapItem(boid, boid.id)
		go boid.start()
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Boids in a box")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
