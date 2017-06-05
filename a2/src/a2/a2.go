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

	// fmt.Println(len(test))	
	
	return true
}
func getTokenContent(token IJsonToken)string{
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
				buffer +=","
			}
			buffer +="\n"
		}
		buffer =buffer +string(containerSigns[1])+""
	}
	
	switch token.(type) {
		case Pair:
			pair,_ := token.(Pair)
			buffer = buffer +"\"" +pair.Key.stringContent +"\":" +getTokenContent(pair.Val)
		case String:
			str,_ := token.(String)
			buffer = buffer +"\"" +str.stringContent +"\""
		case Unknown:
			str,_ := token.(Unknown)
			buffer = buffer +" " +str.Content +" "
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
			parseUnknown(buffer, tokenPool, currentToken)
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
				container.AddChild(pair)
				pairAwaiting = false
			}
			i = i + count
		}else if char ==":"{
			pairAwaiting=true
			i++
		}else if char ==","{
			parseUnknown( buffer,tokenPool, currentToken)
			
			i++
		}else if char == " "||char == "	"||char=="\n"{
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
			buffer = buffer + char
			i++
		}
	}
	return result,i
}

func parseUnknown(buffer string, tokenPool []IJsonToken, currentToken IJsonToken){
	content := strings.TrimSpace(buffer)
	if content==""{
		return
	}
	unkToken := Unknown{Content : content}
	tokenPool = append(tokenPool, unkToken)
	if currentToken.GetTypeString()=="Object"{
	key, _ := tokenPool[len(tokenPool)-2].(String)
	pair := Pair{Key : key, Val : tokenPool[len(tokenPool)-1] }
	tokenPool = tokenPool[:len(tokenPool)-2]
	container,_ :=currentToken.(*Object)
	container.AddChild(pair)

	}else if currentToken.GetTypeString()=="Array" {
		container,_ :=currentToken.(*Array)
		container.AddChild(unkToken)
	}
	buffer=""
	return
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
	fmt.Println("gagadsd")
	arr.Elements = append(arr.Elements, child)
}
func (arr *Array)GetTypeString() string {
	return "Array"
}
func (arr *Array)GetChilds() []IJsonToken{
	return arr.Elements
}

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

type Unknown struct{
	Content string
}
func (unk Unknown)GetTypeString() string {
	return "Unknown"
}
