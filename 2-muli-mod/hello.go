package main

import "fmt"
import "github.com/buger/jsonparser"

func main() {
	fmt.Println("main")
	_, _ = jsonparser.ParseString([]byte("{}"))
}
