// Le package `sort` implémente le tri pour les type de
// base et ceux définis par l'utilisateur. Regardons
// d'abord le tri pour les types du langage.

package main

import "fmt"
import "sort"

func main() {

    // Les méthodes de tri sont spécifiques à un type ;
    // voici un exemple pour les chaines de caractères. A
    // noter que le tri se fait sur place, donc la slice
    // fournie est modifiée et la fonction n'en renvoit 
    // pas une nouvelle.
    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    // Un exemple de tri sur les `int`.
    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    // On peut également utiliser sort pour tester si 
    // une slice est déjà triée.
    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}
