package main

import "fmt"

var m = make(map[int]int)

func main(){



	for _, n := range []int{5, 1, 9, 98, 6} {
		x := fib(n)
		fmt.Println(n, "fib", x)
	}

}