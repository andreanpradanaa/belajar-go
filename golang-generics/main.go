package main

import "fmt"

type Number interface {
	int64 | float64 | int32 | float32
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[int]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, val := range m {
		s += val
	}
	return s
}

func main() {
	ints := map[string]int64{
		"first":  1,
		"second": 2,
	}

	floats := map[int]float64{
		1: 1.0,
		2: 2.0,
	}

	fmt.Println(
		"non-generic func: ",
		SumInts(ints),
		SumFloats(floats),
	)

	fmt.Println(
		"generic func",
		SumNumbers(ints),
		SumNumbers(floats),
	)

}
