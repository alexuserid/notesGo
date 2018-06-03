package main

import "fmt"

type number interface {
	addOne()
}

type float float64
func (n *float) addOne() {
	var one float = 1.0
	*n += one
}

type intn int64
func (n *intn) addOne() {
	var one intn = 1
	*n += one
}

func add(n number) {
	n.addOne()
}

func main() {
	var nfp = new(float)
	*nfp = 2.3
	var nip = new(intn)
	*nip = 5
	fmt.Println(*nfp)
	fmt.Println(*nip)

	add(nfp)
	add(nip)
	fmt.Println(*nfp)
	fmt.Println(*nip)
}
