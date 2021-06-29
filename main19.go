package main

import (
	"fmt"
)

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum //その都度送信
	}
	close(c)
}

func main() {
	ch := make(chan int, 2) //容量２のキュー
	ch <- 100
	fmt.Println(len(ch))
	x := <-ch
	fmt.Println(len(ch))
	fmt.Println(x)
	ch <- 100
	ch <- 200
	//ch <- 3//error
	fmt.Println(len(ch))
	x = <-ch
	fmt.Println(len(ch))
	fmt.Println(x)
	ch <- 300 //ok
	close(ch) //これがないと終端が不明なのでforでエラーになる。
	for c := range ch {
		fmt.Println(c)
	}

	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))
	go goroutine1(s, c)
	for i := range c {
		fmt.Println(i) //都度送信されるやつを取り出して表示
	}
	fmt.Println("end")

}
