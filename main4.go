package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["new"] = 300
	fmt.Println(m)

	v, ok := m["new"]
	fmt.Println(v, ok)
	v, ok = m["nothing"]
	fmt.Println(v, ok)

	m2 := make(map[string]int)
	m2["pc"] = 500
	fmt.Println(m2)

	var m3 map[string]int
	//    m3["a"] = 5000

	var m4 map[string]int = map[string]int{"a": 1}
	fmt.Println(m4)
	m3 = map[string]int{"a": 1}
	m3["d"] = 4
	fmt.Println(m3)

	var a []int
	fmt.Println(a)
	a = append(a, 3)
	fmt.Println(a)

	b := make([]int, 3, 10)
	fmt.Println(b)

	c := []byte{72, 73}
	fmt.Println(c)
	fmt.Println(string(c))

	c2 := []byte("HI")
	fmt.Println(c2)
	fmt.Println(string(c2))

}
