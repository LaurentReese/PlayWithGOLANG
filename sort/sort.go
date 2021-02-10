package main

import ("fmt"
		"sort"
	)

type personne struct {
	nom string
	prenom string
	dateDeNaissance int
	numSecu uint
}

type myIntSlice []int // pour mémoire

type meilleurTri []personne // synonyme, mais va me permettre de définir une autre méthode de tri

func (a meilleurTri) Len() int { return len(a) }
func (a meilleurTri) Less(i, j int) bool {
	// Je vais jusqu'au numéro de sécu si tous les champs sont les mêmes
	if a[i].nom != a[j].nom {return a[i].nom < a[j].nom}
	if a[i].prenom != a[j].prenom {return a[i].prenom < a[j].prenom}
	if a[i].dateDeNaissance != a[j].dateDeNaissance {return a[i].dateDeNaissance < a[j].dateDeNaissance}	
	if a[i].numSecu != a[j].numSecu {return a[i].numSecu < a[j].numSecu}
	return true // do not swap if all the fields are equals
}
func (a meilleurTri) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	test := []int{1,3,2}
	test = append(test, 5, 4)
	fmt.Println(test)	
	sort.Ints(test)	// sort.Ints, sort.Float64s, sort.Strings sont des fonctions toutes prêtes
	fmt.Println(test)		
	// Mais je peux définir ma propre fonction de comparaison, par exemple pour trier dans l'ordre inverse
	sort.Slice(test, func(i, j int) bool {
		return test[i] > test[j]
	})
	fmt.Println(test)			

	//var tabPersonnes plusieursPersonnes
	// je peux aussi faire
	var tabPersonnes []personne
//	tabPersonnes = append(tabPersonnes,	{"REESE", "Laurent", 19680508, 168050608809410 },
//			 {"CASANOVA","Sandrine", 19700123, 270012008809410 })
			 
	tabPersonnes = append(tabPersonnes,
				personne{"REESE", "Laurent", 19680208, 168050608809411 },
			 	personne{"CASANOVA","Sandrine", 19700223, 270012008809411 },
				personne{"AAAESE", "Laurent", 19680308, 168050608809419 },
				personne{"CCCBBBCASANOVA","Sandrine", 19700323, 270012008809419 },
				personne{"REESE", "Laurent", 19680908, 168050608809418 },
				personne{"CASANOVA","Sandrine", 19700923, 270012008809418 },
				personne{"REESE", "Laurent", 19680808, 168050608809412 },
				personne{"CASANOVA","Sandrine", 19700823, 270012008809412},
				personne{"REESE", "Laurent", 19680608, 168050608809415 },
				personne{"CASANOVA","Sandrine", 19700623, 270012008809415 },
				personne{"DDDREESE", "Laurent", 19680708, 168050608809413 },
				personne{"BBBCASANOVA","Sandrine", 19700723, 270012008809413 } )
	fmt.Println(tabPersonnes)
	fmt.Println()

	sort.Slice(tabPersonnes, func(i, j int) bool {
			return tabPersonnes[i].nom < tabPersonnes[j].nom })
	fmt.Println(tabPersonnes)
	fmt.Println()	

	sort.Sort(meilleurTri(tabPersonnes))
	fmt.Println(tabPersonnes)
	fmt.Println()	
	
	// something else, my "own" sort function
	myTab := []int{12, 5, 4, 1, 2, 3, 8, 7, 6, 9,0}
	var i int
	var j int

	for i = 0; i < len(myTab)-1; i++ {
		for j = 0; j < len(myTab)-1-i; j++ {
			if myTab[j] > myTab[j+1] {
				myTab[j], myTab[j+1] = myTab[j+1], myTab[j] // do the swap
			}
		}
	}
	
	for index, item := range myTab {
		fmt.Printf("myTab[%d]=%d\n",index,item)
	}	
}


