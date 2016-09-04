// Go a plusieurs types de valeurs, incluant les chaines
// de caractères, les entiers, les nombres flottants, les
// booléens, etc. Voici quelques exemples basiques.
package main

import "fmt"

func main() {

    // Les chaines de caractères, que l'on peut
    // concaténer avec `+`.
    fmt.Println("go" + "lang")

    // Les entiers et les flottants.
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    // Les booléens, avec les opérateurs booléens tels qu'on les attend.
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
