// Go supporte de base les _valeurs de retour multiples_.
// On s'en sert souvent en Go idiomatique, par exemple 
// pour renvoyer le résultat et la valeur d'erreur d'une
// fonction.

package main

import "fmt"


// Le `(int, int)` dans cette signature de fonction 
// montre que la fonction renvoie deux `int`s.
func vals() (int, int) {
    return 3, 7
}

func main() {

    // Ici nous utilisons les deux valeurs de retour
    // différentes de l'appel avec un _assignement 
    // multiple_.
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // Si vous voulez seulement un sous-ensemble des 
    // valeurs de retour, vous pouvez utiliser 
    // l'identifieur blanc `_`.
    _, c := vals()
    fmt.Println(c)
}

// todo: paramètres de retour nommé
// todo: retours nus
