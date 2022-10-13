package pack


import (
	"fmt"
	"math"
)




func (c Cube) Volume() float64 {
	return c.Length * c.Length * c.Length
}

func (s Sphere) Volume() float64 {
	return (4 * math.Pi * math.Pow(s.Radius, float64(3))) / 3
}

func (b Box) Volume() float64 {
	return b.Length * b.Width * b.Height
}

func CalculateVolume(kind OfStructure, called string) {
	fmt.Printf("The Volume calculated for our %s is: %f \n", called, kind.Volume())
}

func (c Cube) String() string {
	return fmt.Sprintf("The cube has length %v , breadth %v , height %v",c.Length,c.Length,c.Length )
}

func (s Sphere) String() string {
	return fmt.Sprintf("The Sphere has radius %v",s.Radius )
}

func (b Box) String() string {
	return fmt.Sprintf("The cube has length %v , breadth %v , height %v",b.Length,b.Length,b.Length )
}

