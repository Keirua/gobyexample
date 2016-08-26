// Les envois et receptions basiques sur les canaux
// sont bloquants.
// Cependant, on peut utilise `select` avec une clause
// `default` pour implémenter des envois et réceptions 
// _non bloquants_, et même des `select` non bloquants
// sur plusieurs canaux.

package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    // Voici une réception non bloquante. Si une valeur 
    // est disponible sur `messages`, alors `select` va
    // prendre le cas `<-messages` avec cette valeur.
    // Si non, il prendra immédiatement le cas `default`.
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    // Un envoi non bloquand marche de manière similaire.
    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    // On peut utiliser plusieurs `case` au dessus de 
    // la clause `default` pour implémenter un select non
    // bloquant multiple. Ici, on essaye de recevoir des
    // messages de manière non bloquante sur les canaux 
    // `messages` et `signals` simultanément.
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}
