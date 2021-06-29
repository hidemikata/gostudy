package main

import (
	"fmt"
	"strconv"
)

func add(x int, y int) (int, int) {
	return x + y, x - y
}

func addconvstr(x, y int) (result string) {
	z := x + y
	result = strconv.Itoa(z)
	//    return result 省略可能
	return
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)
	fmt.Println(addconvstr(20, 30))

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	func(x int) {
		fmt.Println("inner func", x)
	}(3)

	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	cirlcle := hankeiGenerator(3.14)
	fmt.Println(cirlcle(5))
	cirlcle2 := hankeiGenerator(3)
	fmt.Println(cirlcle2(5))
}

func hankeiGenerator(pi float64) func(r float64) float64 {
	return func(r float64) float64 {
		return r * r * pi
	}
}

func incrementGenerator() func() int {
	x := 1
	return func() int {
		x++
		return x
	}
}
