package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// 	func Fields(s string) []string

func WordCount(s string) map[string]int {
	var theMap map[string]int = make(map[string]int)
	var theFields []string = strings.Fields(s)
	for i:= range theFields { theMap[theFields[i]] ++ }
	return theMap
	//return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
