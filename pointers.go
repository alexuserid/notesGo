package main 

import "fmt"

//mA == memory adress
func main() {
	var n int = 6 
	fmt.Println(n, &n) //n == 6, &n == mA1

	var pn *int //pn == <nil>, &pn == mA2
	fmt.Println(pn, &pn)

	// *pn = 3
	// n = *pn

	pn = &n //n == 6, &n == mA1, *pn == 6, &pn == mA2
	fmt.Println(n, &n, *pn, &pn)

	n = 7 //n == 7, &n == mA1, pn == mA1, *pn == 7, &pn == mA2
	fmt.Println(n, &n, pn, *pn, &pn)

	*pn = 3 //n == 3, &n == mA1, pn == mA1, &pn == mA2, *pn == 3
	fmt.Println(n, &n, pn, &pn, *pn)

	n = *pn //n == 3, &n == mA1, pn == mA1, &pn == mA2, *pn == 3
	fmt.Println(n, &n, pn, &pn, *pn)
}