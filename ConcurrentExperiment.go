// A concurrent prime sieve

package main

func producer(ch chan int) {
	ch <- 233
	close(ch)
}
