package main

import (
	"fmt"
	"strconv"
)

func main(){

	/*
	数组操作
	 */
	//var arr = [5]int{1,3,5,7}

	//没有初始值默认为0
	//var arr [5]int
	//arr[0] = 1


	//转换
	var sarr = [...]string{
		strconv.Itoa(1 ), strconv.Itoa(3 ), strconv.Itoa(5 ),
	}
	fmt.Print( sarr, "\r\n" )
	//strconv.Atoi()	//字符串转整型

	//声明变长数组, 数组必须有初始值   cannot convert "Mon" (type untyped string) to type int
	var arr = [...]string{
		`Mon`, `Tue` , `Wed`, `Thu`, `Fri`, `Sat`, `Sun`, //TODO 这里必须加, 否则会报错、或者直接 }
	}



	for day := range arr {
		fmt.Print( arr[ day ], "\r\n" )
	}


	//多维数组
	var multiarr = [2][3]int{
		{ 1, 3, 5 },
		{ 2, 4, 6 },
		//{ 7, 9 , 10 } //, 这里会报错  array index 2 out of bounds [0:2]
	}

	fmt.Print( len( multiarr ), "\r\n" )

	for iarr := range multiarr {
		var sum int
		for varr := range multiarr[ iarr ] {
			sum += multiarr[ iarr ][ varr ]
		}
		fmt.Print( sum, "\r\n" )
	}


	/*
	切片
	 */
	var arr2 = [5]int{ 1, 3, 5, 7, 9 }

	//fmt.Print( arr2[min:max] )	//todo 包括min， 但是不包括max
	fmt.Print( arr2[0:1], "\r\n" )
	fmt.Print( arr2[2:], "\r\n" )
	//fmt.Print( arr2[:6] )	//报错 ./0311.go:57:17: invalid slice index 6 (out of bounds for 5-element array)
	fmt.Print( arr2[:5] )
}
