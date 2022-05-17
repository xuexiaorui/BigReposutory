package main

import "fmt"

//变量初始化
// 实际上包变量a 并没有被初始化成3 init 函数中 a:=3  相当于声明了一 新的局部变量

var a int

func init() {
	a := 3
	fmt.Printf("a: %v\n", a)
}

func main() {
	fmt.Printf("a: %v\n", a)
}
