package main

import (
    "fmt"
    "time"
)

func main() {
    timeout := make(chan bool, 1)
    go func() {
        time.Sleep(2 * time.Second)
        timeout <- true
    }()
    ch := make(chan int)

	// Create a timer which can be canceled.
	timer := time.NewTimer(time.Second * 1)
	defer timer.Stop()

    select {
		case <-ch:
		case <-timeout:
			fmt.Println("Open5GS")
		case <-timer.C:
			fmt.Println("free5GC")
    }
}