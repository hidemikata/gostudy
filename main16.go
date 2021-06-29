package main

import (
	"fmt"
)

func hello(h interface{}) {
	//    hh := h.(int)//type assersion .(型)こう書く.
	//    fmt.Println(hh)

	switch v := h.(type) { //switch type文こう書く
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v, "!")
	default:
		fmt.Printf("%T not found\n", v)
	}
}

type MyNotFoundError struct { //カスタムエラー
	Mess1 string
	Mess2 string
}

func (e *MyNotFoundError) Error() string { //カスタムエラー
	return fmt.Sprintf("error msg:", e.Mess1)
}

func errfunc() error { //カスタムエラー
	return &MyNotFoundError{"error naiyou1", "err2"}
}

func main() {
	hello(10)
	hello("hello")
	hello(float64(1))
	fmt.Println(errfunc())

}
