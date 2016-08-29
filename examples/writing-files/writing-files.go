// Ecrire dans des fichiers en Go se fait de manière similaire à ce que nous avons vu pour l'écriture.

package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // Pour commencer, voici comment balancer une chaîne (ou simplement des octets) dans un fichier.
    d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
    check(err)

    // Pour des écritures plus granulaires, on ouvre un fichier pour y écrire.
    f, err := os.Create("/tmp/dat2")
    check(err)

    // Il est idiomatique de reporter la fermeture avec `Close` immédiatement après l'ouverture avec `defer`.
    defer f.Close()

    // On peut écrite des slices d'octets comme on s'y attend avec `Write`.
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    // `WriteString` est également disponible.
    n3, err := f.WriteString("writes\n")
    fmt.Printf("wrote %d bytes\n", n3)

    // Appelez `Sync` pour vider les écriture
    f.Sync()

    // `bufio` fournit des writers avec buffer en complémenter des readers que nous avons vu plus tôt.
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    fmt.Printf("wrote %d bytes\n", n4)

    // Utilisez `Flush` pour vous assurer que toutes les opéraitions avec buffer ont été appliquées au writer sous-jacent.
    w.Flush()

}
