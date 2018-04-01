package main

import "fmt"

func Insertion(data []string) []string {
	var tmp string

	for i := 1; i < len(data); i++ {
		tmp = data[i]
		for j := i - 1; j >= 0; j-- {
			if data[j] < tmp {
				break
			}
			data[j+1] = data[j]
			data[j] = tmp
		}
	}
	return data
}

func main() {
	data := []string{"f", "e", "d", "c", "b", "a"}
	fmt.Println(Insertion(data))
}
