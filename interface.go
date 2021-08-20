package main

import "fmt"

//接口定义
type Phone interface {
	call()
}

type NokiaPhone struct {
	title string
}

type IPhone struct {
	title string
}

//实现接口
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

//实现接口
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	//接口变量
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
