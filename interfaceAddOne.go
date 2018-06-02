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

func printer(n number) {
	if num, ok := n.(*float); ok {
		fmt.Println(*num)
		return
	}
	fmt.Println(*n.(*intn))
}

func main() {
	var nfp = new(float)
	*nfp = 2.3
	var nip = new(intn)
	*nip = 5

	nfp.addOne()
	printer(nfp)

	nip.addOne()
	printer(nip)
}
