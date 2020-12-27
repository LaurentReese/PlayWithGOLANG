package main
import (
    "fmt"
	"os"
	"time"
)
func main() {
	var args []string = os.Args
	fmt.Println("nombre d'arguments =", len(args))
	fmt.Print("Arguments = ")
	for _, arg := range args {
		/*time.Sleep(time.Second) or */ time.Sleep(1000 * time.Millisecond)
		fmt.Print(arg," ")
	}
	fmt.Println()
}