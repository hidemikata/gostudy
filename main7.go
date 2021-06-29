package main

import (
	"fmt"
)

func by2(a int) string {
	if a%2 == 0 {
		return "OK"
	}
	return "NG"
}

func main() {
	x := 3
	y := 2
	if x < 2 {
		fmt.Println("x<2")
	} else if x == y {
		fmt.Println("x<2")
	} else {
		fmt.Println("else")
	}
	if result := by2(2); result == "OK" {
		fmt.Println("by2ok")
	}
	if result := by2(3); result == "OK" {
		fmt.Println("by2ok")
	}

	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("san")
			continue
		}
		if i > 5 {
			fmt.Println("five")
			break
		}

		fmt.Println(i)
	}
	sum := 1
	//for ; sum < 10; {
	for sum < 10 {
		fmt.Println("sum=", sum)
		sum++
	}

	//    for {
	//        fmt.Print("no stop loop")
	//    }

	l := []string{"A", "B", "C"}
	for i := 0; i < len(l); i++ {
		fmt.Println(l[i])
	}
	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"A": 1, "B": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for _, v := range m {
		fmt.Println(v)
	}
	for k, _ := range m {
		fmt.Println(k)
	}

}
