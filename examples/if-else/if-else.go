// Les conditions avec `if` et `else` sont simples en Go

package main

import "fmt"

func main() {

    // Voici un exemple simple.
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    // On peut avec un `if` sans `else`.
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    // Une déclaration peut précéder la condition; Toute
    // variable déclarée là est disponible dans toutes
    // les branches
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}

// A noter qu'il n'y a pas besoin de parenthèses autour
// des conditions en Go, mais que les accolades sont
// requises.
