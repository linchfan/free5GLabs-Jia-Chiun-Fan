package main

import (
    "fmt"
    "sync"
)

type Counter struct {
    i int // integer
    m sync.Mutex // mutex
}

func (c *Counter) Increment() {
    c.m.Lock() // Lock
    c.i += 1 // i++
    c.m.Unlock() // Unlock
}

func main() {
    var c Counter
    var wg sync.WaitGroup

    for i := 0; i < 100000; i++ {
        wg.Add(1)

        go func() {
            c.Increment()
            wg.Done()
        }()
    }

    wg.Wait()

    fmt.Println(c.i)
}
