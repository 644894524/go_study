package main

import "fmt"

type books struct {
	title string
	author string
	price float32
	id int
}

func main(){
	var book1 books
	book1.title = `php`
	book1.author = `sl`
	book1.price = 15.50
	book1.id = 1

	var str string = readbook( book1 )
	fmt.Println( str )
}

func readbook( book books ) string {
	var str string = fmt.Sprintf( `ID:%d, this is %s, It's author %s, It's price %0.1f`, book.id, book.title, book.author, book.price )
	return str
}