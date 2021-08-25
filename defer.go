package main

import "fmt"

func main() {

	//匿名返回值
	fmt.Println(a())
	//有名返回值
	fmt.Println(ml())
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

}

/*
	defer:
		1、必须在panic之前，必须要先声明defer，否则不能捕获到panic异常
		2、defer通常用于资源释放、打印日志、异常捕获等
		3、如果有多个defer函数，调用顺序类似于栈，越后面的defer函数越先被执行(后进先出)
*/

/*
	defer：延迟函数
	含有defer函数的外层函数，返回的过程是这样的:
		1、匿名返回值
		2、有名返回值
*/

//匿名返回值
func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("a defer2:", i) // 打印结果为 a defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("a defer1:", i) // 打印结果为 a defer1: 1
	}()
	return i
}

//有名返回值
func ml() (mm string) {
	t := "test"
	return t //会把t赋值给mm
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0 //返回1
}
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t //先给r赋值 r = t 然后t变化不影响r
}
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r) //r是参数传进函数
	return 1
}
