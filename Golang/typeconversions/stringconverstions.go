package main

import (
	"fmt"
	// "math"
	"strconv"
)

func stringToInt(str string) int {

	val, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}
	return val
}

func stringToFloat(str string) float64 {
	// convert to float with 4 digits of precision

	val, err := strconv.ParseFloat(str, 0)
	if err != nil {
		panic(err)
	}
	val = float64(int(val*10000)) / 10000
	return val
}

func FloatToString(value float64) string {

	val := float64(int(value*100)) / 100
	str := fmt.Sprint(val)
	return str
}

func TestStringConverstions() {
	isFailed := false
	if stringToInt("10") != 10 {
		fmt.Println("Failed: stringToInt")
		isFailed = true
	}

	if stringToFloat("123.33333333333") != 123.3333 {
		fmt.Println("Failed: stringToFloat")
		isFailed = true
	}

	if FloatToString(1.0/3) != "0.33" {
		fmt.Println("Failed: stringToFloat")
		isFailed = true
	}

	if !isFailed {
		fmt.Println("All tests passed")
	}
}
