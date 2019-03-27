package main

import "fmt"

func main(){
	defer func() {
		fmt.Println( "终止执行" )
	}()

	var arr = [...]int{ 1, 3, 5, 7, 9 }
	for k, v := range arr {
		fmt.Println( k,v )
	}


}