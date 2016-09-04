// En Go, un _tableau_ ou _array_ est une séquence
// numérotée d'éléments d'une longueur donnée

package main

import "fmt"

func main() {

    // Ici, nous créons un tableau `a` qui contiendra
    // exactement 5 `int`. Le type des éléments et la
    // longueur font tous les deux partie du type du
    // tableau. Par défaut, les valeurs des éléments du
    // tableau sont nulles, c'est à dire 0 chez les int
    var a [5]int
    fmt.Println("emp:", a)

    // On peut affecter la valeur à un index particulier
    // avec la syntaxe `array[index] = value`.
    // On obtient sa valeur avec `array[index]`.
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    // La fonction `len`, intégrée au langage, renvoie
    // la longueur du tableau.
    fmt.Println("len:", len(a))

    // On peut utiliser cette syntaxe pour déclarer et
    // initialiser un tableau en une ligne.
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // Les tableaux sont à une dimension, mais on peut
    // les composer pour obtenir des structures de
    // données multi-dimensionnelles.
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
