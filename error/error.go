package main

// My rules in front of errors (since 30 years of dev)
// 1) DO NOT wait too long, and if possible DO NOT wait AT ALL, before managing errors.
// By the past I have made the mistake to not handle the errors immediately. You want to progress with your program,
// and write important things (but handling errors is also an important thing), and later,
// you don't even think of managing the errors and your program is of bad quality.
// 2) When you meet an error :
// - log it (with line, file, and error message)
// - decide wether your program can or cannot continue the execution with this error (preferably, continue execution).
// - But if it is a serious error, which prevents your program to work (e.g. a network error in a web application), then
//   you may decide to stop the program.
// - if you decide to go on the program execution then...
// 3) You decide to go on the program execution, but :
// - You may be in "degraded" mode, which means you have to deal with this error, that's where point 1 is important
// - in GOLANG : you use defer() to launch a recovering function when necessary
// - in this "defer function", you do some stuff necessary to deal with the error, and then you recover() if you want to go on the execution
// - otherwise the program will stop after this defer function

// Maintenant de façon pratique on peut faire un panic chaque fois qu'on a une erreur qui doit être traitée
// Il faudra faire un grep de panic (ou de log.Panic) et vérifier qu'on a traité TOUS les cas d'erreur
// Important : dans le plan de test, simuler tous les cas d'erreur (tous les panic) : certains ne seront pas faciles à simuler (du moins de façon automatique)

import ("fmt"
	"strconv"
	"log")

var myErrNumber int

func recoveryFunction() {
  fmt.Println("This is recovery function of error ", myErrNumber)
  //recover()
  fmt.Println("This is recovery 2 function of error ", myErrNumber)
}

func executePanic() {
  defer recoveryFunction()
  // ici j'exécute du code
  // bla bla bla bla
  // ce code peut me provoquer une erreur, valant 5 par exemple
  myErrNumber = 5
  //panic("This is Panic Situation " + strconv.Itoa(myErrNumber))
  log.Panic("This is Panic Situation " + strconv.Itoa(myErrNumber))
  fmt.Println("The function executes Completely")  
}

func main() {
  defer a()
  defer b()
  myErrNumber = 0
  executePanic()
  fmt.Println("Main block is executed completely, despite error ", myErrNumber)
}

func a() {
fmt.Println("a")}

func b() {
fmt.Println("b")}

