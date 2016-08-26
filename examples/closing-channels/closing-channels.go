// _Fermer_ un canal indique que nous n'enverrons plus
// de valeurs dessus. Cela peut être utile pour
// indiquer au receveur qu'on a terminé.

package main

import "fmt"

// Dans cet exemple, nous allons utilise un canal `jobs`
// pour communiquer le travail à faire de `main()` à une
// autre goroutine.
// Quand nous n'aurons plus de travail à envoyer, nous
// fermons le canal `jobs` avec `close`.
func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    // Voici la goroutine qui traite les travaux.
    // Elle reçoit de manière répétée dans `j` des
    // valeurs de `jobs`avec `j, more := <-jobs`.
    // Dans cette forme particulière de réception à deux
    // valeurs, la valeur de `more` sera `false`
    // si `jobs` a été fermée et que toutes les valeurs
    // ont déjà été reçues.
    // Nous utilisons cela pour notifier le canal `done`
    // lorsque nous avons traité tous les jobs.
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    // Cela envoit 3 jobs à travers le canal `jobs`, puis
    // le ferme.
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    // On attend le worker en utilisation l'approche de
    // [synchronisation](channel-synchronization) vue plus tôt.
    <-done
}
