package main

import "fmt"

func main() {
	//1.数组的基本操作
	var a [5]int
	a[4] = 100
	fmt.Println("get:", a[2])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//2.使用range遍历数组
	list := []int {11, 12, 13}
	for index, item := range list {
		fmt.Printf("index:%v, number:%v", index, item)
	}
	//若只想用item
	for _, item := range list {
		fmt.Printf("number:%v", item)
	}
	//若只想使用index
	for index := range list {
		fmt.Printf("index%v", index)
	}

	//结构体数组进行反序列化的时候, 必须传入&数组才行, 但是可以直接传进函数并进行修改
	//json反序列化可以对结构体指针数组进行, 所以尽量用指针数组


	//区分数组和切片
}
