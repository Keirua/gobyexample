// `for` est la seule manière de faire une boucle en Go.
// Voici trois types basiques de boucles `for`.

package main

import "fmt"

func main() {

    // La plus simple, avec une unique condition.
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // La boucle `for` classique, en 3 étapes :
    // initialisation/condition/incrémentation.
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    // `for` sans condition va boucler indéfiniment,
    // jusqu'à ce qu'on `break` pour en sortir, ou
    // qu'un `return` fasse sortir de la fonction
    // correspondante.
    for {
        fmt.Println("loop")
        break
    }
}
