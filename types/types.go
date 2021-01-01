package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// Lancer avec les arguments : sdfjk 12 -5 1234 -576 23.45 -12.17 sfsdf aa 0x1c8 true false

func main() {
	var args []string = os.Args
	fmt.Println("nombre d'arguments =", len(args))
	fmt.Println("Types des arguments : ")

	for _, arg := range args {
		fmt.Print(arg, " est de type ")

		b, errBool := strconv.ParseBool(arg)
		u, errUint := strconv.ParseUint(arg, 0, 64) // 0 en 2ème paramètre : permet de parser de l'hexa, genre "0x1c8"
		i, errInt := strconv.ParseInt(arg, 0, 64)
		// i, errInt := strconv.ParseInt(arg, 10, 64) // on est en base 10, mais ça pourrait être de l'hexadécimal
		// i, errInt := strconv.ParseInt(arg, 16, 64) // on est en base 16 : genre deadbeef
		f, errFloat := strconv.ParseFloat(arg, 64) // ParseUfloat n'existe pas

		if errBool == nil { // c'est un booléen
			fmt.Println(reflect.TypeOf(b).String())
		} else {
			if errUint == nil { // c'est un entier non signé
				fmt.Println(reflect.TypeOf(u).String())
			} else {
				if errInt == nil { // c'est un entier signé
					fmt.Println(reflect.TypeOf(i).String())
				} else {
					if errFloat == nil { // c'est un float
						fmt.Println(reflect.TypeOf(f).String())
					} else {
						fmt.Println(reflect.TypeOf(arg).String())
					}
				}
			}
		}
	}
}
// Je pourrais affiner quand c'est un entier pour déterminer si c'est un uint8 ou uint16 ou uint32 ou uint64 ou int8 ou int16 ou int32 ou int64

