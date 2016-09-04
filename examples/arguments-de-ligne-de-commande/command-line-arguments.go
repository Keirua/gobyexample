// Les [_arguments de ligne de commande_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// sont une manière courante de paramètrer l'exécution des programmes.
// Par exemple, `go run hello.go` envoie les arguments`run` et
// `hello.go` au programme `go`.

package main

import "os"
import "fmt"

func main() {

    // `os.Args` donne accès au arguments brutes de la
    // ligne de commande. A noter que la première valeur
    // de cette slice est le chemin du programme, et
    // `os.Args[1:]` contient les arguments du programme.
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    // On peut accéder aux arguments individuer par leur
    // index.
    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}
