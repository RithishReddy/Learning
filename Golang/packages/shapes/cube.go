package shapes

import (
	"fmt"
)

type Cube struct {
	Length float64
}

func (c Cube) Volume() float64 {
	return c.Length * c.Length * c.Length
}

func (c Cube) String() string {
	return fmt.Sprintf("The cube has length %v , breadth %v , height %v", c.Length, c.Length, c.Length)
}