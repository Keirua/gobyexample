// Go offre de base un support des [expressions régulières](https://fr.wikipedia.org/wiki/Expression_rationnelle) (ou regex, pour regular expression).
// Voici quelques exemples de tâches courantes sur les `regex` en Go.

package main

import "bytes"
import "fmt"
import "regexp"

func main() {

    // Ceci teste si un motif ("p([a-z]+)ch") correspond à une chaîne.
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

    // Au dessus nous avec utiliser un motif directement, mais pour d'autres tâches il faut la compile en une structure `Regexp` optimisée.
    r, _ := regexp.Compile("p([a-z]+)ch")

    // Plusieurs méthodes sont disponibles. Voici le même test que celui vu plus tôt.
    fmt.Println(r.MatchString("peach"))

    // Ceci trouve l'association pour la regex.
    fmt.Println(r.FindString("peach punch"))

    // Ceci trouve également la première correspondance mais renvoie l'indice de début et fin de la correspondance, au lieu de renvoyer le text correspondant.
    fmt.Println(r.FindStringIndex("peach punch"))

    // La variante `Submatch` inclue des informations sur les associations au niveau du pattern entier, ainsi que les sous-correspondances à l'intérieur du motif.
    // Par exemple ici, cela renverra des informations à la fois sur `p([a-z]+)ch` et `([a-z]+)`.
    fmt.Println(r.FindStringSubmatch("peach punch"))

    // De la même manière, cela donne des inforamtions sur les indices de correspondances globales et les sous-correspondances.
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

    // Les variantes  `All` de ces fonctions s'appliquent à toutes les correspondances du texte d'entrée, pas juste la première.
    // Par exemple pour trouver toutes les correspondances d'une regex.
    fmt.Println(r.FindAllString("peach punch pinch", -1))

    // Ces variantes `All` sont également disponibles pour toutes les autres méthodes vues plus haut
    fmt.Println(r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))

    // Fournir un entier non négatif en second argument à ces fonctions limite le nombre de résultats.
    fmt.Println(r.FindAllString("peach punch pinch", 2))

    // Nos examples ci-dessus avaient comme arguments des strings et avaient des noms comme `MatchString`. On peut également fournir des arguments `[]byte` et oublier le `String` des nombres de fonction.
    fmt.Println(r.Match([]byte("peach")))

    // Lorsqu'on crée des constrantes avec des expressions régulières, on peut utiliser la variation `MustCompile`  de `Compile`. Un simple `Compile` ne marchera pas car il a deux valeurs de retour.
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)

    // Le package `regexp` peut également être utilisé pour remplacer des sous-ensembles de chaînes avec d'autres valeurs.
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    // La variante `Func` permet de transformer les textes de correspondance selon une fonction donnée.
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}
