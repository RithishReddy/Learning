package main

import (

 "Exercise/shapes/pack"
 "fmt"
)



func main() {

	c := pack.Cube{
		Length: 7,
	}

	b := pack.Box{
		Length: 5.5,
		Width:  5.5,
		Height: 7.7,
	}

	s := pack.Sphere{
		Radius: 7.14,
	}

	pack.CalculateVolume(c, "Cube")
	pack.CalculateVolume(b, "Box")
	pack.CalculateVolume(s, "Sphere")

	fmt.Println(c)
	fmt.Println(s)
	fmt.Println(b)
}