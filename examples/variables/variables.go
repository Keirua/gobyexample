// En Go, les _variables_ sont déclarées explicitement et
// utilisées par le compilateur, par exemple pour vérifier
// que le type de retour des appels de fonction
// est correct

package main

import "fmt"

func main() {

    // `var` déclare une ou plusieurs varaibles.
    var a string = "initial"
    fmt.Println(a)

    // On peut déclarer plusieurs variables à la  fois
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go déduira le type des variables non initialisées
    var d = true
    fmt.Println(d)

    // Les variables déclarées sans être initialisées ont
    // une _valeur nulle_. Par exemple, la valeur nulle
    // d'un `int` est `0`.
    var e int
    fmt.Println(e)

    // La syntaxe `:=` est un raccourci pour déclarer et
    // initialiser une variable, par exemple pour
    // `var f string = "short"` ici.
    f := "short"
    fmt.Println(f)
}
