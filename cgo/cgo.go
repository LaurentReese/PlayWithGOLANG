
package main

import "fmt"
import "unsafe"

// #include <stdio.h>
// // see https://golang.org/cmd/cgo/
// int fibonacci(int x)
// {
//   // printf("%ld\n",sizeof(int));	
//   if (x==0)
//     return 0;
//   return (x + fibonacci(x - 1));
// }
import "C"

var m C.int

func main() {
	var t int = int(C.fibonacci(10)) // Attention il y a des précautions à prendre car le type int en GO fait 8 octets alors qu'il n'en fait que 4 en C !
	fmt.Println(t, unsafe.Sizeof(t))
	var z *int
	z = &t
	fmt.Printf("%T\n", z)
	*z = 12
	t = int(C.fibonacci(C.int(*z))) // Attention il y a des précautions à prendre car le type int en GO fait 8 octets alors qu'il n'en fait que 4 en C !	
	fmt.Println(t, unsafe.Sizeof(t), unsafe.Sizeof(m))
}

// N.B. voir https://dave.cheney.net/tag/cgo
// Il ne faut pas abuser du cgo, il vaut mieux, quand c'est possible, porter le C en GO