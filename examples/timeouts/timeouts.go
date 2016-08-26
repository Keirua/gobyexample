// Les _timemouts_ (dépasser le temps d'exécution alloué)
// sont importants pour des programmes qui se connectent
// à des ressources externes ou qui doivent limiter leur
// temps d'exécution.
// Implémenter des timeout en Go est facile et élégant
// grâce aux canaux et à `select`.

package main

import "time"
import "fmt"

func main() {

    // Pour notre exemple, supposons que l'on exécute un
    // appel externe qui renvoie son résultat sur le
    // canal `c1` après 2s.
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()

    // Voici le `select` qui implémente le timeout.
    // `res := <-c1` attend le résultat et `<-Time.After`
    // attend qu'une valeur soit envoyée après 1s.
    // Comme `select` traite la première réception qui
    // arrive, nous prenons le cas du timeout si
    // l'opération prend plus de temps qu'il ne lui
    // en est alloué
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(time.Second * 1):
        fmt.Println("timeout 1")
    }

    // Si nous permettons un plus gros timeout de 3s,
    // alors la réception depuis `c2` va réussir et nous
    // afficherons le résultat.
    c2 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("timeout 2")
    }
}

// todo: annulation?
