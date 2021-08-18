package main

import (
	"fmt"
)

//import (
//	f "fmt"//别名
//	. "fmt" //使用fmt中的方法时可以省略fmt
//	_ "test1"//只执行init方法
//)

/*
包引用：
如果是同一个包下的：不要求大写
	只要方法名或者结构首字母大写，直接可以调用，不必引用包？ go run main.go hello.go
引用同级子目录下的包：要引用的变量，常量，函数，结构体等首字母需要大写
	引入子目录名，既目录下文件的包名

init函数：
init函数会先运行，引入的包文件中的init函数优先于当前init函数

常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：

{ 不能在单独的行上

*/

/* 声明全局变量 */
var g int

func init() {
	fmt.Println("this is main init")
}
func main() {
	// 使用自定义包
	// test1.Say()

	/*
		1、变量申明 var a type  #必须加类型, var a,b,c type 可申明多个
		2、申明赋值 var a int=0 #可以不加类型，根据值判定
		3、省略var v_name := value #只能在函数内使用，_回收变量
	*/
	// a, _, c := 1.2, 2, 3 //_回收变量
	// var d = int(a)
	// fmt.Println(a, c, d, reflect.TypeOf(d))

	/*
		常量申明 const identifier [type] = value #可以不加类型
		常量用作枚举：
			const (
				Unknown = 0
				Female = 1
				Male = 2
			)
		iota: 特殊的常量，可以认为是一个可以被编译器修改的常量。
				iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引
	*/
	// const ( //const中每新增一行常量声明将使iota计数一次
	// 	a1 = iota
	// 	a2 = iota
	// 	_  //跳值使用法
	// 	a3 = iota
	// )
	// fmt.Println(a1, a2, a3) //0 1 3
	// const (
	// 	a1 = iota
	// 	a2 = 3.14 //插队使用法
	// 	a3 = iota
	// )
	// fmt.Println(a1, a2, a3) //0 3.14 2
	// const (
	// 	a1 = iota * 2
	// 	a2 //隐式使用法。省略iota，会继承上一个非空表达式
	// 	a3
	// )
	// fmt.Println(a1, a2, a3) //0 2 4
	// const (
	// 	a1, a2 = iota, iota + 3 //同一行iota不会增加
	// 	b1, b2                  //b1继承a1, b2继承a2
	// 	f      = iota
	// )
	// fmt.Println(a1, a2, b1, b2, f) //0 3 1 4 2

}
