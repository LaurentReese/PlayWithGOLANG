// see https://tour.golang.org/methods/18

package main

import "fmt"

type IPAddr [4]byte

// "String() string" method added to IPAddr.
// C'est quelque chose de très courant en GO: ajouter la méthode String() sur les types importants,
// car après c'est utilisé par le logger, les fonctions d'erreur (panic), le Printf avec %v, etc
/*func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0],ip[1],ip[2],ip[3])
}*/

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

/*
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
*/
