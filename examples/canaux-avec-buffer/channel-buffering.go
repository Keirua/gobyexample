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
    // Ici on crée avec `make` un canal avec buffer, qui
    // accumule jusqu'à 2 string.
    messages := make(chan string, 2)

    // Comme ce canal a un buffer, nous pouvons envoyer
    // ces deux valeurs dans le canal sans une réception
    // concurrente correspondante.
    messages <- "buffered"
    messages <- "channel"

    // Ensuite, on peut recevoir ces deux valeurs comme
    // d'habitude.
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
