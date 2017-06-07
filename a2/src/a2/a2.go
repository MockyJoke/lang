package main

import (
    "strings"
    "fmt"
    "io/ioutil"
    "os"
    "html"
)

func main() {
    if len(os.Args) == 1 {
        fmt.Println("No proper input json supplied.")
        return
    }
    fileName := os.Args[1]
    if fileName == "" {
        // no json file supplied
        fmt.Println("No proper input json supplied.")
    }
    rootJsonToken := parseJson(fileName)
    html := getTokenHtml(rootJsonToken, 0)
    fmt.Println(getFinalHtml("template.html", html))
}

// Parse json file to a tree structure
func parseJson(filename string) IJsonToken {
    filename = strings.TrimSpace(filename)
    content, err := ioutil.ReadFile(filename) // read the file
    if err != nil {
        fmt.Println(err)
        return &Object {}
    }
    jsonString := string(content)
    if len(jsonString) == 0 {
        fmt.Println("Error. Supplied file is empty.")
        return &Object {}
    }
    token, _ := parseToken(jsonString)
    return token
}

// Print json structre into html
func getTokenHtml(token IJsonToken, depth int) string {
    var indentChar = " "
        // ident used for members
    var indent string = ""
    i := 0
    for i <= depth {
            indent += indentChar
            i++
        }
        // ident used for closing brackets
    var endingIndent = ""
    i = 0
    for i < depth {
        endingIndent += indentChar
        i++
    }

    buffer := ""
    container, ok := token.(IContainer)
    if ok {
        // Current token is a container and recursively print the childs first.
        childs := container.GetChild()
        containerSigns := "{}"
        containerColor := "red"
        if (token.GetTypeString() == "Array") {
            containerSigns = "[]"
            containerColor = "blue"
        }
        buffer = buffer + getHtmlTag(string(containerSigns[0]), containerColor) + "\n"
        for i, ele := range childs {
            buffer = buffer + getTokenHtml(ele, depth + 1)
            if i < len(childs) - 1 {
                buffer += getHtmlTag(",", "orange")
            }
            buffer += "\n"
        }
        buffer = buffer + endingIndent + getHtmlTag(string(containerSigns[1]), containerColor) + ""
        return buffer
    }

    // Current token has no child
    switch token.(type) {
        case Pair:
            pair, _ := token.(Pair)
            buffer = buffer + indent + "\"" + getHtmlTag(getStringHtml(html.EscapeString(pair.Key.stringContent)), "SteelBlue") + "\" " + getHtmlTag(":", "DarkGreen") + " " + strings.TrimLeft(getTokenHtml(pair.Val, depth + 1), indentChar)
        case String:
            str, _ := token.(String)
            buffer = buffer + indent + "\"" + getHtmlTag(getStringHtml(html.EscapeString(str.stringContent)), "IndianRed") + "\""
        case Unknown:
            str, _ := token.(Unknown)
            unknownColor := "DarkOrchid"
            if str.Content == "true" || str.Content == "false" || str.Content == "null" {
                unknownColor = "GoldenRod"
            }
            buffer = buffer + indent + "" + getHtmlTag(str.Content, unknownColor) + ""
    }
    return buffer
}

// Helper for converting string into html, handling escape characters & html encodings
func getStringHtml(str string) string {
    buffer := ""
    i := 0
    for i < len(str) {
        char := string(str[i])
        if char == "\\" {
            if string(str[i + 1]) == "u" {
                buffer = buffer + getHtmlTag(string(str[i: i + 6]), "LightSeaGreen")
                i = i + 6
            } else {
                buffer = buffer + getHtmlTag(string(str[i: i + 2]), "MediumBlue")
                i = i + 2
            }
        } else {
            buffer = buffer + char
            i = i + 1
        }
    }
    return buffer
}

// Warp a string with span tag with color
func getHtmlTag(content string, color string) string {
    start := fmt.Sprintf("<span style=\"color:%v\">", color)
    mid := content
    end := fmt.Sprintf("</span>")
    return start + mid + end
}

// Embed formatteed json html in a actual html file
func getFinalHtml(templateFilename string, codeHtml string) string {
    templateFilename = strings.TrimSpace(templateFilename)
    content, err := ioutil.ReadFile(templateFilename) // read the file
    if err != nil {
        fmt.Println(err)
        return "Cannot open html template"
    }
    return fmt.Sprintf(string(content), codeHtml)
}

// Print json structure into formatted json text
func getTokenContent(token IJsonToken, depth int) string {
    var indentChar = " "
    var indent string = ""
    var endingIndent = ""
    i := 0
    for i <= depth {
        indent += indentChar
        i++
    }
    i = 0
    for i < depth {
        endingIndent += indentChar
        i++
    }

    var buffer string = ""
    container, ok := token.(IContainer)
    if ok {
        childs := container.GetChild()
        containerSigns := "{}"
        if (token.GetTypeString() == "Array") {
            containerSigns = "[]"
        }
        buffer = buffer + string(containerSigns[0]) + "\n"
        for i, ele := range childs {
            buffer = buffer + getTokenContent(ele, depth + 1)
            if i < len(childs) - 1 {
                buffer += ","
            }
            buffer += "\n"
        }
        buffer = buffer + endingIndent + string(containerSigns[1]) + ""
        return buffer
    }

    switch token.(type) {
        case Pair:
            pair, _ := token.(Pair)
            buffer = buffer + indent + "\"" + pair.Key.stringContent + "\":" + strings.TrimLeft(getTokenContent(pair.Val, depth + 1), indentChar)
        case String:
            str, _ := token.(String)
            buffer = buffer + indent + "\"" + str.stringContent + "\""
        case Unknown:
            str, _ := token.(Unknown)
            buffer = buffer + indent + " " + str.Content + " "
    }
    return buffer
}

// Method for parsing json text into structured json tokens
func parseToken(text string)(IJsonToken, int) {
    currentToken := createNewToken(string(text[0]))
    if currentToken.GetTypeString() == "String" {
        return parseString(text)
    }
    var tokenPool[] IJsonToken
    var pairAwaiting bool = false
    var i int = 1
    var buffer string = ""
    for i < len(text) {
        char := string(text[i])
        if char == "}" || char == "]" {
			// The end of a container, wrap anything left in buffer to a unknow token
            parseUnknown(buffer, & tokenPool, currentToken)
            buffer = ""
            pairAwaiting = false
            i++
            break

        } else if char == "{" || char == "[" || char == "\"" {
			// Start of a new container, recursively parse the inner-container
            innerToken, count := parseToken(text[i: ])
            tokenPool = append(tokenPool, innerToken)
            if currentToken.GetTypeString() == "Array" {
                container, _ := currentToken.(IContainer)
                container.AddChild(innerToken)
            } else if pairAwaiting {
                key, _ := tokenPool[len(tokenPool) - 2].(String)
                pair := Pair { Key: key,  Val: tokenPool[len(tokenPool) - 1] }
                tokenPool = tokenPool[: len(tokenPool) - 2]
                container, _ := currentToken.( * Object)
                container.AddChild(pair)
                pairAwaiting = false
            }
            i = i + count
        } else if char == ":" {
            pairAwaiting = true
            i++
        } else if char == "," {
            parseUnknown(buffer, & tokenPool, currentToken)
            buffer = ""
            pairAwaiting = false
            i++
        } else if char == " " || char == "	" || char == "\n" {
            i++
        } else {
            buffer = buffer + char
            i++
        }
    }
    return currentToken, i
}

// Get type for the current token
func createNewToken(char string) IJsonToken {
    var token IJsonToken
    if char == "{" {
        token = & Object {}
    } else if char == "[" {
        token = & Array {}
    } else if char == "\"" {
        token = & String {}
    } else {
        //token = Unknown{}
    }
    return token
}

// Helper for parsing string type
func parseString(text string)(String, int) {
    result := String {}
    var i int = 1
    var buffer string = ""
    for i < len(text) {
        char := string(text[i])
        if char == "\\" {
            buffer = buffer + char + string(text[i + 1])
            i += 2
            continue
        } else if char == "\"" {
            result.stringContent = buffer
            i++
            break
        } else {
            buffer = buffer + char
            i++
        }
    }
    return result, i
}

func debugTokenPool(tokenPool[] IJsonToken) {
    fmt.Printf("TokenPool has %v items.\n", len(tokenPool))
    if len(tokenPool) >= 1 {
        fmt.Printf("Top is %v\n", tokenPool[len(tokenPool) - 1].GetTypeString())
    }
}

// Helper for parsing unknown type
func parseUnknown(buffer string, tokenPool * [] IJsonToken, currentToken IJsonToken) {
    content := strings.TrimSpace(buffer)
    if content == "" {
        return
    }
    unkToken := Unknown { Content: content } 
	*tokenPool = append(*tokenPool, unkToken)
    if currentToken.GetTypeString() == "Object" {
        key, _ := (*tokenPool)[len(*tokenPool) - 2].(String)
        pair := Pair { Key: key, Val: unkToken }
        container, _ := currentToken.( * Object)
        container.AddChild(pair)
    } else if currentToken.GetTypeString() == "Array" {
        container, _ := currentToken.( * Array)
        container.AddChild(unkToken)
    }
    return
}

// General type for Json token, could be anything
type IJsonToken interface {
    GetTypeString() string
    GetContent() string
}

// General type for Json Object & Array, which could have children
type IContainer interface {
    AddChild(child IJsonToken)
    GetChild()[] IJsonToken
}

// Struct for Json Object, start and end with {}
type Object struct {
    Members[] IJsonToken
}
func(obj * Object) AddChild(child IJsonToken) {
    obj.Members = append(obj.Members, child)
}
func(obj * Object) GetTypeString() string {
    return "Object"
}
func(obj * Object) GetContent() string {
    return fmt.Sprintf("Object with %v items, last:%v", len(obj.Members), obj.Members[len(obj.Members) - 1].GetContent())
}
func(obj * Object) GetChild()[] IJsonToken {
    return obj.Members
}

// Struct for Json Object, start and end with []
type Array struct {
    Elements[] IJsonToken
}
func(arr * Array) AddChild(child IJsonToken) {
    arr.Elements = append(arr.Elements, child)
}
func(arr * Array) GetTypeString() string {
    return "Array"
}
func(arr * Array) GetContent() string {
    return fmt.Sprintf("Array with %v items", len(arr.Elements))
}
func(arr * Array) GetChild()[] IJsonToken {
    return arr.Elements
}

// Struct for Json key-value pair, in the form of key:value
type Pair struct {
    Key String
    Val IJsonToken
}
func(pair Pair) AddChild(child IJsonToken) {
    return
}
func(pair Pair) GetTypeString() string {
    return "Pair"
}
func(pair Pair) GetContent() string {
    return fmt.Sprintf("Pair, key: %v value: %v", pair.Key.GetContent(), pair.Val.GetContent())
}

// Struct for Json string, start and end with ""
type String struct {
    stringContent string
}
func(str String) GetTypeString() string {
    return "String"
}
func(str String) GetContent() string {
    return fmt.Sprintf("String:%v", str.stringContent)
}

// Struct for other Json value, could be anything
type Unknown struct {
    Content string
}
func(unk Unknown) GetTypeString() string {
    return "Unknown:"
}
func(unk Unknown) GetContent() string {
    return fmt.Sprintf("Unknown:%v", unk.Content)
}