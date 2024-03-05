package main

import (
	"fmt"
	"math"
	"time"

	// "github.com/veandco/go-sdl2/mix"
	sdl "github.com/veandco/go-sdl2/sdl"
	// "github.com/veandco/go-sdl2/ttf"
	mog "sdl2-3d-sandbox/internal/mog"
)

const (
	WindowSizeX = 800
	WindowSizeY = 600
)

func main() {
	mog.Init()
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(mog.MW, "\nFailed to initialize a window: %v", err)
		panic("Failed to initialize a window")
	}

	window, err := sdl.CreateWindow("Simple SDL2 Project",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		WindowSizeX, WindowSizeY,
		sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)

	if err != nil {
		fmt.Fprintf(mog.MW, "\nFailed to create the SDL Window: %v", err)
		panic("Failed to create the SDL Window")
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(mog.MW, "\nFailed to Initialize renderer: %v", err)
		panic("Failed to Initialize renderer")
	}
	defer renderer.Destroy()
	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				sdl.Quit()
				return
			}
		}
		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()
		renderer.SetDrawColor(255, 255, 255, 255)

		updateAndDraw(renderer)
		renderer.Present()
	}
}

// Main Code

var points []Point3D = []Point3D{
	{-1.0, -1.0, -1.0}, {-1.0, -1.0, 1.0},
	{1.0, -1.0, -1.0}, {-1.0, 1.0, -1.0},
	{-1.0, 1.0, 1.0}, {1.0, -1.0, 1.0},
	{1.0, 1.0, -1.0}, {1.0, 1.0, 1.0},
}

var vertices []Vertex = []Vertex{
	{0, 1}, {0, 2}, {0, 3},
	{2, 5}, {3, 6}, {3, 4},
	{4, 7}, {6, 7}, {7, 5},
	{5, 1}, {4, 1}, {2, 6},
}

var Rotation float64 = 0
var FOV float64 = 10.0
var DeltaTime float64 = 0

var time1 time.Time = time.Now()

func updateAndDraw(renderer *sdl.Renderer) {
	Rotation += 1 * DeltaTime

	for _, vert := range vertices {
		rotStart := getRotPoint(points[vert.Start])
		rotEnd := getRotPoint(points[vert.End])
		start := projection(rotStart)
		end := projection(rotEnd)
		renderer.DrawLine(int32(start.X), int32(start.Y), int32(end.X), int32(end.Y))
	}

	time2 := time.Now()
	duration := time2.Sub(time1)
	DeltaTime = duration.Seconds()
	time1 = time2
}

func getRotPoint(p Point3D) Point3D {
	return rotateX(rotateY(p))
}

func projection(p Point3D) Point2D {
	x := WindowSizeX/2 + (FOV*p.X)/(FOV+p.Z)*100
	y := WindowSizeY/2 + (FOV*p.Y)/(FOV+p.Z)*100
	return Point2D{x, y}
}

func rotateX(p Point3D) Point3D {
	x := p.X
	y := math.Cos(Rotation)*p.Y - math.Sin(Rotation)*p.Z
	z := math.Sin(Rotation)*p.Y + math.Cos(Rotation)*p.Z
	return Point3D{x, y, z}
}
func rotateY(p Point3D) Point3D {
	x := math.Cos(Rotation)*p.X - math.Sin(Rotation)*p.Z
	y := p.Y
	z := math.Sin(Rotation)*p.X + math.Cos(Rotation)*p.Z
	return Point3D{x, y, z}
}

type Point2D struct {
	X float64
	Y float64
}
type Point3D struct {
	X float64
	Y float64
	Z float64
}
type Vertex struct {
	Start int
	End   int
}
