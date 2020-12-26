// But du programme : je remplis un énorme tableau de valeurs aléatoires, et je calcule la somme de toutes les valeurs de façon concurentielle
// Je mesure le temps de calcul avec ou sans parallélisme, histoire de vérifier que le parallélisme apporte quelque chose
// Conclusion : oui ça marche, je gagne un facteur 3 en utilisant 4 processeurs par rapport à la version sans parallélisme

package main

import (
    "runtime"
	"fmt"
	"math/rand"
	"time"
)

const MY_SIZE = uint64(1 << 28) // power of 2 is preferable

var my_array []uint64
var sum uint64

func main() {
	my_array = make([]uint64, MY_SIZE, MY_SIZE) // en fait c un slice
	rand.Seed(time.Now().UnixNano()) // to reset random generator

	cores := uint64(runtime.NumCPU()) // plus tard je le forcerai à 1 pour mesurer le temps sans parallélisme
	// cores = 1 // si je veux le forcer pour comparer les temps "parallèle" vs "non parallèle"
	myRange := MY_SIZE/cores	
	fmt.Println("Nombre de processeurs (logiques ou physiques) :", cores)
	// Moi g que 2 cpus : donc à essayer sur l'ordi de Thomas ou Nathan
	// pour vérifier mon algo

	// (for the moment) I want no concurrency to fill in the array
	var i uint64
	for i=0;i<MY_SIZE;i++ {
		my_array[i] = uint64(rand.Intn(400000))
	}

	// Now I want concurrency to compute a sum of all those ints
	start := time.Now()
	var coeur uint64
	c := make(chan uint64)
	for coeur=0; coeur < cores; coeur++ {
		go computeSum(coeur, myRange, c) // n sommes intermédiaires calculées en parallèle
	}
	// Il faut récupérer ce qui est bufferisé sur le channel autant de fois qu'on a bourré le channel
	for coeur=0; coeur < cores; coeur++ { // somme des n sommes intermédiaires
		sum += <- c // point de synchronisation du calcul de toutes les sommes intermédiaires (envoyées sur le canal c)
	}
	fmt.Printf("Calcul effectué en %v\n", time.Since(start))	
	fmt.Printf("Somme de %d nombres = %d\n", MY_SIZE, sum)
}

func computeSum(coeur uint64, theRange uint64, c chan uint64) {
	var i, oneSum uint64
	for i=coeur * theRange; i < (coeur+1) * theRange; i++ {
		oneSum += my_array[i]
	}
	c <- oneSum // Le channel c est "bourré" avec la somme intermédiaire (il faudra le "dépiler" pour récupérer cette somme)
}
