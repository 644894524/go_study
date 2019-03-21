package main

//go 语言控制结构

import (
	"fmt"
	"math"
)

const (
	name   string = `sunlong`
	age    int    = 10
	height int    = 180
)

func main() {
	//fmt.Print( name )
	//var age int = 10;

	//使用函数
	control_if()

	var score = control_switch(20)
	fmt.Print("\r\n 你的结果是：", score)

	control_for()
}

func control_if() {
	//不需要使用(), 使用也无所谓，执行go fmt 会去掉
	if age >= 10 && name != `sunlong` {
		fmt.Printf(`my name is %s, my age is %d`, name, age)
	}

	//不能够去掉（）
	if (age >= 10 && name == `sunlong2`) || (height >= 180) {
		fmt.Printf(`my name is %s, my age is %d`, name, age) //注意 age 是整形 ，不能使用 %s
	}
}

func control_switch(score float64) string {
	var res float64 = math.Floor(score / 10) //函数返回浮点值 , 必须都是浮点只
	var resstr string
	switch res {
	case 7:
		resstr = `合格`

	case 9:
		resstr = `优秀`
	default:
		resstr = `不合格`
	}

	return resstr
}

func control_for() {

	//for 循环几种写法
	var i int = 1
	for i <= 10 {
		fmt.Print("\r\n 第", i, "个")
		i++
	}

	for v := 10; v <= 20; v++ {
		fmt.Print("\r\n v : ", v)
	}

	//return a
}
