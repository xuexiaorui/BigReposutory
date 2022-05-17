package main

import "fmt"

// 指针

func incr(p *int) int {
	*p++
	return *p
}

func newIntA() *int {
	var a int
	return &a

}

func newIntB() *int {
	return new(int)
}

/* p指针指向变量a, *p++ 代表修改当前指针变量的字面值,相当于修改变量v 的值
可以理解为 *p 是一个变量 v 的一个别名，对于函数incr()不需要知道调用者的变量是谁
但是可以修改变量的值，这是一把双刃剑slice、map和chan，甚至 结构体、数组和接口
都会创建所引用变量的别名(来自go语言圣经)
*/

func main() {

	v := 1
	p := &v
	fmt.Printf("p: %v\n", p)
	fmt.Println(incr(p))
	fmt.Println(incr(p))
	fmt.Printf("p: %v\n", p)

	i := newIntA()
	j := newIntA()

	fmt.Printf("i: %v\n", i)
	fmt.Printf("j: %v\n", j)

}
