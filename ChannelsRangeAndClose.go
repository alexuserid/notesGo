package main

import "fmt"

func chf(ch chan int, arr []int) {
	for _, v := range arr {
		ch<-v
	}
	close(ch)
}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	ch := make(chan int, 10)

	go chf(ch, arr)
	for c := <-ch; c > 0; c = <-ch {
		fmt.Println(c)
	}
}
