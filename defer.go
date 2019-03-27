package main

import "fmt"

func recoverName() {
	if r := recover(); r!= nil {
		fmt.Println("recovered from ", r)
	}
}


func main(){
	defer func() {
		fmt.Println( "终止执行" )
	}()

	defer recoverName()

	var arr = [...]int{ 1, 3, 5, 7, 9 }
	for k, v := range arr {
		fmt.Println( k,v )
	}


	panic( "this is error" )
	fmt.Println( "还会继续执行吗" )			//panic 终止程序，所以不会有输出了
}