package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ticker := time.NewTicker(3 * time.Second)
	fmt.Println("当前时间为:", time.Now())
	count := 1
	sy := sync.WaitGroup{}
	sy.Add(2)
	is_exit := false
	go func() {
		sy.Done()
		for {
			t := <-ticker.C
			fmt.Println("当前时间为:", t)

			count++
			if count == 10 {
				is_exit = true
				ticker.Stop()
			}
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			if is_exit == true {
				fmt.Println("协程可以正常退出")
				sy.Done()
			}
		}
	}()

	sy.Wait()
	fmt.Println("exit success")
}