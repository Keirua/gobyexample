// On a souvent besoin que les programmes réalisent des
// opérations sur des collections de données, comme
// sélectionner tous les éléments qui correspondent à un
// prédicat, ou associer les éléments à une autre
// collection selon une fonction donnée.

// Dans certains langages il est idiomatiques d'utiliser
// des structures de données et des algorithmes
// [génériques](https://fr.wikipedia.org/wiki/G%C3%A9n%C3%A9ricit%C3%A9).
// Go ne supporte pas les types génériques; en Go, il est
// courant de fournir des fonctions sur les collections
// spécialement pour un type donné, si et lorsque c'est
// spécifiquement nécessaire pour le programme.

// Voici quelques exemples de fonctions sur les 
// collections pour des slices de `strings`. Vous pouvez
// utiliser ces exemples pour concevoir vos propres
// fonctions. A noter que dans certains cas, il peut être
// plus clair d'utiliser le code de manipulation de
// collections inline, plutôt que créer et appeler des
// fonctions auxilliaires.

package main

import "strings"
import "fmt"

// Renvoie le premier index où se trouve la chaîne
// cible `t`, ou -1 si aucune correspondance
// n'est trouvée.
func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

// Renvoie `true` si la chaîne cible t est dans la slice.
func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

// Renvoie `true` si une des chaines de la slice
// satisfait le prédicat `f`.
func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

// Renvoie `true` si toutes les chaines de la
// slice satisfassent le prédicat `f`.
func All(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

// Renvoie une nouvelle slice contenant toutes les
// chaines de la slice qui satisfasse le prédicat `f`.
func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

// Renvoie une nouvelle slice contenant le résultat de
// l'application de la fonction `f` à chacune des chaines
// de la slice de départ.
func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

func main() {

    // Maintenant, on essaye nos différentes fonctions
    // sur les collections.
    var strs = []string{"peach", "apple", "pear", "plum"}

    fmt.Println(Index(strs, "pear"))

    fmt.Println(Include(strs, "grape"))

    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
    }))

    // Les exemples ci-dessus ont tous utilisé des
    // fonctions anonymes, mais vous pouvez également
    // utiliser des fonctions nommées (à condition
    // qu'elles soient du bon type).
    fmt.Println(Map(strs, strings.ToUpper))

}
