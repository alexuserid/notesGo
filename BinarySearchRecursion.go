package main

import "fmt"

func finder(data []int, n int) int {
	m := len(data) / 2

	if n < data[0] || n > data[len(data)-1] {
		fmt.Println("The number is out of range")
		return 0

	} else if n < data[m] {
		fmt.Println(data[:m])
		return finder(data[:m], n)

	} else if n > data[m] {
		fmt.Println(data[m:])
		return finder(data[m:], n)

	} else if n == data[m] {
		return data[m]
	}

	return 0
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	res := finder(data, 11)

	if res == 0 {
		fmt.Println("Something goes wrong")
		return
	}

	fmt.Println(res)
}
