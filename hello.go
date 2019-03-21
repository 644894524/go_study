package main

import "fmt"

func main()  {
	fmt.Println( "hello world" )

	//数组声明
	var a = [...]int{ 1, 2, 3, 4, 5 }
	fmt.Println( a )


	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] =3
	fmt.Println( arr )
}


