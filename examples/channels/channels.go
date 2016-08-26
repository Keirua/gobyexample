// Les _canaux_ (channels) sont des tuyaux qui relient
// des goroutines concurrentes. On peut envoyer des
// valeurs d'une goroutine à une autre via un canal

package main

import "fmt"

func main() {

    // On crée un nouveau canal avec `make(chan val-type)`.
    // Les canaux sont typés avec les valeurs qu'ils
    // transportent.
    messages := make(chan string)

    // On _envoie_ une valeur dans un canal avec la
    // syntaxe `canal <-`. Ici on envoie `"ping"`  dans
    // le canal `messages` que l'on a créé au dessus,
    // depuis une nouvelle goroutine.
    go func() { messages <- "ping" }()

    // La syntaxe`<-canal` _recoit une valeur depuis un canal.
    // Ici on affiche le message `"ping"` envoié au
    // dessus, et l'affiche.
    msg := <-messages
    fmt.Println(msg)
}
