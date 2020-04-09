package main

import "fmt"

//import (
//	f "fmt"//别名
//	. "fmt"
//	_ "mmm"//只执行init方法
//)

//大写字母开头的变量可在其他包使用
//小写字母开头的变量不能在其他包使用
//常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：

//定义结构体
type Books struct {
	book_id int
	uid     int
	title   string
	content string
}

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

	// test1.Feo()
	//a,_,c := 1.2,2,3//只能在函数内使用  _回收变量
	//var d = int(a)
	//fmt.Println(a,c,d, reflect.TypeOf(d))

	//iota碰到const关键字时重置为0
	//const x = iota;
	//const y = iota;
	//fmt.Println(x, y)

	// const ( //const中每新增一行常量声明将使iota计数一次
	// 	a1 = iota
	// 	a2 = iota
	// 	_  //跳值使用法
	// 	a3 = iota
	// )
	// fmt.Println(a1, a2, a3)
	// const (
	// 	a1 = iota
	// 	a2 = 3.14 //插队使用法
	// 	a3 = iota
	// )
	// fmt.Println(a1, a2, a3)
	// const (
	// 	a1 = iota * 2
	// 	a2 //隐式使用法。省略iota，会继承上一个非空表达式
	// 	a3
	// )
	// fmt.Println(a1, a2, a3)

	// const (
	// 	a1, a2 = iota, iota + 3 //同一行iota不会增加
	// 	b1, b2                  //b1继承a1, b2继承a2
	// 	f      = iota
	// )
	// fmt.Println(a1, a2, b1, b2, f)

	//if a< 3 {
	//
	//}
	//switch value {
	//case value1:
	//
	//}
	//for 循环的 range 格式可以对 slice（切片）、map、数组、字符串等进行迭代循环。格式如下：
	//for key, value := range oldMap {
	//	newMap[key] = value
	//}

	//函数定义
	//var m = max(1, 2)
	//fmt.Println(m)
	//x,y := swap("google", "baidu")
	//fmt.Println(x, y)

	//数组
	// var arr = [5]int{1, 2, 3, 4, 5}
	// arr := []int{1, 2, 3, 4, 6}
	// var arr1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	// var arr2 []int //定义数组
	// fmt.Println(arr, arr1, arr2)

	//指针
	// var m = 123
	// var o *int = &m
	// var oo **int = &o
	// var ooo ***int = &oo
	// fmt.Println(m, o, *o, oo, *oo, **oo, ooo, ***ooo) //指向指针的指针

	//指针数组
	// var arr = []int{1, 2, 3}
	// var arr_ptr [3]*int
	// arr_ptr[0] = &arr[0]
	// arr_ptr[1] = &arr[1]
	// arr_ptr[2] = &arr[2]
	// fmt.Println(arr, arr_ptr, *arr_ptr[0])

	//指针函数  	//函数参数传引用
	//var a,b = 1,2
	//c,d := swap2(&a, &b);
	//fmt.Println(a,b, c,d)

	//结构体
	// fmt.Println(Books{1, 2, "i am title", "i am content"})
	// fmt.Println(Books{book_id: 1, uid: 2, title: "i am title", content: "i am content"})
	// var books1 Books = Books{1, 2, "i am title", "content"}
	// books1.title = "haha"
	// fmt.Println(books1)
	// fmt.Println(books1.title)

	//结构体作为函数参数
	// fmt.Println(printBooks(books1))

	//结构体指针
	// var struct_pointer *Books
	// struct_pointer = &books1
	// fmt.Println(struct_pointer, &books1, *&books1)
	// fmt.Println(pritBooks2(&books1), books1)

	//Slice 切片
	// var identifier_nil []int          //定义空切片
	// identifier := make([]int, 10, 20) //创建切片
	// var slice1 []int = make([]int, 2)
	// s := identifier[:]  //初始化切片s,是数组identifier的引用
	// m := identifier[:5] //初始化切片s,是数组identifier的引用
	// fmt.Println(identifier, slice1, s, m, len(s), cap(s))
	// fmt.Println(identifier_nil == nil)

	//numbers1 := make([]int,0,5)
	//fmt.Println(numbers1, len(numbers1), cap(numbers1))

	//切片append() copy()
	//var numb []int
	//numb = []int{1,2,3}
	//num1 := append(numb, 5,4,3)
	//num2 := make([]int, len(num1), cap(num1)*2)
	//copy(num2, num1)
	//fmt.Println(numb, num1, num2)

	//map  如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
	//var mMap map[string]string //定义空map  不能存放键值对
	//var nMap = make(map[string]string)//创建map并初始化
	//nMap ["title"] = "ajeif"
	//nMap ["content"] = "faeoifjqje"
	//nMap ["name"] = "mxq"
	//fmt.Println(mMap,nMap)

	//for key,value := range nMap {
	//	fmt.Println(key, value)
	//}

	//判断map中元素是否存在
	//mkey, mvalue := nMap["uid"]
	//fmt.Println(mkey, mvalue)

	//map delete()
	/* 创建map */
	//countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	//fmt.Println("原始地图", countryCapitalMap)
	//for key, value := range countryCapitalMap {
	//	fmt.Println(key, value, countryCapitalMap[key])
	//}
	////fmt.Println("\n")
	//delete(countryCapitalMap, "France")
	//for key, value := range countryCapitalMap {
	//	fmt.Println(key, value, countryCapitalMap[key])
	//}

	//接口
	// var phone Phone
	//
	// phone = new(NokiaPhone)
	// phone.call()
	//
	// phone = new(IPhone)
	// phone.call()

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

func ml() (mm int) {
	t := 10
	return t //会把t赋值给mm
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
func pritBooks2(book *Books) string {
	fmt.Println((*book).title, book.title) //为什么是一样的 结构体访问数据是一样的
	book.title = "change"
	return book.title
}

func printBooks(book Books) string {
	return book.title
}

func swap2(a *int, b *int) (int, int) {
	var tem int
	tem = *a
	*a = *b
	*b = tem
	return *a, *b
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

func swap(x string, y string) (string, string) {
	return y, x
}
