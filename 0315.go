package main

import (
	"fmt"
)

func main()  {
	fmt.Print( "开始 \n" )


	//练习goto，以及标签		goto label的label(标签)既可以定义在for循环前面,也可以定义在for循环后面
	var i int = 0

	for {
		i++
		//fmt.Println( i )

		for i > 10 {
			goto LABEL1	//break LABEL1  会跳出到LABEL1 所在的同级块， 不是死循环
		}

	}

	fmt.Println( `label1 OK` )
LABEL1:


	var a int = 0
LABEL2:
	for {
		a++
		//fmt.Println( a )

		for a > 10 {
			break LABEL2			//也不会是死循环	必须在前边
		}
	}

	fmt.Println( `label2 ok` )

}