package main

import (
	"fmt"
	"time"
)

func  main()  {
	currtime := time.Now()		//当前时间
	fmt.Println( currtime.Unix() )			//转换时间戳
	tmp := currtime.Format( "2006-01-02 15:04:05" )		//格式化时间
	fmt.Println( tmp )


	//时间戳转时间
	ctime := currtime.Unix()
	str_time := time.Unix( ctime, 0 ).Format( "2006-01-02 15:04:05" )
	fmt.Println( str_time )

	//字符串转时间戳
	ntime, _ := time.Parse( "2006-01-02 15:04:05", str_time )
	fmt.Println( ntime.Unix() )

	//毫秒级时间戳
	mtime := time.Now().UnixNano() / 1e6
	fmt.Println( mtime )
}


