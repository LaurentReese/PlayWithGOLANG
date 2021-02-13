// https://golang.org/pkg/os/

// Re-write the famous "cat" program, except that there is no stdin, we just want filenames as arguments

// Just to have in mind the functions : os.Open(), File.Stat(), File.Read(), Stat.Size()

package main

import (   "os"
 			"fmt"
			 "errors")

func check(err error) {
	if err!=nil { panic(err) }
}

func cat(filename string) {
	file, err := os.Open(filename) // For read access.
	check(err) // TO DO : if error ignore just this file ?
	stat, err := file.Stat()
	check(err)  // TO DO : if error ignore just this file ?
	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	check(err)
	fmt.Printf("%s",data)
}

func main() {
	if len(os.Args) == 1 { check(errors.New("No file specified")) }
	for index, filename := range os.Args {
		if index==0 {continue} // skip the executable
		cat(filename)
	}
}			 



