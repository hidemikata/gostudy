package main

import (
	"fmt"
)

type Vertex struct {
	X int //小文字にしたらprivateになります。
	Y int
	S string
}

func vtest(A Vertex) {
	A.X = 100
}
func ptest(A *Vertex) {
	A.X = 200
	//(*A).x = 100//こう書かなくても.で勝手に参照してくれる。
}

func main() {
	v1 := Vertex{X: 10, Y: 20}
	fmt.Println(v1)
	fmt.Println(v1.X, v1.Y)

	v2 := Vertex{}
	fmt.Println(v2)

	var v3 Vertex
	fmt.Println(v3)

	v4 := new(Vertex)
	fmt.Printf("%T %v\n", v4, v4)

	v5 := &Vertex{} //v4と同じ挙動。pointerってわかりやすいのでこっちつかう。
	fmt.Printf("%T %v\n", v5, v5)

	//mapsliceはmake、structは＆でやるのがおおいらしい。

	test1 := Vertex{1, 2, "test1"}
	vtest(test1) //値渡し
	fmt.Println(test1)

	test2 := &Vertex{1, 2, "test2"}
	ptest(test2)
	fmt.Println(*test2)
	fmt.Println((*test2).X)

	ptest(&test1)
	fmt.Println(test1)

}
