// Les _fonctions_ sont centrales en Go. Nous allons les 
// découvrir à travers quelques exemples différents.

package main

import "fmt"

// Voici une fonction qui prend deux `int` en paramètres,
// et renvoie leur somme, un int.
func plus(a int, b int) int {

    // Go a besoin de retours explicites : il ne renverra
    // pas automatiquement la valeur de la dernière
    // expression.
    return a + b
}

// Quand vous avez plusieurs paramètres consécutifs du 
// même type, vous pouvez vous passez des déclarations
// de type jusqu'au dernier, qui le déclare.
func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {
    // On appelle une fonction comme on s'y attend, avec
    // `nom(arguments)`
    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}
