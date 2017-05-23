package a1

import (
	"math"
	"strings"
	"fmt"
	"io/ioutil"
)

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

func countStrings(filename string) map[string]int{
	filename = strings.TrimSpace(filename)
	content, err := ioutil.ReadFile(filename) // read the file
    if err != nil {
        fmt.Print(err)
		return nil
    }
	var words = strings.Fields(string(content))
	dict := make(map[string]int)
	for _, element := range words {
		if val, ok := dict[element]; ok {
			// word already exist, doing increment
			dict[element] = val+1
		}else{
			// word do not exist, init with one
			dict[element] = 1
		}
	}
	return dict
}