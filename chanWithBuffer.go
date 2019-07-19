package main

import (
	"fmt"
	"sync"
)

func chanWithBuffer() {
	ch := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				ch <- i*10 + j
			}
			fmt.Println("done")
		}(i)
	}
	wg.Add(4)

	go func() {
		fmt.Println("wait")
		wg.Wait()
		fmt.Println("ready to close")
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
