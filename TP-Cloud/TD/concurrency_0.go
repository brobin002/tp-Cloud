package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		//etape 1 time.Sleep(100 * time.Millisecond)
		//etape 2 time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}