package a1

import "testing"
import "fmt"

func TestCountPrimes(t *testing.T) {
	// test stuff here...
	var error bool = false
	if countPrimes(0)!=0{
		fmt.Println("Failed at testcase n = 0.")
		error = true
	}
	if countPrimes(1)!=0{
		fmt.Println("Failed at testcase n = 1.")
		error = true
	}
	if countPrimes(2)!=1{
		fmt.Println("Failed at testcase n = 2.")
		error = true
	}
	if countPrimes(5)!=3{
		fmt.Println("Failed at testcase n = 5.")
		error = true
	}
	if countPrimes(10000)!=1229{
		fmt.Println("Failed at testcase n = 10000.")
		error = true
	}
	if error{
		t.Error("Failed")
	}

}

func TestIsPrime(t *testing.T) {
	// test stuff here...
	var error bool = false
	if isPrime(0)!=false{
		fmt.Println("0 is not a prime")
		error = true
	}
	if isPrime(1)!=false{
		fmt.Println("1 is not a prime")
		error = true
	}
	if isPrime(2)!=true{
		fmt.Println("2 is a prime")
		error = true
	}
	if isPrime(3)!=true{
		fmt.Println("3 is a prime")
		error = true
	}
	if isPrime(4)!=false{
		fmt.Println("4 is not a prime")
		error = true
	}
	if error{
		t.Error("Failed")
	}
}