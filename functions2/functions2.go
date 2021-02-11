package main

// Exemple pas vraiment extraordinaire sur les fonctions, mais juste pour le garder :
// J'adore la façon dont on peut utiliser les fonctions en tant que variables

import "fmt"

func se1() int {
	return 3
}

func se2() int {
	return 4
}

func use(a func() int, b func() int) int {
  return a() + b()
}
//==========================================
func main(){
	f := se1;
	g := se2;
	fmt.Println(f(),g())
	fmt.Println(use(f,g))
	h := mySe1;
	i := mySe2;
	fmt.Println(myUse(h, i, 2, 3))	// should return 3 + 2 + 4 + 3 = 12
	fmt.Println(myOtherUse(h, i, 2, 3))	// should display 5, then 7 (because variadic)
}
//==========================================
func mySe1(x int) int {
	return 3 + x
}

func mySe2(y int) int {
	return 4 + y
}

func myUse(a func(int) int, b func(int) int, val1 int, val2 int) int {
  return a(val1) + b(val2) 
}

func myOtherUse(a func(int) int, b func(int) int, val1 int, val2 int) (int, int) {
  return a(val1), b(val2) 
}

