package main

import (
"fmt"
)

func panicCreate(){
    panic("panic!")//普通は自分で実装しない
}

func save(){
    defer func(){//呼び出しより先にdeferする
        s:=recover()//panicをキャッチして処理する。
        fmt.Println(s)
    }()
    panicCreate()
}

func main() {
    save()
    fmt.Println("ok?")

    l := []int{100, 300, 23,11, 23, 2,4,6,4}

    var min int
    for i, v := range l {
        if i == 0{
            min = v
            continue
        }
        if v < min {
            min = v
        }
    }
    fmt.Println(min)


    m := map[string]int{
        "apple":  200,
        "banana": 300,
        "grapes": 150,
        "orange": 80,
        "papaya": 500,
        "kiwi":   90,
    }
    sum := 0
    for _, v := range m {
        sum +=v
    }
    fmt.Println(sum)







}
