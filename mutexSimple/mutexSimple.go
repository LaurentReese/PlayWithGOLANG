// Ici j'ai deux chaines de caractères et uen chaine de digits, que je veux afficher en parallèle, donc avec des go
// Je ne veux pas que les lettres majuscules et minuscules se croisent
// Par contre je m'autorise à mettre des chiffres n'importe où (n'importe quand)
// Donc j'utilise le mutex pour l'affichage des chaines, mais pas pour l'affichage des digits

package main

import (
	"fmt"
	"sync"
	"time"
)

func displayStringSlowly(myString string) {
	myMutex.Lock()
	defer myMutex.Unlock() // s'il y a quoi que ce soit comme problème dans mon traitement, je ne veux pas rester en dead lock
	// à la place de ma boucle for je pourrais par exemple avoir un accès à une base de données avec des erreurs potentielles
	for _, x := range myString {
		fmt.Printf("%c\n", x)
		time.Sleep(100 * time.Millisecond)
	}
}

func displayDigitsSlowly(myDigits string) {
	for _, x := range myDigits {
		fmt.Printf("%c\n", x)
		time.Sleep(100 * time.Millisecond)
	}
}

var myMutex sync.Mutex

func main() {
	go displayDigitsSlowly("73467425658787908112413453479905785661123095778456347346742565878790811241345347990578566112309577845634")
	go displayStringSlowly("i want that sentence to be written in small letters")
	go displayStringSlowly("I WANT THAT SENTENCE TO BE WRITTEN IN BIG LETTERS")
	time.Sleep(10 * time.Second) // to let it execute and finish
}

