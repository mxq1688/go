package main

import (
	"fmt"
	"time"
)

func main() {
	//if
	a := 10
	if a < 3 {
		fmt.Println("a<3")
	} else {
		fmt.Println("a>3")
	}

	//switch 不用break
	value := 0
	switch value {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	/*
		select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。
		select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。

		每个 case 都必须是一个通信
		所有 channel 表达式都会被求值
		所有被发送的表达式都会被求值
		如果任意某个通信可以进行，它就执行，其他被忽略。
		如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
		否则：
			如果有 default 子句，则执行该语句。
			如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
	*/
	ch := make(chan int)
	c := 0
	stopCh := make(chan bool)
	go Chann(ch, stopCh)
	for {
		select {
		case c = <-ch:
			fmt.Println("Receive", c)
			fmt.Println("channel")
		case s := <-ch:
			fmt.Println("Receive", s)
		case _ = <-stopCh:
			goto end
		}
	}
end:
}

func Chann(ch chan int, stopCh chan bool) {
	for j := 0; j < 10; j++ {
		ch <- j
		time.Sleep(time.Second)
	}
	stopCh <- true
}
