package main

import (
	"fmt"
	"sync"
)

func hello2(wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("world")
	}
	c <- 1
}
func hello(c chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("hello")
	}
	c <- 2
}

func main() {
	c := make(chan int) //maintとgoroutineでアタイの受け渡しができないので
	//チャネルを使って受け渡しする。
	var wg sync.WaitGroup
	wg.Add(1)
	go hello2(&wg, c)
	go hello(c)
	x := <-c //受け取ったやつを取り出す。//キューイングされる。
	fmt.Println(x)
	x = <-c //受け取ったやつを取り出す。
	fmt.Println(x)
	wg.Wait()
}
