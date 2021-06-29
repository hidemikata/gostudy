package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v   map[string]int
	mux sync.Mutex //鍵と一緒に宣言
}

func (c *Counter) Inc(key string) {
	c.mux.Lock() //ロックを取る
	fmt.Println(c.v[key], "+1")
	c.v[key]++
	c.mux.Unlock() //ロック解除
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}
func main() {
	c := Counter{v: make(map[string]int)} //mapはmakeで作る。
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()
	go func() {
		for i := 0; i < 13; i++ {
			c.Inc("key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c, c.Value("key"))
}
