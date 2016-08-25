// _range_ itère sur les éléments de différents types de
// structures de données. Voyons comment utiliser `range`
// avec certaines des structures de données que nous
// avons déjà rencontré.

package main

import "fmt"

func main() {
    // Ici on utilise `range` pour additionner les
    // nombres d'une slice.
    // Cela fonctionne de la même manière avec un array.
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // `range` sur les array et les slices fournit à la
    // fois la clé et la valeur pour chaque entrée.
    // Au dessus, nous n'avons pas besoin de la clé, donc
    // nous l'avions ignoré avec l'identifieur blanc `_`.
    // On a cependant parfois besoin de le récupérer.
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // `range` sur une map itère sur les paires
    // clé/valeur.
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // `range` sur une chaine itère sur les points de
    // code [Unicode](https://fr.wikipedia.org/wiki/Unicode).
    // La première valeur est l'index de l'octet de
    // départ de la `rune`, et le second la rune elle-même.
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
