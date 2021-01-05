// Ce n'est pas une création de classes et d'objets comme en C++
//
// On crée une interface (une classe pour faire la comparaison avec le C++)
// Dans cette interface on définit des fonctions (fonctions virtuelles pour faire la comparaison avec le C++)
// Ensuite on crée des types, puis on ajoute à ces types les fonctions de l'interface (on crée ainsi une sorte d'héritage au sens C++)
// On pourra ainsi affecter ces types à des variables de type interface et travailler sur des objets en faisant abstraction de leur type exact (on sait juste qu'ils sont "hérités" de l'interface)

// Je rajoute aussi un Stringer pour afficher mes slices d'objets

package main

import ("fmt"
"reflect"
"strconv"
"runtime"
)

type vehicle interface {
	getSpeed() uint
	getNbWheels() uint
	getPrice() uint
}

type voiture struct {}
type velo struct {}
type fusee struct {}

func (voiture) getSpeed() uint {	return 100 }
func (velo) getSpeed() uint {	return 20}
func (f fusee) getSpeed() uint {	return 2000} // f est superflu si je n'ai pas besoin des ses champs dans la fonction getSpeed

func (voiture) getNbWheels() uint {	return 4}
func (velo) getNbWheels() uint {	return 2}
func (fusee) getNbWheels() uint {	return 0}

func (voiture) getPrice() uint {	return 30000}
func (velo) getPrice() uint {	return 300}
func (fusee) getPrice() uint { return 30000000}
// Stringers
func (fusee) String() string { return "Une fusee " }
func (v voiture) String() string { return "Une voiture qui coute " + strconv.Itoa(int(v.getSpeed())) + " Euros "}
func (v velo) String() string { return "Un velo " }

type vehicles []vehicle
type pVehicle *vehicle
type pVehicles []pVehicle

func (p pVehicles) String() string {
	var vo voiture
	var ve velo
	var fu fusee
	var myString string
	for _, one := range p {
		//fmt.Printf("type=%T\n", *one)
		switch (reflect.TypeOf(*one) ) {
			case reflect.TypeOf(vo) :
				vo = (*one).(voiture)	// type assertion (je suis sûr que ça se passe bien car je vient de tester le type grâce à TypeOf)
				myString += vo.String()
			case reflect.TypeOf(ve) :
				ve = (*one).(velo)		// type assertion		
				myString += ve.String()					
			case reflect.TypeOf(fu) :
				fu = (*one).(fusee)		// type assertion	
				myString += fu.String()
			default :
				_, fileName, fileLine, _ := runtime.Caller(1)
				s := fmt.Sprintf("véhicule inconnu ==> %s:%d", fileName, fileLine)
				panic(s)
		}
	}
	return myString
}

func main() {
	// 1) slice d'objets héritant tous de vehicle
	var v voiture
	fmt.Printf("%d\n", v.getSpeed())
	var ve velo
	fmt.Printf("%d\n", ve.getSpeed())
	var fu fusee
	fmt.Printf("%d\n", fu.getSpeed())	
	vehicules := vehicles{v, ve, fu}  // Je pourrais aussi faire "vehicules := []vehicle{v, ve, fu}"
	fmt.Println(vehicules)

	//var otherVehicles []*vehicle
	//var otherVehicles []pVehicle
	// 2) slice de pointeurs d'objets héritant tous de vehicle	
	var otherVehicles pVehicles
	z :=  vehicle(v)
	otherVehicles = append(otherVehicles, &z)
	zz := vehicle(ve)
	otherVehicles = append(otherVehicles, &zz)
	zzz := vehicle(fu)
	otherVehicles = append(otherVehicles, &zzz)
	fmt.Println(otherVehicles)	
}
