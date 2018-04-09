package main

import "fmt"

func main()  {
	a()
	b()
	c()
}

func a(){
	fmt.Println( `func a` )
}


func b(){
	fmt.Println( `func b` )
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println( err )
		}
	}(  )
	panic( `PANIC B` )
}

func c(){
	fmt.Println( `func c` )
}
