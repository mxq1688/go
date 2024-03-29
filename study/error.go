package main

import (
	"errors"
	"fmt"
)

/*
	error类型是一个接口类型，这是它的定义：
		type error interface {
			Error() string
		}
*/

type DivideError struct {
	dividee int
	divider int
}

//实现error接口 参考函数方法
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

func main() {
	//正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println(errorMsg)
	}
	// 使用errors.New 可返回一个错误信息
	var err error
	err = errors.New("this is an error")
	fmt.Println(err)
	if err != nil {
		panic(err)
	}

}
