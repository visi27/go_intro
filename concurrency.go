package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id int
}

func main() {
	c := make(chan int)
	for i := 0; i < 4; i++ {
		worker := Worker{id: i}
		go worker.process(c)
	}

	for i := 0; i < 2000; i++ {
		select {
		case c <- rand.Int():
			fmt.Println("RECEIVED")
		case <-time.After(time.Millisecond * 50):
			fmt.Println("timed out *********************************************************************")
			//default:
			//	fmt.Println("DROPPED")
		}

		//fmt.Printf("Buffer length %d\n", len(c))
		time.Sleep(time.Millisecond * 100)
	}

	time.Sleep(time.Millisecond * 10000)
}

func (w Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 300)
	}
}
