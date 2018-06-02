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

func main() {
	var nf = new(float)
	*nf = 2.3
	var ni = new(intn)
	*ni = 6

	nf.addOne()
	ni.addOne()
	fmt.Println(*nf)
	fmt.Println(*ni)
}
