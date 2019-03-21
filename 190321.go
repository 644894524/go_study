package main

import (
	"fmt"
	"strconv"
)

func main(){
	//字符串声明
	var str string = `this is string type %d`
	var num int = 8
	var format string = fmt.Sprintf( str, num )
	fmt.Println( format )


	//数组声明
	var arr = [...]int{1, 3, 5, 7}
	fmt.Println( arr )

	//切片, 切片尾部不能超过数组长度
	//fmt.Println( arr[1:9] )		//invalid slice index 9 (out of bounds for 4-element array)
	fmt.Println( arr[1:3] )

	//最外层必须声明数组长度,多维数组
	var darr = [...][2]int{
		{1,3},
		{3,5},
		}
	fmt.Println( darr )

	//字符串转整型
	var fstr string = `1234`
	stoi,_ := strconv.Atoi( fstr )
	fmt.Println( stoi )

	//整型转字符串
	itos := strconv.Itoa( stoi )
	fmt.Println( itos )

	//函数声明以及使用
	var he int = umath( 1, 5 )
	fmt.Println( he )

	//函数作为返回值，闭包
	ufun := ufuncmath( 1, 5 )
	var qiuhestr string = ufun()
	fmt.Println( qiuhestr )
}

func umath( num1 int, num2 int ) int {
	return num1 + num2
}


func ufuncmath( num1 int, num2 int ) func() string {
	ufun := func() string {
		var qiuhe int = num1 + num2
		return fmt.Sprintf( "jie guo shi %d", qiuhe )
	}

	return ufun
}