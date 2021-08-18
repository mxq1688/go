package main

import "fmt"

//接口定义
type Phone interface {
	call()
}

type NokiaPhone struct {
	title string
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println(nokiaPhone)
}

type IPhone struct {
	title string
}

func (iPhone IPhone) call() {
	fmt.Println(iPhone)
}

func main() {
	//接口
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
