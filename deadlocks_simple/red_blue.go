package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock1 = sync.Mutex{}
	lock2 = sync.Mutex{}
)

func blueRobot() {
	for {
		println("Blue: Acquiring lock1")
		lock1.Lock()
		println("Blue: Acquiring lock2")
		lock2.Lock()
		println("Blue: Both locks Acquired")
		lock1.Unlock()
		lock2.Unlock()
		println("Blue: Both locks Released")
	}
}
func redRobot() {
	for {
		println("Red: Acquiring lock2")
		lock2.Lock()
		println("Red: Acquiring lock1")
		lock1.Lock()
		println("Red: Both locks Acquired")
		lock1.Unlock()
		lock2.Unlock()
		println("Red: Both locks Released")
	}
}

func main() {
	go blueRobot()
	go redRobot()

	time.Sleep(20 * time.Second)
	fmt.Println("Done")
}
