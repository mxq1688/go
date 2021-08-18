package main

/*
	切片是对数组的抽象。
	定义切片
		声明一个未指定大小的数组来定义切片： var identifier []type
		使用 make() 函数来创建切片:
			var slice1 []type = make([]type, len)
			也可以简写为
			slice1 := make([]type, len)

*/
func main() {
	//Slice 切片
	// var identifier_nil []int          //定义空切片
	// identifier := make([]int, 10, 20) //创建切片
	// var slice1 []int = make([]int, 2)
	// s := identifier[:]  //初始化切片s,是数组identifier的引用
	// m := identifier[:5] //初始化切片s,是数组identifier的引用
	// fmt.Println(identifier, slice1, s, m, len(s), cap(s))
	// fmt.Println(identifier_nil == nil)

	//numbers1 := make([]int,0,5)
	//fmt.Println(numbers1, len(numbers1), cap(numbers1))

	//切片append() copy()
	//var numb []int
	//numb = []int{1,2,3}
	//num1 := append(numb, 5,4,3)
	//num2 := make([]int, len(num1), cap(num1)*2)
	//copy(num2, num1)
	//fmt.Println(numb, num1, num2)
}
