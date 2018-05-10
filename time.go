package main

import (
	"time"
	"fmt"
)

func main(){
	stime := time.Now()
	fmt.Println( stime )
	y, m, d := stime.Date()
	h, min, s := stime.Clock()

	fmt.Printf( `%d %02d %02d %02d:%02d:%02d`, y, m, d, h, min, s )

}


