package a1

import (
	"math"
	"strings"
	"fmt"
	"io/ioutil"
	"errors"
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

type Time24 struct {
    hour, minute, second uint8
}

func equalsTime24(t Time24,otherT Time24) bool {
    if t.hour==otherT.hour && t.minute==otherT.minute && t.second==otherT.second{
		return true
	}
	return false
}

func lessThanTime24(t Time24,otherT Time24) bool {
	var t1 uint32 = t.getTick()
	var t2 uint32 = otherT.getTick()
	if t1 < t2{
		return true
	}
    return false
}

func (t Time24) getTick() uint32 {
	return uint32(t.hour)*3600 +uint32(t.minute)*60 +uint32(t.second)
}

func (t Time24) String() string {
	var str = fmt.Sprintf("%02d:%02d:%02d", t.hour, t.minute, t.second)
	return str;
}

func (t Time24) validTime24() bool{
	if t.hour >= 0 && t.hour <24 &&
	t.minute >= 0 && t.minute <60 &&
	t.second >= 0 && t.second <60 {
		return true
	}
	return false
}

func minTime24(times []Time24) (Time24, error){
	if len(times)==0 {
		error := errors.New("Encountered empty Time24 array")
		return Time24 { hour:0, minute:0, second:0 },error
	}

	var minIndex int = 0
	var minTick uint32 = times[0].getTick()
	for i, element := range times {
		if element.getTick() < minTick{
			minIndex = i
			minTick=element.getTick()
		}
	}
	return times[minIndex],nil
}