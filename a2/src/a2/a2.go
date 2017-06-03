package a2

import (
	// "math"
	"strings"
	"fmt"
	"io/ioutil"
	// "errors"
	// "strconv"
)

func parseJson(filename string) bool {
	filename = strings.TrimSpace(filename)
	content, err := ioutil.ReadFile(filename) // read the file
	if err != nil {
        fmt.Print(err)
		return false
    }
	fmt.Print(string(content))
	return true
}

type Time24 struct {
   hour, minute, second uint8
}
