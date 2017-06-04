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
	token, _ := parseToken(jsonString)
	fmt.Print(getTokenContent(token))
	// var test []Object
	// test = append(test, Object{})
	// test = append(test, Object{})
	// test = append(test, Object{})
	// test = append(test, Object{})
	// fmt.Println(len(test))	

	//testJson,_ := parseToken(jsonString)
	//if json==testJson {
		//fmt.Print(string(content))
	//}
	//fmt.Print(string(content[1:]))
	
	return true
}
func getTokenContent(token IJsonToken)string{
	//fmt.Println(token.GetTypeString())
	var buffer string = ""
	container,ok := token.(IContainer)
	if ok{
		childs := container.GetChilds()
		//fmt.Printf("Found container, got %v childs.\n",len(childs))
		containerSigns := "{}"
		if(token.GetTypeString()=="Array"){
			containerSigns="[]"
		}
		buffer =buffer +string(containerSigns[0])+"\n"
		for i,ele := range childs{
			buffer =buffer + getTokenContent(ele)
			if i<len(childs)-1{
				buffer +=",\n"
			}
		}
		buffer =buffer +string(containerSigns[1])+"\n"
	}
	
	switch token.(type) {
		case Pair:
			pair,_ := token.(Pair)
			buffer = buffer +"\"" +pair.Key.stringContent +"\":" +getTokenContent(pair.Val)
		case String:
			str,_ := token.(String)
			buffer = buffer +"\"" +str.stringContent +"\""
	}
	return buffer
}

func parseToken(text string) (IJsonToken,int){
	currentToken := createNewToken(string(text[0]))
	if currentToken.GetTypeString()=="String"{
		return parseString(text)
	}
	var tokenPool []IJsonToken
	var pairAwaiting bool = false 
	var i int = 1
	var buffer string = ""
	for i < len(text) {
		char := string(text[i])
		
		//container,_ :=currentToken.(*Object)
		//fmt.Printf("At:(%v)current %v childs.\n",char,len(container.Members))
		if char == "}" ||char == "]"{
			fmt.Println("Token:")
			fmt.Println(buffer)
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")
			i++
			break
		
		}else if char == "{" ||char == "[" ||char=="\"" {
			innerToken,count := parseToken(text[i:])
			tokenPool = append(tokenPool, innerToken)
			if currentToken.GetTypeString()=="Array"{
				container,_ :=currentToken.(IContainer)
				container.AddChild(innerToken)
			}
			if pairAwaiting {
				key, _ := tokenPool[len(tokenPool)-2].(String)
				pair := Pair{Key : key, Val : tokenPool[len(tokenPool)-1] }
				tokenPool = tokenPool[:len(tokenPool)-2]
				container,_ :=currentToken.(*Object)
				//fmt.Printf("current %v childs.\n",len(container.Members))
				container.AddChild(pair)
				//fmt.Println("Added pair "+key.stringContent+ " to "+currentToken.GetTypeString()+".")
				//container,_ =currentToken.(Object)
				
				//fmt.Printf("current %v childs.\n",len(container.Members))
				
				pairAwaiting = false
			}else{
			}
			buffer=buffer +"|"+innerToken.GetTypeString()+"|"
			
			i = i + count
		}else if char ==":"{
			pairAwaiting=true
			i++
		}else{
			buffer = buffer+char
			i++
		}
	}
	return currentToken,i
}

func createNewToken(char string) IJsonToken{
	var token IJsonToken
	if char == "{"{
		token = &Object{}
	}else if char == "["{
		fmt.Println("yoyo")
		token = &Array{}
	}else if char =="\""{
		token = &String{}
	}else{
		//token = Unknown{}
	}
	return token
}
func parseString(text string) (String,int){
	result := String{}
	var i int = 1
	var buffer string = ""
	for i < len(text) {
		char := string(text[i])
		if char == "\\"{
			i+=2
			continue
		}else if char == "\""{
			result.stringContent = buffer
			i++
			break
		}else{
			buffer = buffer +char
			i++
		}
	}
	return result,i
}

type IJsonToken interface {
	GetTypeString() string
}
type IContainer interface{
	AddChild(child IJsonToken)
	GetChilds() []IJsonToken
}

type Object struct{
   	Members []IJsonToken
}
func (obj *Object)AddChild(child IJsonToken){
	// member,ok := child.(Member)
	// if !ok{
	// 	fmt.Println("Failed to add a Member.")
	// 	fmt.Println("Type is ."+child.GetTypeString())
	// }
	//fmt.Println("Adding "+child.GetTypeString()+" to Object.")
	obj.Members = append(obj.Members, child)
	//fmt.Printf("Currently %v childs\n",len(obj.Members))
	
}
func (obj *Object)GetTypeString() string {
	return "Object"
}
func (obj *Object)GetChilds() []IJsonToken{
	return obj.Members
}

type Array struct{
	Elements []IJsonToken
}
func (arr *Array)AddChild(child IJsonToken){
	// element,ok := child.(Element)
	// if !ok{
	// 	fmt.Println("Failed to add a Element.")
	// }
	//fmt.Println("Adding "+child.GetTypeString()+" to Array.")
	fmt.Println("gagadsd")
	arr.Elements = append(arr.Elements, child)
	//fmt.Printf("Currently %v childs\n",len(arr.Elements))
}
func (arr *Array)GetTypeString() string {
	return "Array"
}
func (arr *Array)GetChilds() []IJsonToken{
	return arr.Elements
}
// type Element struct{
// 	Values []IJsonToken
// }
// func (ele Element)AddChild(child IJsonToken){
// 	val,ok := child.(Value)
// 	if !ok{
// 		fmt.Println("Failed to add a Value.")
// 	}
// 	ele.Values = append(ele.Values, val)
// }
// func (ele Element)GetTypeString() string {
// 	return "Element"
// }

// type Member struct{
// 	Pairs []Pair
// }
// func (member Member)AddChild(child IJsonToken){
// 	pair,ok := child.(Pair)
// 	if !ok{
// 		fmt.Println("Failed to add a Pair.")
// 	}
// 	member.Pairs = append(member.Pairs, pair)
// }
// func (member Member)GetTypeString() string {
// 	return "Member"
// }

type Pair struct{
	Key String
	Val IJsonToken
}
func (pair Pair)AddChild(child IJsonToken){
	return
}
func (pair Pair)GetTypeString() string {
	return "Pair"
}


type String struct{
	stringContent string
}
func (str String)GetTypeString() string {
	return "String"
}

// type Value struct{
// 	// Tokens IJsonToken
// }
// func (val Value)AddChild(child IJsonToken){
// 	return
// }
// func (val Value)GetTypeString() string {
// 	return "Value"
// }
