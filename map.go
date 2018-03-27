package main

import (
	"fmt"
	"sort"
)

func main(){
	//创建map
	m1 := make( map[string]string )
	m1[`name`] = `sunlong`
	m1[`age`] = `18`

	fmt.Println( m1 )


	//初始化，赋值
	m2 := map[string]string {
		`name`:`zhangsan`,
		`age`:`28`,
	}


	//循环输出map
	for key := range m2 {
		fmt.Println( key + "=>" + m2[key] )
	}

	//查看元素在集合中是否存在
	captial ,isexist := m2[`name`]	//这里不知道原因
	if isexist {
		fmt.Println( `存在,值是`, captial )
	}else{
		fmt.Println( `不存在` )
	}

	delete( m2, `age` )
	fmt.Println( m2 )
	//haha2,haha3 := 2,3
	//fmt.Println( haha2, haha3 )


	m := map[string]int{
		"unix":         0,
		"python":       1,
		"go":           2,
		"javascript":   3,
		"testing":      4,
		"philosophy":   5,
		"startups":     6,
		"productivity": 7,
		"hn":           8,
		"reddit":       9,
		"C++":          10,
	}

	//是随机的
	for mk := range m {
		fmt.Println( mk, "=>", m[mk] )
	}

	//按照顺序输出
	var mkeys []string
	for mk := range m {
		mkeys = append( mkeys, mk )
	}

	//go 排序 ,
	//go 分别提供了 sort.Ints() 、 sort.Float64s() 和 sort.Strings() 函数， 默认都是从小到大排序。
	//


	fmt.Println( mkeys )

}
