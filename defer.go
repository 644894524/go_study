package main

import (
	"fmt"
)

func recoverName() {
	if r := recover(); r!= nil {
		fmt.Println("recovered from ", r)
	}
}


func main(){
	//延迟函数，在报错之前执行；程序终止时执行
	defer func() {
		fmt.Println( "终止执行" )
	}()

	//错误处理
	defer recoverName()


	bb := make([]int,10)
	fmt.Println( bb )

	var map_name = map[string]string{}
	map_name["one"] = "one"
	map_name["two"] = "two"
	for k, v := range map_name {
		fmt.Println( k + "=》" + v )
	}


	var cc = make( map[int]string )
	cc[0] = "haah"
	cc[1] = "xixi"
	fmt.Println( len(cc) )

	var arr = [...]int{ 1, 3, 5, 7, 9 }
	for k, v := range arr {
		fmt.Println( k,v )
	}

	panic( "this is error" )
	fmt.Println( "还会继续执行吗" )			//panic 终止程序，所以不会有输出了
}