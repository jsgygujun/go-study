package main

import (
	"fmt"
	"time"
)

func fun1() {
	msg := make(chan string)
	go func() {
		msg <- "ping"
	}()
	s := <-msg
	fmt.Println(s)
}

func bufferedChannel() {
	msg := make(chan string, 2) // 带有缓冲的channel
	msg <- "buffered"
	msg <- "channel"
	fmt.Println(<-msg)
	fmt.Println(<-msg)
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second * 1)
	fmt.Println("done")
	done <- true
}

func syncChannel() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// 有方向的 channel
func dirChannel() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func main() {
	fun1()
	bufferedChannel()
	syncChannel()
	dirChannel()
}
