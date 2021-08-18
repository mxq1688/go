package main

import "fmt"

func main() {
	//指针
	var m = 123
	var o *int = &m
	var oo **int = &o
	var ooo ***int = &oo
	fmt.Println(m, o, *o, oo, *oo, **oo, ooo, ***ooo) //指向指针的指针

	//指针数组
	// var arr = []int{1, 2, 3}
	// var arr_ptr [3]*int
	// arr_ptr[0] = &arr[0]
	// arr_ptr[1] = &arr[1]
	// arr_ptr[2] = &arr[2]
	// fmt.Println(arr, arr_ptr, *arr_ptr[0])

	//指针函数  	//函数参数传引用
	//var a,b = 1,2
	//c,d := swap2(&a, &b);
	//fmt.Println(a,b, c,d)
}
