package main

import "fmt"
import "os"
import "errors"
import "log"

func ignoreErrors() {
  fmt.Println("The error is not a serious error")
  recover()
}

func check(err error) {
  // Verify if there is an error and then launch a panic to log and stop the program
  if err!=nil {
    log.Println(err)    
    panic(err)
  }
}

func treatArgOne() {
  // defer ignoreErrors() // keep to "ignore" (or repear errors), comment it to log and stop the program if error
  var err error
  // Ici je crée ma propre erreur
  if len(os.Args)==1 { err = errors.New("Please enter a filepath as parameter") }
  check(err)
  var node string = os.Args[1]
  _ , err = os.Stat(node)
  // Ici c'est une erreur système comme quoi le fichier ou répertoire n'existe pas
  check(err)
  fmt.Println(node,"exists")
}

func main() {
  treatArgOne()
  fmt.Println("Here I am to continue")
}

