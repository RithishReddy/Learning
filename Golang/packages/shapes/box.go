package shapes

import (
	"fmt"
)

type Box struct {
	Length float64
	Width  float64
	Height float64
}

func (b Box) String() string {
	return fmt.Sprintf("The cube has length %v , breadth %v , height %v", b.Length, b.Length, b.Length)
}

func (b Box) Volume() float64 {
	return b.Length * b.Width * b.Height
}