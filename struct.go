package main

import "fmt"

//定义结构体
type Books struct {
	book_id int
	uid     int
	title   string
	content string
}

func printBooks(book Books) string {
	return book.title
}
func pritBooks2(book *Books) string {
	fmt.Println((*book).title, book.title) //为什么是一样的 结构体访问数据是一样的
	book.title = "change"
	return book.title
}
func swap2(a *int, b *int) (int, int) {
	var tem int
	tem = *a
	*a = *b
	*b = tem
	return *a, *b
}

func main() {
	//结构体
	fmt.Println(Books{1, 2, "i am title", "i am content"})
	fmt.Println(Books{book_id: 1, uid: 2, title: "i am title", content: "i am content"})
	var books1 Books = Books{1, 2, "i am title", "content"}
	books1.title = "haha"
	fmt.Println(books1)
	fmt.Println(books1.title)

	//结构体作为函数参数
	// fmt.Println(printBooks(books1))

	//结构体指针
	// var struct_pointer *Books
	// struct_pointer = &books1
	// fmt.Println(struct_pointer, &books1, *&books1)
	// fmt.Println(pritBooks2(&books1), books1)
}
