package main

import "fmt"

func sum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	fmt.Println(arr)
	x := len(arr) - 1
	return arr[x] + sum(arr[:x])
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(sum(arr))
}
