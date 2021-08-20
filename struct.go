package main

import "fmt"

//定义结构体
type Books struct {
	book_id int
	uid     int
	title   string
	content string
}

func main() {
	// 创建一个新的结构体
	fmt.Println(Books{1, 2, "i am title", "i am content"})
	// 也可以使用 key => value 格式
	fmt.Println(Books{book_id: 1, uid: 2, title: "i am title", content: "i am content"})

	//用于变量的申明
	var books1 Books = Books{1, 2, "i am title", "content"}

	// 访问结构体成员
	books1.title = "haha"
	fmt.Println(books1)
	fmt.Println(books1.title)

	//结构体作为函数参数
	fmt.Println(printBooks(books1))

	//结构体指针
	var struct_pointer *Books
	struct_pointer = &books1
	fmt.Println(struct_pointer, &books1, *&books1)

	//结构体指针作为函数参数，传递的是引用，会影响外面的值
	fmt.Println(pritBooks2(&books1), books1)
}

func printBooks(book Books) string {
	return book.title
}
func pritBooks2(book *Books) string {
	fmt.Println((*book).title, book.title) //内部访问可以直接访问，或者*book
	book.title = "change"
	return book.title
}
