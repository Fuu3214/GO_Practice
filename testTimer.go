package main

import (
	"fmt"
	"math/rand"
	"time"
)

func testTimer() {
	for {
		select {
		case <-time.After(getElectionTimeOut()):
			fmt.Println("?")
		}
	}
}

func getElectionTimeOut() time.Duration {
	randomTimeout := 0 + rand.Intn(2000)
	fmt.Println(randomTimeout)
	electionTimeout := time.Duration(randomTimeout) * time.Millisecond
	return electionTimeout
}
