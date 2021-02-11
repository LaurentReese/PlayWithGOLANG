package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	os.Args = append(os.Args, "un", "deux", "trois") // extend it with my "own" args
	var args []string = os.Args
//	args = append(args, "un", "deux", "trois") // extend it with my "own" args
//	os.Args = args
	fmt.Println("nombre d'arguments =", len(args))
	fmt.Print("Arguments = ")
	for _, arg := range args {
		/*time.Sleep(time.Second) or */ time.Sleep(1000 * time.Millisecond)
		fmt.Print(arg, " ")
	}
	fmt.Println()
}
