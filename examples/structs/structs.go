// Les _structures_ en Go sont des collections de champs 
// typés.
// Elles sont utiles pour regrouper ensemble les données.

package main

import "fmt"

// Cette structure `person` a des champs `name` et `age`.
type person struct {
    name string
    age  int
}

func main() {

    // Cette syntaxe crée une nouvelle structure.
    fmt.Println(person{"Bob", 20})

    // On peut nommer les champs lorsqu'on initialise
    //  une structure
    fmt.Println(person{name: "Alice", age: 30})

    // Les champs non spécifiés seront de valeur nulle.
    fmt.Println(person{name: "Fred"})

    // Un préfixe `&` fournit un pointeur sur la structure.
    fmt.Println(&person{name: "Ann", age: 40})

    // On accède aux champs de la structure avec un point
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    // On peut aussi utiliser le point avec des pointeurs
    // de structure : les pointeurs sont automatiquement déréférencés.
    sp := &s
    fmt.Println(sp.age)

    // Les structures sont mutables.
    sp.age = 51
    fmt.Println(sp.age)
}
