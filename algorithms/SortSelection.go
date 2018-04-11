package main

import "fmt"

func Selection(data []string) {
	for i:=0; i<len(data)-1; i++ {
		min := i
		
		for j:=i+1; j<len(data); j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}
}

func main() {
	data := []string{"f", "e", "d", "c", "b", "a"}
	Selection(data)
	fmt.Println(data)
}
