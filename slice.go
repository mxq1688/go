package main

import "fmt"

/*
	切片是对数组的抽象。
	定义切片
		声明一个未指定大小的数组来定义切片： var identifier []type
		使用 make() 函数来创建切片:
			var slice1 []type = make([]type, len, cap)  # len 是数组的长度并且也是切片的初始长度; cap 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
				也可以简写为
					slice1 := make([]type, len)  #不指定cap和len一样
	len() 方法获取长度。
	cap() 可以测量切片最长可以达到多少
*/
func main() {
	//定义
	// var identifier_nil []int          //定义空切片
	// identifier := make([]int, 10, 20) //创建切片
	// var slice1 []int = make([]int, 2)
	// s := identifier[:]  //初始化切片s,是数组identifier的引用
	// m := identifier[:5] //不包含右边
	// fmt.Println(identifier, slice1, s, m, len(s), cap(s))
	// fmt.Println(identifier_nil == nil)

	/* 在切片上施加切片 */
	// numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// numbers1 := make([]int, 0, 5)
	// printSlice(numbers1)
	// /* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	// number2 := numbers[2:5] //cap根据 2 来决定（9-2）
	// printSlice(number2)

	/*
		append() copy()
		当append(list, [params]) 原切片list 添加的切片parmas 结果切片tar

		1）如果slice的容量还有剩余，元素1直接追加到slice指向的底层数组。
		2) 如果slice没有剩余容量，append函数会创建一个新的切片，新切片容量为slice的两倍，并将slice中的数据拷贝到新切片，然后追加新元素。
	*/
	var numbers []int
	printSlice(numbers)
	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)
	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)
	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
