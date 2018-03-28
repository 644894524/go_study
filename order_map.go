package main

import (
	"fmt"
	"sort"
)

func main(){

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

	//按照顺序输出
	var mkeys []string
	for mk := range m {
		mkeys = append( mkeys, mk )
	}


	//fmt.Println( mkeys )

	//排序
	/*
	sort.Strings( mkeys )
	for _,mv := range mkeys {
		fmt.Println( mv, `=>`, m[mv] )
	}
	*/

	//降序排
	sort.Sort( sort.Reverse( sort.StringSlice( mkeys ) ) )
	fmt.Println( mkeys )
	for _,mv := range mkeys {
		fmt.Println( mv, `=>`, m[mv] )
	}
}