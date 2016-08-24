// Go support les [_fonctions anonymes_](https://fr.wikipedia.org/wiki/Fonction_anonyme),
// qui peuvent former une [closure](https://fr.wikipedia.org/wiki/Fermeture_(informatique)) (fermeture ou clotûre en français.)
// Les fonctions anonymes sont utiles lorsque l'on veut 
// définir des fonctions à la volée, sans avoir à les
//  nommer.

package main

import "fmt"

// Cette fonction `intSeq` renvoie une autre fonction, 
// qui est définie anonymement dans le corps de `intSeq`.
// La fonction retournée _embarque_ la variable `i` pour 
// former une closure.
func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func main() {

    // On appelle `intSeq`, et assigne le résultat (une 
    // function) à `nextInt`. Cette fonction capture sa
    // propre valeur de `i`, qui sera mise à jour à 
    // chaque fois que l'on appelle `nextInt`.
    nextInt := intSeq()

    // Voyez l'effet de la closure en appellant `nextInt`
    // plusieurs fois.
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // Pour confirmer que l'état est unique à cette 
    // fonction particulière, créez et testez en une
    // nouvelle.
    newInts := intSeq()
    fmt.Println(newInts())
}
