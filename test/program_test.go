package main

import ("testing"
		//"strings"
)

func TestEmptyPassword(t *testing.T) {
	b, err := checkPassword("")
    if b || err==nil {
		t.Fatalf("empty password not well handled")
	}
}

func TestTrivialPassword(t *testing.T) {
    b, err := checkPassword("password")
    if b || err==nil {
		t.Fatalf("trivial password not well handled")
	}
}

func TestShortPassword(t *testing.T) {
    b, err := checkPassword("short")
    if b || err==nil {
		t.Fatalf("short password not well handled")
	}
}
 
func TestLongPassword(t *testing.T) {
    b, err := checkPassword("mot de passe extremement long, mais vraiment vraiment trop trop long, ah ça c'est sûr qu'il est trop long")
    if b || err==nil {
		t.Fatalf("long password not well handled")
	}
}

func TestFirstName(t *testing.T) {
	b, err := checkPassword("Laurent")
    if b || err==nil {
		t.Fatalf("First name as password not well handled")
	}
	b, err = checkPassword("laurent")
    if b || err==nil {
		t.Fatalf("First name as password not well handled")
	}	
	b, err = checkPassword("LAURENT")
    if b || err==nil {
		t.Fatalf("First name as password not well handled")
	}	
}
