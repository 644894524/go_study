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
	//timehour := stime.Unix()
	//fmt.Println( timehour.Format("02/01/2006 15:04:05 PM") )



	//时间戳，时间相互转换
	timestamp := time.Now().Unix() - 3600 * 24	//时间戳
	fmt.Println( timestamp )

	etime := time.Unix( timestamp, 0 )
	fmt.Println( etime.Format( "2006-01-02 15:04:05" ) )	//对时间戳进行格式化


	//字符串转换时间戳
	t, _ := time.Parse("2006-01-02 15:04:05", "2018-06-26 11:01:01")
	currtime := t.Unix()
	fmt.Println( currtime )
}


