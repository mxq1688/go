package main

import "fmt"

type DivideError struct {
	dividee int
	divider int
}

//实现error接口
func (de *DivideError) Error() string {
	strFormat := `
		Cannot proceed, the divider is zero.
    	dividee: %d
    	divider: 0
	`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

func ml() (mm int) {
	t := 10
	return t //会把t赋值给mm
}

func main() {
	//错误处理
	//if result, errorMsg := Divide(100, 10); errorMsg == "" {
	//	//正常情况
	//	fmt.Println("100/10 = ", result)
	//}
	// 当被除数为零的时候会返回错误信息
	//if _, errorMsg := Divide(100, 0); errorMsg != "" {
	//	//fmt.Println("errorMsg is: ", errorMsg)
	//	m := errors.New(errorMsg)
	//	fmt.Println(m)
	//}

	//var mm error
	//mm = errors.New("faew")
	//if(mm != nil){
	//	panic(mm)
	//}

	//panic相当于其他语言里的throw，而recover相当于其他语言里的cacth，可是由于golang的recover机制要求必须在defer的函数里才能执行catch panic
	//defer必须在panic之前
	// 必须要先声明defer，否则不能捕获到panic异常  //defer通常用于资源释放、打印日志、异常捕获等
	//如果有多个defer函数，调用顺序类似于栈，越后面的defer函数越先被执行(后进先出)
	defer func() {

		fmt.Println("before-end")

		if err := recover(); err != nil {

			fmt.Println(err) // 这里的err其实就是panic传入的内容，55

		}

		fmt.Println("end")

	}()
	// fmt.Println(f1())
	// fmt.Println(f2())
	// fmt.Println(f3())
	// fmt.Println(ml())
	f()
}

func f() {

	fmt.Println("a")

	panic(55) //panic后就中断

	fmt.Println("b")

	fmt.Println("f")

}

//含有defer函数的外层函数，返回的过程是这样的：先给返回值赋值，然后调用defer函数，最后才是返回到更上一级调用函数中 下面三例子
func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
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
