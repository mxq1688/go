package main

import "fmt"

func main() {
	/*
		和 C 语言的 for 一样：for init; condition; post { }
		和 C 的 while 一样：for condition { }
		和 C 的 for(;;) 一样：or { }
	*/
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum1 := 1
	for sum1 <= 10 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	// sum2 := 0
	// for {
	// 	sum2++ // 无限循环下去
	// }
	// fmt.Println(sum2) // 无法输出

	//for 循环的 range 格式可以对 slice（切片）、map、数组、字符串等进行迭代循环。格式如下：
	str := "abcdefg"
	for key, value := range str {
		fmt.Println(key, value)
	}
}
