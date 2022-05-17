package main

// 变量声明

import (
	"fmt"
)

var s string

/*  变量声明可以一行声明多个变量，
包级别声明的变量会在main入口函数执行前完成初始化，
局部变量将在声明语句被执行到的时候完成初始化 */
var i, j, k int

func main() {

	fmt.Printf("s: %v\n", s)
	fmt.Printf("i: %T\n", i)
	fmt.Printf("j: %T\n", j)
	fmt.Printf("k: %T\n", k)

	/* 	简短变量声明采用 ：= 的方式
	其参数类型可以通过返回参数自行推断
	简短变量声明语句中必须至少要声明一个新的变量
	i:=3  no new variables on left side of :=compile*,
	修改为 i，j := 3,5*/

	z := 3 * 5
	i := 2
	//i:=3
	i, j := 3, 5

	fmt.Printf("z: %v\n", z)
	fmt.Printf("i: %v\n", i)
	fmt.Printf("j: %v\n", j)

}
