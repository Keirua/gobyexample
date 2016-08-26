// On peut utiliser les canaux pour synchronier 
// l'exécution à travers des goroutines. Voici un exemple
// qui utilise une réception bloquante pour attendre
// qu'une goroutine se termine.

package main

import "fmt"
import "time"

// Voici la fonction qui va tourner dans une goroutine.
// Le canal `done` sera utilisé pour notifier une autre
// goroutine que le travail de cette fonction est terminé
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    // Envoie une valeur pour notifier qu'on a terminé.
    done <- true
}

func main() {
    // On démarre la goroutine, en lui fournissant 
    // le canal à notifier.
    done := make(chan bool, 1)
    go worker(done)

    // On bloque jusqu'à ce qu'on ait reçu une
    // notification sur ce canal.
    <-done
}
