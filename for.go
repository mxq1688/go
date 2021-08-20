package main

import "fmt"

func main() {
	/*
		和 C 语言的 for 一样：for init; condition; post { }
		和 C 的 while 一样：for condition { }
		和 C 的 for(;;) 一样：or { }

		break 语句	经常用于中断当前 for 循环或跳出 switch 语句
		continue 语句	跳过当前循环的剩余语句，然后继续进行下一轮循环。
		goto 语句	将控制转移到被标记的语句。
			goto label
			..
			.
			label: statement
	*/
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//模拟while
	sum1 := 1
	for sum1 <= 10 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	//类似while(1)
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
