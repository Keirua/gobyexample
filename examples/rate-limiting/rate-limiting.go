// _[Limiter le débit](https://fr.wikipedia.org/wiki/Rate_limiting)_
// est un important mécanisme pour controller l'utilisation des ressources et maintenir la qualité de service. Go supporte élégamment la limitation de débit avec les goroutines, les canaux et les [tickers](tickers).

package main

import "time"
import "fmt"

func main() {

    // Nous allons d'abord regarder une limitation de débit simple. Supposons que l'on veuille limiter notre gestion des requêtes entrantes. Nous allons servir ces requêtes depuis un canal du même nom.
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    // Ce canal `limiter` va recevoir une valeur toutes les 200 millisecondes. C'est le régulateur de notre système de limitation du débit.
    limiter := time.Tick(time.Millisecond * 200)

    // En bloquant sur une réception du canal `limiter` avant de servir chaque requête, on se limite à une requête toutes les 200ms.
    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }

    // On pourrait vouloir permettre de courtes augmentations du nombre de requêtes, tout en préservant la limitation globale. On peut pour cela mettre un buffer sur le canal limitant. Ce canal `burstyLimiter` va permettre de recevoir jusqu'à 3 évènements.
    burstyLimiter := make(chan time.Time, 3)

    // On remplit le canal pour représenter l'augmentation considérée.
    // Fill up the channel to represent allowed bursting.
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    // Toutes les 200ms nous essaierons d'ajouter une nouvelle valeur dans `burstyLimiter`, jusqu'à sa limite de 3.
    go func() {
        for t := range time.Tick(time.Millisecond * 200) {
            burstyLimiter <- t
        }
    }()

    // Maintenant on simule 5 nouvelles requêtes entrantes. Les 3 premières d'entre elles bénéficieront des possibilités de sur-traitement de `burstyLimiter`.
    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}
