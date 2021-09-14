package main

import "fmt"

/*
	取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
	一个指针变量指向了一个值的内存地址。
	var var_name *var-type #var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针

	在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。
*/

func main() {
	//指针变量
	var a = 123
	var b *int = &a //申明指针变量o并赋值， 使用指针访问该地址的值 *b
	//指针的指针
	var c **int = &b
	var d ***int = &c
	fmt.Println(a, b, *b, c, *c, **c, d, ***d) //指向指针的指针

	/*
		指针数组：var ptr [size]*int

	*/
	var arr = []int{1, 2, 3}
	var arr_ptr [3]*int
	arr_ptr[0] = &arr[0]
	arr_ptr[1] = &arr[1]
	arr_ptr[2] = &arr[2]
	fmt.Println(arr, arr_ptr, *arr_ptr[0], *arr_ptr[1], *arr_ptr[2])

	/*
		指针函数：允许向函数传递指针，只需要在函数定义的参数上设置为指针类型即可
		实参：传地址
		形参：类型使用指针
	*/
	var m, n = 1, 2
	x, y := swap2(&m, &n)
	fmt.Println(m, n)
	fmt.Println(x, y)
}

func swap2(a *int, b *int) (int, int) {
	var tem int
	tem = *a
	*a = *b
	*b = tem
	return *a, *b
}
