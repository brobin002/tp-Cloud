package main
import (
    "fmt"
    "time"
)

func main() {
    // create new channel of type int
    ch := make(chan int)

    // start new anonymous goroutine
    go func() {
        time.Sleep(1 * time.Second)
        // send 42 to channel
        ch <- 42
    }()
    // read from channel
    v := <-ch
    fmt.Println(v)
    fmt.Println("main thread") 
}
