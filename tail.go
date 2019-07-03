package main

import (
	"github.com/hpcloud/tail"
	"fmt"
)

func main(){

	t, err := tail.TailFile("/tmp/code_auto_push_0601.log", tail.Config{ Follow: true, Poll: true } )
	if( err == nil ) {
		for line := range t.Lines {
			fmt.Println(line.Text)
		}
	}

}
