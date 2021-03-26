package main

import "fmt"

//定义结构体
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func definition() {
	fmt.Println(Books{"Go", "zzz", "GoGo", 2777})
	fmt.Println(Books{title: "Go", author: "zzz", subject: "GoGo", book_id: 2777})
	fmt.Println(Books{author: "zzz", book_id: 2777})
}

//访问结构体成员
func visit() {
	var Book1 Books // 声明Book1 为 Books 得类型
	var Book2 Books

	Book1.title = "Go"
	Book1.author = "zzz"
	Book1.subject = "GoGO"
	Book1.book_id = 2777

	Book2.title = "PHP"
	Book2.author = "zz"
	Book2.subject = "PHP PHP"
	Book2.book_id = 22777

	// 打印Book1 信息
	fmt.Printf("Book 1 title %s\n", Book1.title)

	//打印Book2 信息
	fmt.Printf("Book 2 title %s\n", Book2.title)

}

//结构体作为函数参数
func structureParameter() {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go"

	Book2.title = "PHP"

	printBook(&Book1)
	printBook(&Book2)

}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	if len(book.author) > 0 {
		fmt.Printf("Book author : %s \n", book.author)
	}
	if len(book.subject) > 0 {
		fmt.Printf("Book subject : %s \n", book.subject)
	}
	if book.book_id > 0 {
		fmt.Printf("Book book_id : %s \n", book.book_id)
	}
}

func main() {
	definition()
	visit()
	structureParameter()
}
