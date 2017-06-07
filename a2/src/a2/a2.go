package main

import (
	"strings"
	"fmt"
	"io/ioutil"
	"os"
	"html"
)

func main(){
	fileName := os.Args[1]
	if fileName==""{
		fmt.Print("No proper input json supplied.")
	}
	rootJsonToken := parseJson(fileName)
	//json := getTokenContent(rootJsonToken,0)
	html := getTokenHtml(rootJsonToken, 0)
	//fmt.Println(json)
	fmt.Println(getFinalHtml("template.html",html))
	
}

func parseJson(filename string) IJsonToken {
	filename = strings.TrimSpace(filename)
	content, err := ioutil.ReadFile(filename) // read the file
	if err != nil {
        fmt.Println(err)
		return &Object{}
    }
	jsonString := string(content)
	if len(jsonString)==0{
		return &Object{}
	}
	token, _ := parseToken(jsonString)
	return token

}

func getTokenHtml(token IJsonToken, depth int)string{
	var indentChar =" "
	var indent string =""
	var endingIndent = ""
	i := 0
	for i <= depth{
		indent += indentChar
		i++
	}
	i = 0
	for i < depth{
		endingIndent += indentChar
		i++
	}
	var buffer string = ""
	container,ok := token.(IContainer)
	if ok{
		childs := container.GetChilds()
		containerSigns := "{}"
		containerColor := "red"
		if(token.GetTypeString()=="Array"){
			containerSigns="[]"
			containerColor = "blue"
		}
		buffer =buffer +getHtmlTag(string(containerSigns[0]),containerColor)+"\n"
		for i,ele := range childs{
			buffer =buffer+ getTokenHtml(ele,depth+1)
			if i<len(childs)-1{
				buffer +=getHtmlTag(",","orange")
			}
			buffer +="\n"
		}
		buffer =buffer +endingIndent+getHtmlTag(string(containerSigns[1]),containerColor)+""
		return buffer
	}
	
	switch token.(type) {
		case Pair:
			pair,_ := token.(Pair)
			buffer = buffer +indent+"\""+pair.Key.stringContent +"\" "+getHtmlTag(":","DarkGreen")+" " +strings.TrimLeft(getTokenHtml(pair.Val, depth+1),indentChar)
		case String:
			str,_ := token.(String)
			buffer = buffer+indent+"\"" +getHtmlTag(html.EscapeString(str.stringContent),"IndianRed") +"\""
		case Unknown:
			str,_ := token.(Unknown)
			unknownColor:= "DarkOrchid"
			if str.Content == "true" || str.Content =="false"|| str.Content=="null"{
				unknownColor = "GoldenRod"
			}
			
			buffer = buffer + indent +"" +getHtmlTag(str.Content, unknownColor) +""
	}
	return buffer
}

func getHtmlTag(content string,color string) string{
	//color := "red"
	start := fmt.Sprintf("<span style=\"color:%v\">",color)
	mid := content
	end := fmt.Sprintf("</span>")
	return start+mid+end
}
func getFinalHtml(templateFilename string, codeHtml string) string{
	templateFilename = strings.TrimSpace(templateFilename)
	content, err := ioutil.ReadFile(templateFilename) // read the file
	if err != nil {
        fmt.Println(err)
		return "Cannot open html template"
    }
	return fmt.Sprintf(string(content),codeHtml)
}

func getTokenContent(token IJsonToken,depth int)string{
	var indentChar =" "
	var indent string =""
	var endingIndent = ""
	i := 0
	for i <= depth{
		indent += indentChar
		i++
	}
	i = 0
	for i < depth{
		endingIndent += indentChar
		i++
	}
	
	var buffer string = ""
	container,ok := token.(IContainer)
	if ok{
		childs := container.GetChilds()
		containerSigns := "{}"
		if(token.GetTypeString()=="Array"){
			containerSigns="[]"
		}
		buffer =buffer +string(containerSigns[0])+"\n"
		for i,ele := range childs{
			buffer =buffer+ getTokenContent(ele,depth+1)
			if i<len(childs)-1{
				buffer +=","
			}
			buffer +="\n"
		}
		buffer =buffer +endingIndent+string(containerSigns[1])+""
		return buffer
	}
	
	switch token.(type) {
		case Pair:
			pair,_ := token.(Pair)
			buffer = buffer +indent+"\""+pair.Key.stringContent +"\":" +strings.TrimLeft(getTokenContent(pair.Val, depth+1),indentChar)
		case String:
			str,_ := token.(String)
			buffer = buffer+indent+"\"" +str.stringContent +"\""
		case Unknown:
			str,_ := token.(Unknown)
			buffer = buffer + indent +" " +str.Content +" "
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
		if char == "}" ||char == "]"{
			parseUnknown(buffer, &tokenPool, currentToken)
			buffer=""
			pairAwaiting=false
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
			parseUnknown( buffer,&tokenPool, currentToken)
			buffer=""
			pairAwaiting=false
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
			buffer= buffer + char + string(text[i+1])
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
func debugTokenPool(tokenPool []IJsonToken){
	fmt.Printf("TokenPool has %v items.\n",len(tokenPool))
	if len(tokenPool)>=1{
		fmt.Printf("Top is %v\n",tokenPool[len(tokenPool)-1].GetTypeString())
	}
}
func parseUnknown(buffer string, tokenPool *[]IJsonToken, currentToken IJsonToken){
	content := strings.TrimSpace(buffer)
	if content==""{
		return
	}
	unkToken := Unknown{Content : content}
	*tokenPool = append(*tokenPool, unkToken)
	if currentToken.GetTypeString()=="Object"{
		//debugTokenPool(*tokenPool)
		key, _ := (*tokenPool)[len(*tokenPool)-2].(String)
		pair := Pair{Key : key, Val : unkToken }
		//*tokenPool = (*tokenPool)[:len(*tokenPool)-2]
		container,_ :=currentToken.(*Object)
		container.AddChild(pair)
		

	}else if currentToken.GetTypeString()=="Array" {
		container,_ :=currentToken.(*Array)
		container.AddChild(unkToken)
	}
	return
}


type IJsonToken interface {
	GetTypeString() string
	GetContent() string
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
func (obj *Object)GetContent() string {
	return fmt.Sprintf("Object with %v items, last:%v",len(obj.Members),obj.Members[len(obj.Members)-1].GetContent())
}
func (obj *Object)GetChilds() []IJsonToken{
	return obj.Members
}

type Array struct{
	Elements []IJsonToken
}
func (arr *Array)AddChild(child IJsonToken){
	arr.Elements = append(arr.Elements, child)
}
func (arr *Array)GetTypeString() string {
	return "Array"
}
func (arr *Array)GetContent() string {
	return fmt.Sprintf("Array with %v items",len(arr.Elements))
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
func (pair Pair)GetContent() string {
	return fmt.Sprintf("Pair, key: %v value: %v",pair.Key.GetContent(),pair.Val.GetContent())
}

type String struct{
	stringContent string
}
func (str String)GetTypeString() string {
	return "String"
}
func (str String) GetContent() string {
	return fmt.Sprintf("String:%v",str.stringContent)
}

type Unknown struct{
	Content string
}
func (unk Unknown)GetTypeString() string {
	return "Unknown:"
}
func (unk Unknown) GetContent() string {
	return fmt.Sprintf("Unknown:%v",unk.Content)
}

