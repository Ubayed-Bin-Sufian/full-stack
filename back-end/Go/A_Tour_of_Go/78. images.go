package main

import (
	"fmt"
	"image" // gives us basic image types like Rectangle and RGBA
)

func main() {
	// Create a new RGBA image with bounds from (0,0) up to (100,100).
	// That means the image is 100 pixels wide and 100 pixels tall.
	// "image.Rect" builds a rectangle: Min = (0,0), Max = (100,100).
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))

	// Print the bounds of the image (the rectangle).
	// Output: (0,0)-(100,100)
	fmt.Println(m.Bounds())

	// Ask for the color at pixel (0,0).
	// By default, a new RGBA image is filled with all zeros (transparent black).
	// The "RGBA()" method returns 4 uint32 values: red, green, blue, alpha.
	// Each is scaled to 16 bits (0–65535 range).
	// So you'll see (0, 0, 0, 0) → fully transparent black.
	fmt.Println(m.At(0, 0).RGBA())
}

// Output:
// (0,0)-(100,100)
// 0 0 0 0