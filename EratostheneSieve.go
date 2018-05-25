//prime numbers:
//1, 2, 3, 5, 7
//all next prime numbers should end up with 1, 3, 7, 9
package main
import "fmt"

var (
	length int
	step int
)

func sieve(n *int) []bool {
	if *n < 2 {
		return nil
	}

	arr := make([]bool, *n + 1)
	for i:=2; i<len(arr); i++ {
		arr[i] = true
	}

	for i:=2; i<len(arr); i++ {
		if arr[i] {
			for j:= i+i; j<len(arr); j+=i {
				arr[j] = false
			}
		}
	}
	return arr
}

func distributionAnalysis(arr []bool) []int {
	counter := make([]int, len(arr)/length + 1)
	for i := range arr {
		if arr[i] {
			counter[i/length] += 1
		}
	}
	return counter
}

func distributionPrinter(arr []int) {
	for i := range arr {
		for j:=0; j<arr[i]; j+=step {
			fmt.Print(">")
		}
		fmt.Print("\n")
	}
}

func main() {
	n := new(int)
	fmt.Scanln(n)

	boolPrime := sieve(n)
	if boolPrime == nil {
		fmt.Println("No prime numbers")
		return
	}

	var primes []int
	for i := range boolPrime {
		if boolPrime[i] {
			primes = append(primes, i)
		}
	}

	//values for control printing result in distributionAnalysis
	switch {
	case *n < 10:
		length = 1
	default:
		length = *n/10 //the more value - the more lines will be printed
	}

	switch {
	case *n < 30:
		step = *n/1
	case *n < 100:
		step = *n/30
	case *n < 1000:
		step = *n/100
	case *n < 10000:
		step = *n/1000
	default:
		step = *n/10000
	} //control line size

	amountArr := distributionAnalysis(boolPrime)
	distributionPrinter(amountArr)
}
