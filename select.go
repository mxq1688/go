package main

import (
	"fmt"
	"time"
)

//select在主线程和携程中的表现
/*
	golang for的select体中如何break外层循环：
		1.1 使用带标签的break
			LOOP: for {
				select{
					case <-tick.C:
						//do someting
					case <- stop:
						break LOOP//break的for循环,跳转执行doOther()
				}
				doNext()
			}
			doOther()
		1.2	使用return
			for {
				select{
					case <-tick.C:
						//do someting
					case <- stop:
						return //干净利落，适合退出goroutin的场景
				}
				doNext()
			}
		1.3	使用标志
			isStop := false
			for {
				select{
					case <-tick.C:
						//do someting
					case <- stop:
						isStop = true//干净利落，适合退出goroutin的场景
						break
				}
				if isStop {
					break
				}
				doNext()
			}
			doOther()
		1.4 goto
			for {
				select{
					case <-tick.C:
						//do someting
					case <- stop:
						goto LOOP
				}
				doNext()
			}
			LOOP:
			doOther()
*/

type T int

func main() {
	// test1()
	test2()
}
func test1() {
	c := make(chan T)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("关闭")
		close(c)
	}()
	for {
		select {
		case <-c:
			fmt.Println(true)
			return //主线程中必须加return（或者上面的四种方法）退出循环
		default:
			fmt.Println(false)
		}
	}
}

func test2() {
	c := make(chan T)
	go func() {
		for {
			select {
			case <-c:
				fmt.Println(true)
				// return //携程中可以不加也会退出循环
			default:
				fmt.Println(false)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("关闭")
	close(c)
}
