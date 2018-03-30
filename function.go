package main

import "fmt"

func main()  {
	/*a := func() {
		fmt.Println( `匿名函数` )
	}

	a()*/


	//闭包
	sum := clou( 4 )( 5 )
	fmt.Println( sum )
}


func clou( x int ) func( y int ) ( int ) {
	fu := func( y int ) int {
		return x + y
	}

	return fu

}

//func main()  {
//	a := 1
//	A(&a)
//	fmt.Println( a )
//}
//
//func A(a *int)  {
//	*a = 2
//	fmt.Println( *a )
//}