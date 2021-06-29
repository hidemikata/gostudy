package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

//構造体（クラス）のメソッド見たなやつ
//値レシーバー
func (v Vertex) multiMethod() int {
	return v.X * v.Y
}

func (v *Vertex) pMultiMethod(scale int) { //参照レシーバー
	//ポンター先が書き変わる
	v.X = v.X * scale
	v.Y = v.Y * scale
}

func multi(v Vertex) int { //単なるメソッド
	return v.X * v.Y
}

func New(x, y int) *Vertex { //コンストラクタ的なもの
	return &Vertex{x, y}
	//    v := Vertex{x, y}
	//    return &v //この参照はいつ消えるんだろうか。
}

//継承3D
type Vertex3D struct {
	Vertex
	Z int
}

func (v Vertex3D) multiMethod3D() int {
	return v.X * v.Y * v.Z
}

func (v *Vertex3D) pMultiMethod3D(scale int) { //参照レシーバー
	//ポンター先が書き変わる
	v.X = v.X * scale
	v.Y = v.Y * scale
	v.Z = v.Z * scale
}

func multi3D(v Vertex3D) int { //単なるメソッド
	return v.X * v.Y * v.Z
}

func New3D(x, y, z int) *Vertex3D { //コンストラクタ的なもの
	return &Vertex3D{Vertex{x, y}, z}
	//    v := Vertex{x, y}
	//    return &v //この参照はいつ消えるんだろうか。
}
func main() {
	vertex := Vertex{X: 10, Y: 20}
	xy := multi(vertex)
	fmt.Println(xy)

	xy = vertex.multiMethod()
	fmt.Println(xy)

	vertex.pMultiMethod(10)
	fmt.Println(vertex)

	cv := New(30, 40)
	fmt.Println(*cv)

	vertex3d := Vertex3D{Vertex{10, 20}, 30}
	fmt.Println(vertex3d)
	vertex3d.pMultiMethod3D(10)
	fmt.Println(vertex3d)

	var pver3d *Vertex3D
	pver3d = New3D(10, 20, 30)
	fmt.Println(pver3d)
	fmt.Println(pver3d.Z)
	fmt.Println(pver3d.multiMethod())
	pver3d.pMultiMethod(10)
	fmt.Println(pver3d)
}
