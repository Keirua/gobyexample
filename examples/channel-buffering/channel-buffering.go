// Par défaut les canaux n'ont pas de buffer, ce qui
// signifie qu'ils acceptent uniquement les envois
// (`chan <-`) s'il y a un receveur
// correspondant (`<- chan`) prêt à recevoir la valeur
// envoyée.
// Les _canaux avec buffer_ acceptent un nombre limité
// de valeurs sans receveur correspondant
// pour ces valeurs.

package main

import "fmt"

func main() {

    // Here we `make` a channel of strings buffering up to
    // 2 values.
    messages := make(chan string, 2)

    // Because this channel is buffered, we can send these
    // values into the channel without a corresponding
    // concurrent receive.
    messages <- "buffered"
    messages <- "channel"

    // Later we can receive these two values as usual.
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
