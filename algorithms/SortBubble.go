package main

import "fmt"

func Bubble(data []string) {

	for i := len(data) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func main() {
	data := []string{"f", "e", "d", "c", "b", "a"}
	Bubble(data)
	fmt.Println(data)
}
