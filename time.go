package main

import (
	"time"
	"fmt"
)

func main(){
	stime := time.Now()		//获取时间
	fmt.Println( stime )
	y, m, d := stime.Date()
	h, min, s := stime.Clock()
	fmt.Printf( "%d %02d %02d %02d:%02d:%02d \r\n", y, m, d, h, min, s )		//格式化时间


	//time.Sleep( 1000000 * time.Microsecond )		//毫秒
	time.Sleep( 1 * time.Second )
	fmt.Println( `haha` )


	//时间戳
	fmt.Println( stime.Unix() )


	//前一个小时

}


