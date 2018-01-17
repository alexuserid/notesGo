package pointers


//mA == memory adress
func main() {
	var n int = 6 //n == 6, & == mA1
	var pn *int //pn == <nil>, &pn == mA2

	pn = &n //n == 6, &n == mA1, *pn == 6, &pn == mA2
	n = 7 //n == 7, &n == mA1, pn == mA1, *pn == 7, &pn == mA2
	*pn = 3 //n == 3, &n == mA1, pn == mA1, &pn == mA2, *pn == 3
	n = *pn //n == 3, &n == mA1, pn == mA1, &pn == mA2, *pn == 3
}
