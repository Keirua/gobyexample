// Le _select_ de go permet d'attendre plusieurs
// opérations sur les canaux. Combiner les goroutines et
// les canaux est une fonctionnalité puissante de Go.

package main

import "time"
import "fmt"

func main() {
    // Pour notre exemple, nous ferons des select à
    // travers deux canaux.
    c1 := make(chan string)
    c2 := make(chan string)

    // Chaque canal va recevoir une valeur après un
    // certain temps, pour simuler une opération
    // bloquante (par ex. un appel RPC) qui s'exécute
    // dans une goroutine concurrente.
    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "two"
    }()

    // Nous utilisons `select` pour attendre ces deux valeurs simultanément, en affichant chacune d'elle dès qu'elle arrive.
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
