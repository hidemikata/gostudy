package main

import (
	"fmt"
	"time"
)

func goroutine1(ch1 chan int) {
	for i := 0; i < 5; i++ {
		ch1 <- i + 1
	}
	close(ch1)
}
func goroutine2(ch1 chan int, ch2 chan int) {
	defer close(ch2)
	for v := range ch1 {
		ch2 <- v + 1
	}
	fmt.Println("2 end")
}
func goroutine3(ch2 chan int, ch3 chan int) {
	for v := range ch2 {
		ch3 <- v + 1
	}
	fmt.Println("3 end")
}

func groutineselect1(ch1 chan string) {
	for {
		ch1 <- "packet1"
		time.Sleep(1 * time.Second)
	}
}

func groutineselect2(ch2 chan string) {
	for {
		ch2 <- "packet2"
		time.Sleep(1 * time.Second)
	}
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go goroutine1(ch1)
	go goroutine2(ch1, ch2)
	go goroutine3(ch2, ch3)

	x := 0
	for i := 0; i < 5; i++ {
		x = <-ch3
		fmt.Println(x)
	}
	time.Sleep(2 * time.Second)

	chse1 := make(chan string)
	chse2 := make(chan string)
	go groutineselect1(chse1)
	go groutineselect2(chse2)

	for {
		select { //selectでチャネルに入るまで待つ。
		case msg1 := <-chse1:
			fmt.Println(msg1)
		case msg2 := <-chse2:
			fmt.Println(msg2)
		}
	}

}
