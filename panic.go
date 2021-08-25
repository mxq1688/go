package main

import "fmt"

func main() {
	/*
		panic:
			1、会中断程序
			2、相当于其他语言里的throw，而recover相当于其他语言里的catch，可是由于golang的recover机制要求必须在defer的函数里才能执行catch
	*/
	/*
		defer:
			1、必须在panic之前，必须要先声明defer，否则不能捕获到panic异常
			2、defer通常用于资源释放、打印日志、异常捕获等
			3、如果有多个defer函数，调用顺序类似于栈，越后面的defer函数越先被执行(后进先出)
	*/
	defer func() {

		fmt.Println("before-error")

		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("end")
	}()

	f()
}

func f() {
	fmt.Println("a")
	panic("this is an error!") //会中断程序
	fmt.Println("b")
	fmt.Println("f")
}
