// Il faudra lancer la commande "go test"
// ou "go test -v" pour verbose

package main
// https://golang.org/doc/tutorial/add-a-test
// https://www.synbioz.com/blog/tech/test-en-go
import "fmt"
import "errors"
import "strings"

// Je pourrais tout gérer avec juste l'erreur qui vaut nil si pas d'erreur
func checkPassword(password string) (bool, error) {
	// commenter certaines des lignes ci-dessous pour que certains tests ne passent plus	
	if password == "password" || password == "Password" || password == "PASSWORD" {return false, errors.New("mot de passe trivial") }
	//if len(password) < 6 {return false, errors.New("mot de passe trop court") }	
	if len(password) > 100 {return false, errors.New("mot de passe trop long") }		
	if strings.ToUpper(password) == "LAURENT" {return false, errors.New("prenom en tant que mot de passe est interdit") }		
	return true, nil
}

func main() {
	password := "Laurent"
	b, err := checkPassword(password) // err est de type *errors.errorString
	if b {fmt.Println("Le mot de passe ==>", password,"<== est accepté")
	} else {fmt.Println("Le mot de passe ==>", password,"<== n'est pas accepté,", err)
			panic(err)}

	password = "qdfjklqsjfoqzefjlqsdfj"
	b, err = checkPassword(password)
	if b {fmt.Println("Le mot de passe ==>", password,"<== est accepté")
	} else {fmt.Println("Le mot de passe ==>", password,"<== n'est pas accepté,", err)
			panic(err)}
}