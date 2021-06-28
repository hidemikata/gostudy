package main

import(
"fmt"
)

func test(params ...int){
    fmt.Println(len(params), params)
}

func main(){
    test(10,20)
    test(10,20, 30)

    l := []int{100,200}
    l = append(l, 300)
    test(l...)

    f := 1.11
    fmt.Println(int(f))

    s := []int{1, 2, 5, 6, 2, 3, 1}
    fmt.Println(s[2:4])//5,6
    m := map[string]int{"Mike":20, "Nancy":24, "Messi":30}
    fmt.Printf("%T %v", m, m)
}

