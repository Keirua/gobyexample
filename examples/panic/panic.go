// Un `panic` signifie typiquement que quelque chose
// s'est mal passé et de manière inattendue. On l'utilise
// principalement pour faire échouer rapidement des
// erreurs qui ne devraient normalement pas arriver, ou
// que nous ne pouvons pas traiter.

package main

import "os"

func main() {
    // Nous allons utilise panic dans les exemples de ce
    // site pour traiter les cas imprévus. Ce
    // programme-ci est le seul du site qui va paniquer
    // durant son déroulement normal.
    panic("a problem")

    // Une utilisation courante de `panic`, c'est
    // d'abandonner si une fonction renvoie une valeur
    // d'erreur que nous ne savons (ou ne voulons) pas
    // traiter. Voici un exemple de panique si on a une
    // erreur inattendue lors de la création d'un nouveau
    // fichier.
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}
