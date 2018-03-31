package main

import "fmt"

func Bubble(data []string) []string {
	var tmp string

	for i := len(data) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if data[j] > data[j+1] {
				tmp = data[j]
				data[j] = data[j+1]
				data[j+1] = tmp
			}
		}
	}
	return data
}

func main() {
	data := []string{"f", "e", "d", "c", "b", "a"}
	fmt.Println(Bubble(data))
}
