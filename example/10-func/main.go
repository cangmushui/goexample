package main

import "fmt"


//这里写一些很重要的点, 关于函数和方法中参数中参数的类型
//1.go语言中的值类型: 基本类型, struct, string, 数组
//2.引用类型 : 指针, 切片, 映射, 通道
//3.如果函数的一个参数是结构体, 一般的做法都是将参数设置成结构体指针, 一方面可以对结构体进行修改, 另一方面可以减少参数传递的拷贝成本

func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}
func updateList(list [3]int) {
	list[0] = 11
}
func updateString(str string) {
	str = "bye"
	fmt.Println(str)
}

func main() {
	//函数基本使用演示
	res := add(1, 2)
	fmt.Println(res) // 3

	v, ok := exists(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok) // A True

	//函数传递字符串和数组参数的演示
	a := [3]int{1, 2, 3}
	b := a
	a[0] = 10
	str := "hello"
	updateList(a)
	updateString(str)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(str)
}
