package main

import (
	"fmt"
)

type MyInt int //intをMyIntで置き換える

func (i MyInt) Double() int {
	return int(i * 2) //castしないとMyInt型のまま。
}

func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}
