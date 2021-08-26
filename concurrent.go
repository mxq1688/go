package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

//channel通道是引用类型

/*
并发：
	go 函数名( 参数列表 )
	Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。
	同一个程序中的所有 goroutine 共享同一个地址空间。

通道（channel）：
	通道（channel）是用来传递数据的一个数据结构。
	通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。

	注意：默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。
		channel有两种形式:
			一种是无缓冲的，一个线程向这个channel发送了消息后，会阻塞当前的这个线程，直到其他线程去接收这个channel的消息。
				无缓冲的形式：intChan := make(chan int)
			带缓冲的channel，是可以指定缓冲的消息数量，当消息数量小于指定值时，不会出现阻塞，超过之后才会阻塞，需要等待其他线程去接收channel处理
				带缓冲的形式：intChan := make(chan int， 3)
		var ch chan string; // nil channel 通道零值为nil（没有任何作用）

	判断通道是否关闭：
	 	1 第二个字段为true时，channel可能没关闭，也可能已经关闭，不能证明什么
		2 第二个字段为false时，可以证明channel中已没有残留数据且已关闭
		if v, ok := <- ch; ok {
			fmt.Println(v)
		}

读取通道：
	1、for range
	2、for ok
		for{
			if v, ok := <- ch; ok {
				fmt.Println(v)
			}else{
				break
			}
		}
*/

func main() {
	//并发
	// go say("2")
	// say("1")
	//执行以上代码，你会看到输出的是没有固定先后顺序。因为它们是两个 goroutine 在执行

	/*channel*/
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收
	fmt.Println(x, y, x+y)

	/*
		range 函数遍历每个从通道接收到的数据，因为 d 在发送完 10 个数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据之后就结束了。
		如果 d 通道不关闭，那么 range 函数就不会结束，从而在接收第 11 个数据的时候就阻塞了。
		close()：
			1、试图向已经关闭的channel发送数据会导致panic
			2、关闭已经关闭的channel会导致panic
			3、无论怎样都不应该在接收端关闭channel,因为在接收端无法判断发送端是否还会向通道中发送元素值
	*/
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)
	for i := range ch {
		fmt.Println(i)
	}

	/*	1 个发送者 1个接收者（最简单的），发送端直接关闭channel	*/
	// testClose()
	/* 	1个发送者 N个接收者（进行扩展）	*/
	// testClose1()
	/*	N 个发送者 1个接收者	添加一个 停止通知 接收端告诉发送端不要发送了  */
	// testClose2()
	/*
		select 用法类似与 IO 多路复用，可以同时监听多个 channel 的消息状态

		for循环select时，如果其中一个case通道已经关闭，则每次都会执行到这个case。
			怎么样才能不读关闭后通道:
				当返回的ok为false时，执行c = nil 将通道置为nil
		如果select里边只有一个case，而这个case被关闭了，则会出现死循环。
	*/
	/*	N个发送者 M个接收者	*/
	testClose3()
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, time.Millisecond)
	}
}
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}
func fibonacci(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	close(c)
}

func testClose() {
	notify := make(chan int)
	datach := make(chan int, 100)
	go func() {
		<-notify //接收数据，没接收到数据一直阻塞
		fmt.Println("2 秒后接受到信号开始发送")
		for i := 0; i < 100; i++ {
			datach <- i
		}
		fmt.Println("发送端关闭数据通道")
		close(datach)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("开始通知发送信息")
	notify <- 1
	time.Sleep(1 * time.Second)
	fmt.Println("通知1秒后接收到数据通道数据 ")
	for {
		if i, ok := <-datach; ok {
			fmt.Println(i)
		} else {
			fmt.Println("接收不到数据中止循环")
			break
		}
	}
	time.Sleep(5 * time.Second)
}
func testClose1() {
	notify := make(chan int)
	datach := make(chan int, 100)
	go func() {
		<-notify //接收数据，没接收到数据一直阻塞
		fmt.Println("2 秒后接受到信号开始发送")
		for i := 0; i < 100000; i++ {
			datach <- i
		}
		fmt.Println("发送端关闭数据通道")
		close(datach)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("开始通知发送信息")
	notify <- 1
	time.Sleep(100 * time.Second)
	fmt.Println("3秒后接受到数据通道数据 此时datach 在发送端已经关闭")
	for n := 0; n < 5; n++ {
		go func() {
			for {
				if i, ok := <-datach; ok {
					fmt.Println(i)
				} else {
					break
				}
			}
		}()

	}
	time.Sleep(5 * time.Second)
}

/*	N 个发送者 1个接收者	添加一个 停止通知 接收端告诉发送端不要发送了  */
type T int

func testClose2() {
	dataCh := make(chan T, 1)
	stopCh := make(chan T)
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				value := T(rand.Intn(10000))
				select {
				case <-stopCh:
					// for循环select时，如果其中一个case通道已经关闭，则每次都会执行到这个case。
					fmt.Println("接收到停止发送的信号")
					// return //干净利落，适合退出goroutin的场景
				case dataCh <- value:
				}
			}
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("1秒后开始接收数据")
	for {
		if d, ok := <-dataCh; ok {
			fmt.Println(d)
			if d == 9999 {
				fmt.Println("当在接收端接收到9999时告诉发送端不要发送了")
				close(stopCh)
				return
			}
		} else {
			break
		}
	}
}

func testClose3() {
	dataCh := make(chan T, 100)
	toStop := make(chan string)
	stopCh := make(chan T)
	//简约版调度器
	go func() {
		if t, ok := <-toStop; ok {
			log.Println(t)
			close(stopCh)
		}
	}()
	//生产者
	for i := 0; i < 30; i++ {
		go func(i int) {
			for {
				id := strconv.Itoa(i)
				value := T(rand.Intn(100))
				if value == 99 {
					select {
					case toStop <- "sender# id:" + id + "to close":
					default:
					}
				}
				select {
				case <-stopCh:
					return
				default:
				}
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(i)
	}
	//消费者
	for i := 0; i < 20; i++ {
		go func(i int) {
			id := strconv.Itoa(i)
			for {
				select {
				case <-stopCh:
					return
				default:
				}
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == 98 {
						select {
						case toStop <- "receiver# id:" + id + "to close":
						default:
						}
					}
					log.Println("receiver value :", value)
				}
			}
		}(i)
	}
	time.Sleep(10 * time.Second)
}
