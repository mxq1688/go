package main

import (
	"fmt"
	"math"
)

/* 定义结构体 */
type Circle struct {
	radius float64
}

// 声明一个函数类型
type cb func(int) int

func main() {
	//函数调用
	var m = max(1, 2)
	fmt.Println(m)

	//返回多个值
	x, y := swap("google", "baidu")
	fmt.Println(x, y)

	//闭包
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()
	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	//函数方法 看看接口的实现
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())

	/* 声明函数变量 及使用*/
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	// 函数作为参数传递，实现回调
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})

	// 函数参数返回形式

}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

//交换
func swap(x string, y string) (string, string) {
	return y, x
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func testCallBack(x int, f cb) {
	f(x)
}
func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}
