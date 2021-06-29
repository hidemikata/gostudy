package main

import (
	"fmt"
	"sync"
	"time"
)

//producer, consumer
//producerからチャネルを利用してconsumerに値を渡す。
//両方ともgoroutine

func producer(ch chan int, i int) {
	ch <- i * 2
}
func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("process, ", i*1000)
		wg.Done()
		//chに入るまで待つ。
	}
	fmt.Println("###")
}
func main() {
	ch := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}
	go consumer(ch, &wg)
	wg.Wait()
	close(ch)
	time.Sleep(2 * time.Second)
	fmt.Println("end")

}
