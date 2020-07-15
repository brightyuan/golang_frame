package main

import "fmt"

/**
引用）指针方式
*/

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books /* Declare Book1 of type Book */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* 打印 Book1 信息 */
	printBook(Book1)
	fmt.Println("===============================================")
	//指针方式
	printBookPtr(&Book1)
	fmt.Println("===============================================")

	changeBook(Book1)
	fmt.Println(Book1)
	fmt.Println("===============================================")

	//指针方式，引用方式，使用内存地址，可以修改结构体值
	changeBookPtr(&Book1)
	fmt.Println(Book1)

}
func printBookPtr(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)

}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func changeBookPtr(books *Books) {
	books.title = "test name"

}

func changeBook(books Books) {
	books.title = "test name"

}
