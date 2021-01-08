package main
// pour moi : Site en français dédié à Go
// http://decouvric.cluster013.ovh.net/

import "fmt"

func essai() func(int) int {
	fmt.Println("coucou")
	i := 1
	return func(param int) int {
		i += param + 1
		return i
	}
}

func main(){
	maFonction := essai() // écrit "coucou" une fois car essai() est exécutée, i est instancié, on print "coucou"
	res := maFonction(3) // ici c'est la fonction anonyme f2 (voir commentaires en bas) qui est exécutée (avec 3 comme paramètre)
	fmt.Println(res) // 5 (1 + 3 + 1)
	res = maFonction(4)
	fmt.Println(res) // 10  (5 + 4 + 1)
	res = maFonction(5)
	fmt.Println(res) // 16  (10 + 5 + 1)

	maFonc2 := essai() // deuxième instance de i initialisée à 1. "i est privée et encapsulée dans maFonc2"
	result := maFonc2(1)
	fmt.Println(result) // 3 (1 + 1 + 1)
	result = maFonc2(1)
	fmt.Println(result) //  5  (3 + 1 + 1)
	result = maFonc2(2)	
	fmt.Println(result) //  8  (5 + 2 + 1)	
}

// En fait, avec un peu de théorie :
// On définit une fonction f1 (ici essai) qui retourne une fonction f2 (anonyme)
// On implémente la fonction f1, elle doit respecter le contrat et retourner une fonction prenant un int et retournant un int
// Dans cette implémentation de la fonction f1 on peut être amené à déclarer une variable (i dans ma fonction essai)
// on doit retourner une fonction f2, implémentée de façon anonyme
// c'est là que c'est fort, je peux utiliser une variable déclarée dans f1 (i), dans le corps de f2, elle est bien connue
// Quand j'appelle maFonction, seule la fonction anonyme est appelée, ce qui est normal puisque c'est celle que l'appel à essai() a retourné
// Donc quand j'appelle maFonction, la partie initialisation de i ou print de "coucou" n'est pas appelé, c'est le corps de la fonction anonyme qui est exécuté
// Mais i est connu et stocké dans maFonction
// i est privée et encapsulée dans maFonction, c'est un autre i qui se trouve dans maFonc2


// j'ai voulu faire cet essai2, qui marche : ligne 25, faire un appel à essai2
func essai2() func(int) int {
	//fmt.Println("coucou")
	i := 1
	return func(param int) int {
		i += param + 1
		if i>1000 {return i} // sinon récursif à l'infini
		maFonction := essai2() // j'utilise ma fonction anonyme dans ma fonction anonyme
		res := maFonction(param+1) // +1 sinon récursif à l'infini
		i += res
		fmt.Println(i)		
		return i
	}
}