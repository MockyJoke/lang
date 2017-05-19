package a1

import "math"
// import "fmt"
func isPrime(x int) bool{
	if x < 2{
		return false
	}
	limit := int(math.Sqrt(float64(x)))
	for i:=2; i<=limit; i++{
		if x % i==0 {
			return false
		}
	}
	return true
}
func countPrimes(n int) int{
	var count int = 0
	if n < 2{
		return count
	}
	for i:=2; i<=n; i++{
		if isPrime(i) {
			count++
		}
	}
	return count
}
// func main() {  
//     if isPrime(5){
// 		fmt.Print("Yes")
// 	}else{
// 		fmt.Print("No")
// 	}
// }
