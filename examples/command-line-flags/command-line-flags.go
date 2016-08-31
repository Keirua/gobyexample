// Les options de [ligne de commande](https://fr.wikipedia.org/wiki/Interface_en_ligne_de_commande#Option)
// sont une manière courante de préciser des options pour des programmes en ligne de commande.
// Par exemple, dans `wc -l`, le `-l` est une option.

package main

// Go fournit un package `flag` qui permet d'analyser des options simples. Nous utiliserons ce package pour implémenter notre programme d'exemple.
import "flag"
import "fmt"

func main() {

    // Des déclarations d'options simples sont disponibles pour les string, les entiers et les booléens. Ici on déclare une option `word` avec comme valeur par défaut `"foo"`, et une courte description. La fonction `flag.String` renvoie un pointeur sur la chaîne (pas une chaîne directement). Nous verrons comment utliser ce pointeur après.
    wordPtr := flag.String("word", "foo", "une chaine")

    // On déclare des options `numb` et `fork`, en utilisant une approche similaire.
    numbPtr := flag.Int("numb", 42, "un int")
    boolPtr := flag.Bool("fork", false, "un booleen")

    // Il est également possible de déclarer l'option en utilisant une variable déjà déclarée ailleurs dans le programme. A noter que l'on doit passer un pointeur à la fonction de déclaration d'option.
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // Une fois que toutes les options sont déclarées, il faut appeller `flag.Parse()` pour exécuter l'analyse des paramètres de ligne de commande.
    flag.Parse()

    // Ici, on va afficher les options analysées ainsi que les éventuels paramètres restants. A noter qu'il faut déréférencer les pointeurs, par ex. avec `*wordPtr` pour obtenir la véritable valeur des options.
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}
