package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)  //100msごとにchanelに入る
	boom := time.After(500 * time.Millisecond) //500ms後にchanelに入る。

OutBreakhoge:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom.")
			break OutBreakhoge
			//return//mainを抜けたら止まるが、後続に行かない。
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

	//呼ばれるためにはbreakをしこむ。
	fmt.Println("###")

}
