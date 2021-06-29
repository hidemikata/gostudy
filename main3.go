package main

import (
	"fmt"
)

func main() {
	//array
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	//array
	var b [2]int = [2]int{300, 400}
	fmt.Println(b)

	var s []int = []int{500, 600}
	s = append(s, 700)
	fmt.Println(s)

	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2:4])
	fmt.Println(n[:4])
	fmt.Println(n[4:])

	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
	}
	fmt.Println(board)
	fmt.Println(board[0])
	board[1] = append(board[1], 6, 7, 8)
	fmt.Println(board)

	m := make([]int, 3, 5)
	fmt.Printf("%T, len=%d, cap=%d, value=%v\n", m, len(m), cap(m), m)
	m = append(m, 0, 0, 0)
	fmt.Printf("%T, len=%d, cap=%d, value=%v\n", m, len(m), cap(m), m)

	mm := make([]int, 3)
	fmt.Printf("%T, len=%d, cap=%d, value=%v\n", mm, len(mm), cap(mm), mm)
	mm = append(mm, 0)
	fmt.Printf("%T, len=%d, cap=%d, value=%v\n", mm, len(mm), cap(mm), mm)

	mm1 := make([]int, 0) //0
	var mm2 []int         //nil
	fmt.Printf("%T, len=%d, cap=%d, value=%v\n", mm1, len(mm1), cap(mm1), mm1)
	fmt.Printf("%T, len=%d, cap=%d, value=%v\n", mm2, len(mm2), cap(mm2), mm2)

	//    mm2 = make([]int, 5)
	mm2 = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		mm2 = append(mm2, i)
		fmt.Println(mm2)
	}
	fmt.Println(mm2)

}
