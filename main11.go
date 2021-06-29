package main

import (
	"fmt"
)

func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n)
	one(&n)
	fmt.Println(n)

	var p *int     //これだとnil
	fmt.Println(p) //領域確保してない。
	//    *p//パニック

	var p2 *int = new(int) //領域確保してる
	fmt.Println(p2)
	fmt.Println(*p2) //0が刺されてう

	//map sliceはmake
	//newはポインターを返すものに使う。

	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	pp := new(int)
	fmt.Printf("%T\n", pp)

	ps := new(struct{})
	fmt.Printf("%T\n", ps)

}
