// Lorsqu'on utilise des canaux comme paramètres de 
// fonctions, on peut spécifier si un canal est censé 
// uniquement envoyer ou recevoir des valeurs. Cette 
// spécificité augmente la sécurité du typage du 
// programme.

package main

import "fmt"

// Cette fonction `ping` accepte seulement un canal qui
// envoie des valeurs. On aurait une erreur de 
// compilation si on lui essayait de recevoir sur ce 
// canal.
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// La fonction `pong` accepte un canal pour recevoir
// (`pings`) et un second pour les envois (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
