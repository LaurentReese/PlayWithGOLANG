package main

import ("fmt"
		"os"
		"reflect")

func main() {
	funcWithManyInts()	
	funcWithManyInts(3, 4, 5, 6, 7, 8, 2e8)
	funcWithManyStrings()
	funcWithManyStrings("a", "friend", "in", "need", "is", "a", "friend", "indeed")
	funcWithListOfAnything(3, -7, 12.34, -23.87, 2e4, "Laurent", true) // 2e4 est vu comme une string
	funcWithManyStrings(os.Args...)	// transforme mon slice de strings, à savoir os.Args, en vararg de strings
}

func funcWithManyInts(intSlice ...int) {
	for x, oneInt := range intSlice {
		fmt.Println("int[",x,"]=",oneInt)
	}
}

func funcWithManyStrings(stringSlice ...string) {
	for x, oneString := range stringSlice {
		fmt.Println("string[",x,"]=",oneString)
	}
}

// ok, maintenant un vararg de n'importe quoi
type anything interface {}

func funcWithListOfAnything(anyList ...anything) {
	for _, anyItem := range anyList {
		fmt.Println("type=",reflect.TypeOf(anyItem))
	}
}

