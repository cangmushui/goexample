package main

import "fmt"

func add2(n int) {
	n += 2
}

func add2ptr(n *int) {
	*n += 2
}

//1.在不涉及方法的部分go语言指针的使用和c语言相同, 唯一的区别是无法对指针进行修改, 另外就是go有一个语法糖访问某个指针*p的成员时可以直接
//使用p.variable, 其完整的写法应该是(*p).variable
//2.在涉及方法的部分, go语言有一些自己的特性, 见下方实例中的方法接受参数部分
func main() {
	//传值和传引用的区别
	n := 5
	add2(n)
	fmt.Println(n) // 5
	add2ptr(&n)
	fmt.Println(n) // 7

	//指针的使用
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值
	fmt.Println(p)  //查看此时指针指向的地址


	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值


	//这里的结论就是, 在执行方法时采用(*p).method()和p.method()均可, 最终是传值还是传引用要看方法中接受的参数是指针还是值
	a := &ll{value:1}
	(*a).set(2)
	fmt.Println(a.value)

	
}

type ll struct {
	value int
}
func (l ll) set(value int) {
	l.value = value
}
