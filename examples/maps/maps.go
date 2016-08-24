// _Maps_ sont un type [type de donnée associatif](http://en.wikipedia.org/wiki/Associative_array)
// (parfois appelé _hashes_ ou _dicts_ dans d'autres langages).

package main

import "fmt"

func main() {

    // Pour créer une map vide, on utilise la builtin
    //  `make`: `make(map[key-type]val-type)`.
    m := make(map[string]int)

    // On affecte les paires clé/valeur avec la syntaxe
    // `name[key] = val`
    m["k1"] = 7
    m["k2"] = 13

    // Afficher une map via `Println` affichera toutes
    // les paires clé/valeur
    fmt.Println("map:", m)

    // On récupère une valeur avec `name[key]`.
    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    // La builtin `len` renvoie le nombre de paires
    // clé/valeur lorsqu'on l'appelle sur une map
    fmt.Println("len:", len(m))

    // La builtin `delete` supprime d'une map
    // une paire clé/valeur
    delete(m, "k2")
    fmt.Println("map:", m)

    // La seconde valeur de retour optionnelle lorsque
    // l'on récupère une valeur dans une map indique si
    // la clé était présente. Cela peut être utilisé pour
    // lever l'ambiguité entre des clés manquantes et des
    // clés de valeur nulle, comme `0` ou `""`.
    // Ici, nous n'avions pas besoin de la valeur en
    // elle-même, donc nous l'avions ignorée avec
    // l'_identifieur vide_ `_`.
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    // On peut aussi déclarer et initialiser une nouvelle
    // maps sur la même ligne avec cette syntaxe.
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}
