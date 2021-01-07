
package main

import "fmt"
import "unsafe"
// #include <stdio.h>
// #include <stdlib.h>
// // see https://golang.org/cmd/cgo/
// int fibonacci(int x)
// {
//   // printf("%ld\n",sizeof(int));	
//   if (x==0)
//     return 0;
//   return (x + fibonacci(x - 1));
// }
// char *changeString(char *s)
// { // "poisson" devient "pCHIENn"
//	 s[1] = 'C';
//	 s[2] = 'H';
//	 s[3] = 'I';
//	 s[4] = 'E';
//	 s[5] = 'N';
//   return s;
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


	var s string = "poisson"
	fmt.Println("animal=", s)
	s2 := C.CString(s)
	defer C.free(unsafe.Pointer(s2)) // ! needs to include stdlib.h otherwise the compiler will say : could not determine kind of name for C.free
	// prototype C.GoString :  func C.GoString(*C.char) string
	s3 := C.GoString(C.changeString(s2))
	fmt.Println("animal=", s3)	
}

// N.B. voir https://dave.cheney.net/tag/cgo
// Il ne faut pas abuser du cgo, il vaut mieux, quand c'est possible, porter le C en GO