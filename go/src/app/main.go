package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	var val1 int
	for i := 1; i < x1; i++ {
		val := rand.Intn(1000 - 1)
		fmt.Fprintf(w, "\n sda%d=%d", i, val)

	}
	for i := 1; i < x2; i++ {
		val := rand.Intn(1000 - 1)
		fmt.Fprintf(w, "\n eth%d=%d", i, val)
	}
	val1 = rand.Intn(1000-1) * 1024 * 1024
	fmt.Fprintf(w, "\n Очередь const_mq1_q1 сообщений  %d", val1)
	val1 = rand.Intn(1000-1) * 1024 * 1024
	fmt.Fprintf(w, "\n Очередь const_mq2_q1 сообщений  %d", val1)
	val1 = rand.Intn(1000-1) * 1024 * 1024
	fmt.Fprintf(w, "\n Очередь const_mq3_q1 сообщений  %d", val1)
}

var x1, x2 int

func main() {
	min := 1
	max := 20
	//en := make(chan int)
	//en2 := make(chan int)
	go func() {
		for {
			x1 = rand.Intn(max-min) + min
			time.Sleep(time.Hour * 3)
		}

	}()
	//x1 = <-en
	go func() {
		for {
			x2 = rand.Intn(max-min) + min
			time.Sleep(time.Hour * 3)
		}
	}()
	//x2 = <-en2
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
