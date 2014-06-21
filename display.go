package vm

import "fmt"

const BLANK = "_"

type Display struct {
	canvas [64][32]string

	// Display dimensions
	width  int
	height int

	xScale int
	yScale int
}

func NewDisplay() *Display {
	dm := Display{canvas: [64][32]string{}, width: 64, height: 32}
	dm.Clear()
	return &dm
}

func (display *Display) Clear() {
	for i := 0; i < display.width; i++ {
		for j := 0; j < display.height; j++ {
			display.canvas[i][j] = BLANK
		}
	}
}

func (display *Display) Draw() {
	fmt.Println(display.canvas)
}

func (display *Display) XorPixel(x int, y int) bool {
	// Wrap around vertically
	if x > display.width {
		x -= display.width
	} else if x < 0 {
		x += display.width
	}

	// Wrap around horizontally
	if y > display.height {
		y -= display.height
	} else if y < 0 {
		y += display.height
	}

	// Set the pixel state
	active := display.canvas[x][y] == BLANK
	if active {
		// Javascript was: var active = this.canvas[x][y] ^= 1
		display.canvas[x][y] = "*"
	} else {
		display.canvas[x][y] = " "
	}
	return active
}
