// Lire et écrire dans les fichiers sont des tâches de base nécessaires dans de nombreux programmes en Go. Nous allons tout d'abord regarder comment lire dans des fichiers.

package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

// Lire dans des fichiers nécessite de vérifier les erreurs pour la plupart des appels. Cette fonction va nous permettre de simplifier la gestion d'erreurs ensuite.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // Peut-être une des tâches de lecture les plus simples : aspirer le contenu entier d'un fichier en mémoire.
    dat, err := ioutil.ReadFile("/tmp/dat")
    check(err)
    fmt.Print(string(dat))

    // On veut souvent plus de contrôle sur comment et quelle partie d'un fichier sont lues. Pour ces tâches, on commence par ouvrir le fichier avec `Open` pour obtenir un `os.File` que l'on va manipuler.
    f, err := os.Open("/tmp/dat")
    check(err)

    // Lit quelques octets au début du fichier. Permet d'en lire jusqu'à 5, mais regardez combien sont effectivement lus.
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))

    // On peut aussi se déplacer avec `Seek` à une position connue dans le fichier, et lire depuis là avec `Read`.
    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

    // Le package `io` fournit des fonctions qui peuvent être utiles pour la lecture de fichiers. Par exemple, les lectures comme celle ci-dessus peuvent être réalisées de manière plus robustes avec `ReadAtLeast`.
    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    // Il n'y a pas de retour au début avec `rewind`, mais on peut le faire `Seek(0, 0)`.
    _, err = f.Seek(0, 0)
    check(err)

    // Le package `bufio` implémente un lecteur avec buffer qui peut être utile à la fois pour ses performances sur les petites lectures et par les méthodes additionnelles qu'il propose.
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    // On ferme le fichier quand on a terminé. En général, on le programme automatiquement après l'ouverture avec `defer`.
    f.Close()

}
