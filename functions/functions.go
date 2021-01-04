// https://golangbot.com/first-class-functions/

package main

import (
	"fmt"
)

type mafonc func(string, string)

func main() {
	var x int = 7
	var a func(string, string) = func(one string, two string) {
		fmt.Println(one, two)
	}
	a("Un", "Dauphin")
	fmt.Printf("%T %T\n", a, x)

	// En simplifiant les déclarations
	y := 8 // je pourrais faire aussi "var y int = 8"
	var b mafonc // je pourrais aussi la déclarer comme j'ai déclaré la fonction a
	// je pourrais aussi faire b := func(one string, two string) {.....}
	b = func(one string, two string) { // je pourrais faire aussi var b = func(one string, two string)
		fmt.Println(two, one)
	}
	b("Chiens", "Deux")
	fmt.Printf("%T %T %T\n", b, y, fmt.Println)

	animaux(a,b, "Un", "Dauphin", "Chiens", "Deux", fmt.Println)
}

// fmt.Println is of type func(...interface {}) (int, error)

func animaux(	f1 mafonc,
				f2 func(string, string),
				animal1 string,
				animal2 string,
				animal3 string,
				animal4 string,
				z func(...interface {}) (int, error)) {
	z()	// ça va me sauter une ligne
	f1(animal1, animal2)
	f2(animal3, animal4)
}

