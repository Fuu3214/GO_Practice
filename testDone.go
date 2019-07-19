package main

import (
	"fmt"
	"time"
)

func testDone() {
	done := make(chan struct{})
	for i := 0; i < 2; i++ {
		go func(i int) {
			for {
				select {
				default:
					fmt.Println("working")
					time.Sleep(time.Millisecond * 500)
				case <-done:
					fmt.Println("BOOM!")
					return
				}
			}
		}(i)
	}
	time.Sleep(time.Second * 2)
	close(done)
	time.Sleep(time.Second * 5)
}
