package main

import "fmt"

func main() {
	//函数定义
	var m = max(1, 2)
	fmt.Println(m)
	x, y := swap("google", "baidu")
	fmt.Println(x, y)
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

//交换
func swap(x string, y string) (string, string) {
	return y, x
}
