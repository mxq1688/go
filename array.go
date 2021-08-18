package main

import "fmt"

/*数组申明赋值的四种方式
方式1：
	var arr [2]int
		arr[0]=1
		arr[1]=2
		或者
		arr = [2]int{1,2}
方式2：
	var arr = [2]int{1,2} #或 arr := [2]int{1,2}
方式3：根据初始值的个数自行推断数组的长度
	var arr = [...]int{1,2} #或 arr := [...]int{1,2}
方式4：索引
var arr = [...]int{1:1,0:2} #或 arr := [...]int{1:1,0:2}
*/

func main() {
	//数组申明：var variable_name [SIZE] variable_type
	var arr = [5]int{1, 2, 3, 4, 5}
	var arr1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8} //根据初始值的个数自行推断数组的长度
	var arr2 [2]int
	arr2 = [2]int{1, 2}
	fmt.Println(arr, arr1, arr2, arr3)

	// 二维数组： var arrayName [ x ][ y ] variable_type
	a := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}
	// 数组切片定义二维数组
	values := [][]int{}
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	//for遍历数组

	//指针数组
}
