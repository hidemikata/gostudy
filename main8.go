package main

import (
	"fmt"
)

func getNumber() int {
	return 3
}
func main() {
	defer fmt.Println(1) //deferは関数が終わったときに実行
	defer fmt.Println(2) //先に実行(スタックに対になる。
	//    x:=2
	//    switch x {
	switch x := getNumber(); x { //ここで定義できる。
	case 1:
		fmt.Println("ichi")
	case 2:
		fmt.Println("ni")
	default:
		fmt.Println("sonota")
	}
	y := 0
	switch { //かかなくてもいい
	case y == 0:
		fmt.Println("zero")
	case y < 0:
		fmt.Println("fu")
	case y > 0:
		fmt.Println("sei")
	}
}
