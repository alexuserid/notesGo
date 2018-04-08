package main

import "fmt"

func Insertion(data []string) {
	var tmp string

	for i := 1; i < len(data); i++ {
		tmp = data[i]
		for j := i - 1; j >= 0; j-- {
			if data[j] < tmp {
				break
			}
			data[j], data[j+1] = tmp, data[j]
		}
	}
}

func main() {
	data := []string{"f", "e", "d", "c", "b", "a"}
	Insertion(data)
	fmt.Println(data)
}
