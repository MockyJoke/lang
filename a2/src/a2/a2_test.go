package main

import "testing"
import "fmt"


func TestParseJson(t *testing.T) {
	rootJsonToken := parseJson("test.json")
	html := getTokenHtml(rootJsonToken, 0)
	fmt.Println(getFinalHtml("template.html",html))
}