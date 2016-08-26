// Dans un exemple [précédent](range), nous avons vu
// comment `for` et `range` permettent d'itérer sur des
// structures de données simples.
// On peut également utiliser cette syntaxe pour itérer
// sur des valeurs reçues depuis un canal.

package main

import "fmt"

func main() {

    // Nous allons itérer sur 2 valeurs dans le canal `queue`.
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    // Ce `range` itère sur chaque élément à mesure
    // qu'il est reçu dans `queue`.
    // Comme nous avons fermé le canal plus haut,
    // l'itération se termine après réception des deux
    // éléments. Si nous ne l'avions pas fermé, nous
    // bloquerions sur la réception d'une troisième
    // valeur dans la boucle.
    for elem := range queue {
        fmt.Println(elem)
    }
}
