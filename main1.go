package main

import (
	"fmt"
    "strings"
)

var (
	i   int     = 1
	f64 float64 = 1.2
	s   string  = "test"
	t   bool    = true
)

const Pi = 3.14

const (
	Username = "test user"
	Password = "pass"
)

func main() {

    fmt.Println("hello" + "world")
    fmt.Println("hello"[3])
    fmt.Println(string("hello"[3]))
    var s string = "world"
    s = strings.Replace(s, "w", "a", 1)
    fmt.Println(s)
	fmt.Println(strings.Contains(s, "orld"))
	fmt.Println(`test
      test"`)

    tt, ff := true, false
    fmt.Printf("%T %v %t\n", tt, tt, tt)
    fmt.Printf("%T %v %t\n", tt, tt, 0)
    fmt.Printf("%T %v %t\n", ff, ff, ff)

	fmt.Println(Username)
	fmt.Println(Password)
	i = 3

	fmt.Println(i, f64, s, t)

	xi := 1
	var xf64 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	xi = 2
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T %v\n", xf64, xf64)

	fmt.Println("10 -1  =", 10-1)

}
