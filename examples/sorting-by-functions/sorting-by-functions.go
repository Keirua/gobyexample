// Parfois on veut trier une collection selon un autre
// critère que l'ordre naturel. Par exemple, imaginons
// que l'on veuille trier des chaines par leur longueur
// plutôt que par ordre alphabétique. Voici un exemple
// de tri sur mesure en Go.

package main

import "sort"
import "fmt"

// Afin de trier selon une méthode custom en Go, nous
// avons besoin d'un type correspondant.
// Ici, on crée le type `ByLength`, qui est simplement un
// alias pour le type `[]string` builtin.
type ByLength []string

// Nous implémentons `sort.Interface` - `Len`, `Less`, et
// `Swap` - sur notre type, afin de pouvoir utiliser la
// fonction générique Sort du package `sort`.
// `Len` et `Swap` seront généralement similaires à
// travers les types, et `Less` va implémenter la logique
// de notre méthode de tri.
// Dans notre cas, nous voulons trier par ordre de
// longueur de chaine croissante, donc on utilise
// `len(s[i])` et `len(s[j])` ici.
func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

// Avec tout ceci en place, nous pouvons maintenant
// implémenter notre tri sur mesure en convertissant la
// slice de départ `fruits` en `ByLength`, et ensuite
// utiliser la méthode `sort.Sort` sur cette slice typée.
func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}
