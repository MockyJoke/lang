package a1

import "testing"
import "fmt"
import "reflect"

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

func TestCountStrings(t *testing.T) {
	var dict map[string]int = countStrings("test.txt")
	var target map[string]int =map[string]int {"The":1, "the":1, "big":3, "dog":1, "ate":1, "apple":1}
	result := reflect.DeepEqual(dict, target)
	if !result {
		fmt.Println("Result mapping incorrect")
		fmt.Printf("Result: %v\n", dict)
		fmt.Printf("Target: %v\n", target)
		t.Error("Failed")
	}
}


func TestEqualsTime24(t *testing.T) {
	var error bool = false

	var t1 Time24 = Time24 { hour:2, minute:35, second:27 }
	var t2 Time24 = Time24 { hour:2, minute:35, second:27 }
	var t3 Time24 = Time24 { hour:2, minute:35, second:28 }
	if equalsTime24(t1,t2)!=true{
		error = true
	}
	if equalsTime24(t1,t3)!=false{
		error = true
	}
	if error{
		t.Error("Failed")
	}
}

func TestLessThanTime24(t *testing.T) {
	var error bool = false

	var t1 Time24 = Time24 { hour:2, minute:35, second:27 }
	var t2 Time24 = Time24 { hour:2, minute:35, second:27 }
	var t3 Time24 = Time24 { hour:2, minute:35, second:28 }
	if lessThanTime24(t1,t3)!=true{
		error = true
	}
	if lessThanTime24(t1,t2)!=false{
		error = true
	}
	if error{
		t.Error("Failed")
	}
}

func TestTime24ToString(t *testing.T){
	var error bool = false
	
	var t1 Time24 = Time24 { hour:2, minute:35, second:27 }
	fmt.Println(t1)
	if t1.String()!="02:35:27"{
		error = true
	}
	if error{
		t.Error("Failed")
	}
}

func TestValidTime24(t *testing.T){
	var error bool = false
	
	var t1 Time24 = Time24 { hour:2, minute:35, second:27 }
	var t2 Time24 = Time24 { hour:2, minute:35, second:80 }
	var t3 Time24 = Time24 { hour:30, minute:35, second:0 }
	
	if t1.validTime24()!=true{
		fmt.Printf("Test case failed on %v",t1)
		error = true
	}
	if t2.validTime24()!=false{
		fmt.Printf("Test case failed on %v",t2)
		error = true
	}
	if t3.validTime24()!=false{
		fmt.Printf("Test case failed on %v",t3)
		error = true
	}
	if error{
		t.Error("Failed")
	}
}

func TestMinTime24(t *testing.T){
	var error bool = false
	var t1 Time24 = Time24 { hour:2, minute:35, second:27 }
	var t2 Time24 = Time24 { hour:12, minute:45, second:4 }
	var t3 Time24 = Time24 { hour:23, minute:35, second:0 }

	if _,err := minTime24([]Time24{}) ; err==nil{
		fmt.Println("Test case failed on empty array.")
		error = true
	}
	if minTime,err := minTime24([]Time24{t2,t1,t3}) ; err!=nil || !equalsTime24(minTime,t1) {
		fmt.Println("Test case failed when comparing t2,t1,t3.")
		error = true
	}
	if error{
		t.Error("Failed")
	}
}

