// Go supporte la [récursivité](https://fr.wikipedia.org/wiki/R%C3%A9cursivit%C3%A9)
// Voici l'exemple classique de la factorielle.

package main

import "fmt"

// Cette fonction `fact` s'appelle elle-même jusqu'à ce
// qu'elle atteigne le cas de base, `fact(0)`.
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7))
}
