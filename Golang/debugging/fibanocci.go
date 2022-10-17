package main

func fib(n int) int {
	if n < 2{
		return n
	}

	var f int
	if v, ok := m[n]; ok {
		f = v
	} else {
		f = fib(n-1) +fib(n-2)
		m[n] = f
	}

	return f
}