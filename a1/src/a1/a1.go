package a1

import (
	"math"
	"strings"
	"fmt"
	"io/ioutil"
	"errors"
	"strconv"
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

func linearSearch(x interface{}, lst interface{}) int{
	switch tpe := x.(type) {
		case int:
			slice, ok := lst.([]int)
			if !ok{
				panic("lst is not []int type.")
			}
			for i, element := range slice {
				if element == x.(int){
					return i
				}
			}	
		case string:
			slice, ok := lst.([]string)
			if !ok{
				panic("lst is not []string type.")
			}
			for i, element := range slice {
				if element == x.(string){
					return i
				}
			}
		default:
			var str = fmt.Sprintf("Unsupported type: %T\n", tpe)
			panic(str)
	}
	return -1
}

func allBitSeqs(n int) [][]int{
	if n<=0 {
		return [][]int{}
	}
	n = int(math.Pow(2,float64(n)))-1
	var result [][]int = [][]int{}
	var binLength = len(strconv.FormatInt(int64(n), 2))

	for i := 0; i <= n; i++ {
		var binSlice []int = []int{}
		
		var binStr = strconv.FormatInt(int64(i), 2)
		if(len(binStr)<binLength){
			var fmtStr = "%0"+strconv.Itoa(binLength-len(binStr))+"d"
			var zeroPads = fmt.Sprintf(fmtStr, 0)
			binStr = zeroPads + binStr
		}
		for _,char := range(binStr){
			digit, _ := strconv.Atoi(string(char))
			binSlice = append(binSlice,digit)
		}
		//fmt.Println(binSlice)
		result = append(result,binSlice)
	}
	//fmt.Println(result)
	return result
}