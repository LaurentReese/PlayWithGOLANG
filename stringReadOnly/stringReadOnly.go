package main 
  
import "fmt"

func changeByteArray(b []byte) {
    b[1]= 'G' // changement effectif, sans avoir à retourner b
	// b += "changed"
	}

func changeString(s string) {
    //s[1]= 'G' is forbidden if it's as tring, but authorised if it'a a byte array
	s += "changed" // changement local, à moins que je retourne s
	}

// Main function 
func main() { 
  
    // Creating and initializing a string 
    // using shorthand declaration 
    mystr := []byte("Welcome to GOLANG")
	//mystr := string("Welcome to GOLANG")	
  
    fmt.Println("String:", mystr) 
  
    //mystr[1]= 'G' is forbidden if it's as tring, but authorised if it'a a byte array
	//changeString(mystr)
	changeByteArray(mystr)	
	
    fmt.Println("String:", string(mystr)) 
     
} 
