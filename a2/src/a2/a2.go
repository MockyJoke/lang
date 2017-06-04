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
	jsonString := string(content)
	parseToken(jsonString)
	//testJson,_ := parseToken(jsonString)
	//if json==testJson {
		//fmt.Print(string(content))
	//}
	//fmt.Print(string(content[1:]))
	
	return true
}

func parseToken(text string) (IJsonToken,int){
	currentToken := createNewToken(string(text[0]))
	var i int = 1
	var buffer string = ""
	for i < len(text) {
		char := string(text[i])
		if char == "}" ||char == "]"{
			fmt.Println("Token:")
			fmt.Println(buffer)
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")
			i++
			break
		}else if char == "{" ||char == "["{
			innerToken,count := parseToken(text[i:])
			innerToken.GetTypeString()
			currentToken.AddChild(innerToken)
			buffer=buffer +"|"+innerToken.GetTypeString()+"|"
			//currentToken.SubTokens = append(currentToken.SubTokens, innerToken)
			i = i + count
		}else{
			buffer = buffer+char
		}
		i++
	}
	
	return currentToken,i
}
func createNewToken(char string) IJsonToken{
	var token IJsonToken
	if char == "{"{
		token = Object{}
	}else if char == "["{
		token = Array{}
	}else{
		//token = Unknown{}
	}
	return token
}

type IJsonToken interface {
	AddChild(child IJsonToken)
	GetTypeString() string
}

type Object struct{
   	Members []Member
}
func (obj Object)AddChild(child IJsonToken){
	member,ok := child.(Member)
	if !ok{
		fmt.Println("Failed to add a Member.")
		fmt.Println("Type is ."+child.GetTypeString())
	}
	obj.Members = append(obj.Members, member)
}
func (obj Object)GetTypeString() string {
	return "Object"
}

type Array struct{
	Elements []Element
}
func (arr Array)AddChild(child IJsonToken){
	element,ok := child.(Element)
	if !ok{
		fmt.Println("Failed to add a Element.")
	}
	arr.Elements = append(arr.Elements, element)
}
func (arr Array)GetTypeString() string {
	return "Array"
}

type Element struct{
	Values []Value
}
func (ele Element)AddChild(child IJsonToken){
	val,ok := child.(Value)
	if !ok{
		fmt.Println("Failed to add a Value.")
	}
	ele.Values = append(ele.Values, val)
}
func (ele Element)GetTypeString() string {
	return "Element"
}

type Member struct{
	Pairs []Pair
}
func (member Member)AddChild(child IJsonToken){
	pair,ok := child.(Pair)
	if !ok{
		fmt.Println("Failed to add a Pair.")
	}
	member.Pairs = append(member.Pairs, pair)
}
func (member Member)GetTypeString() string {
	return "Member"
}

type Pair struct{
	// key string
	// val Value
}
func (pair Pair)AddChild(child IJsonToken){
	return
}
func (pair Pair)GetTypeString() string {
	return "Pair"
}

type Value struct{
	// Tokens IJsonToken
}
func (val Value)AddChild(child IJsonToken){
	return
}
func (val Value)GetTypeString() string {
	return "Value"
}

type Unknown struct{
}
func (unk Unknown)AddChild(child IJsonToken){
	return
}
func (unk Unknown)GetTypeString() string {
	return "Unknown"
}