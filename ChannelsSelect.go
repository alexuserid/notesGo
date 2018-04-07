package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	c, a, ch1 := make(chan bool), make(chan bool), make(chan int)

	go func() {
		for i := arr[len(arr)-1] + 1; i <= 10; i++ {
			ch1<- i
		}
		c<- true
	}()

	go func() {
		for {
			select {
			case input := <-ch1:
				arr = append(arr, input)
				fmt.Println(arr)
			case <-c:
				a<- true //channel a guarantees that all massives will be printed
				return
			}
		}
	}()

	<-a
}
