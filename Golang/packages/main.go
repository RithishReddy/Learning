package main

import (

"Exercise/packages/shapes"
"Exercise/packages/shapes/measurements"
 "fmt"
)



func main() {

	c := shapes.Cube{
		Length: 7,
	}

	b := shapes.Box{
		Length: 5.5,
		Width:  5.5,
		Height: 7.7,
	}

	s := shapes.Sphere{
		Radius: 7.14,
	}

	measurements.CalculateVolume(c, "Cube")
	measurements.CalculateVolume(b, "Box")
	measurements.CalculateVolume(s, "Sphere")

	fmt.Println(c)
	fmt.Println(s)
	fmt.Println(b)
}