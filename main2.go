package main

import (
"fmt"
"strconv"
)

func main() {
    var x int = 1
    xx := float64(x)
    fmt.Printf("%T\n", xx)

    var s string = "12"
    i, _ := strconv.Atoi(s)
    fmt.Printf("%T %v\n", i, i)

    h := "hello world"
    fmt.Println(string(h[0]))//ascii code to string
}

