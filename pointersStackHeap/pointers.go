// Donc ici je m'aperçois que je peux retourner une adresse de la pile
// En fait il n'y a pas vraiment de pile et de tas dans GO (en tout cas pas dans la doc)
// même si forcément il les utilise.
// Je peux réutiliser sans problème une adresse allouée localement (voir ma fonction a())
// Ca marche même si je ne retourne pas l'adresse, car dans le main je la saisis avec la valeur affichée dans la fonction a()
// En lisant les docs, j'ai compris que si on retourne une adresse locale, le compilo comprend qu'on va la réutiliser et donc il ne la libère pas.
// Dans ce programme j'ai même testé que sans retouner l'adresse je peux encore l'utiliser... pendant un certain temps, jusqu'à ce que le garbage collector fasse son travail...


package main

import "fmt"
import "reflect"
import "unsafe"
import "runtime"

func a() /*(*int)*/ {
	i := 4
	fmt.Printf("adresse de i = %d\n", unsafe.Pointer(&i))
	fmt.Printf("taille de l'adresse de i = %d\n", reflect.TypeOf(&i).Size())
	/*return &i*/
	}

func main() {
	/*j := */a()
	runtime.GC()
	var j *int
	fmt.Println("=========================")	
	//fmt.Printf("contenu de j = %v\n", *j)
	
	//var z int = 8
	address := uintptr(unsafe.Pointer(j))

	fmt.Scanf("%d", &address) 


	fmt.Printf("address=%d\n", address)
	p := uintptr(unsafe.Pointer(address))
	q := unsafe.Pointer(p)
    fmt.Println(*(*int)(q))

}