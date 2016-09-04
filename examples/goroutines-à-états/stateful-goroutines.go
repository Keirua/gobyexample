// Dans l'exemple précédent, nous avons utilisé un verrou explicite avec des mutexes pour synchroniser l'accès à l'état partagé à travers plusieurs goroutines. Une autre option consiste à utiliser les fonctionnalités natives de synchronisation des goroutines pour obtenir le même résultat.
// Cette approche avec les canaux s'aligne avec l'idée du Go de partager la mémoire en communiquand et en ayant chaque morceau de donnée possédé par exactement une goroutine.

package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

// Dans cet exemple, notre état sera possédé par une unique goroutine. Cela va garantir que la donnée n'est jamais corrompue par des accès concurrents. Afin de lire ou écrire dans cet état, les autres goroutines vont envoyer des messages à la goroutine propriétaire et recevoir les réponses correspondantes. Ces structures `readOp` et `writeOp` encapsulent ces requêtes et une manière pour la goroutine propriétaire de répondre.
type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}

func main() {

    // Comme précédemment, nous allons compter combien d'opérations on réalise.
    var ops int64 = 0

    // Les canaux `reads` et `writes` seront utilisés par les autres goroutines pour envoyer des demande de lecture et d'écriture.
    reads := make(chan *readOp)
    writes := make(chan *writeOp)

    // Voici la goroutine qui possède l'état `state`, qui est une map comme dans l'exemple précédent, mais qui est privée grâce à la goroutine à état. Cette goroutine fait un select de manière répétée sur les canaux `reads` et `writes`, et répond aux requêtes à mesure qu'elles arrivent. Une réponse est exécutée en réalisant tout d'abord l'opération demandée, puis en envoyant une valeur sur le canal de réponse `resp` pour indiquer la réussite (et la valeur désirée dans le cas d'une lecture).
    go func() {
        var state = make(map[int]int)
        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()

    // On démarre 100 goroutines pour faire des lectures sur la goroutine à état, grâce au canal `reads`.
    // Chaque lecture nécessite de construire un objet `readOp`, l'envoyer à travers le canal`reads`, puis recevoir le résultat à travers le canal `resp` fourni.
    for r := 0; r < 100; r++ {
        go func() {
            for {
                read := &readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int)}
                reads <- read
                <-read.resp
                atomic.AddInt64(&ops, 1)
            }
        }()
    }

    // On démarre 10 écritures également, selon la même approche.
    for w := 0; w < 10; w++ {
        go func() {
            for {
                write := &writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddInt64(&ops, 1)
            }
        }()
    }

    // On laisse les goroutines travailler pendant une seconde.
    // Let the goroutines work for a second.
    time.Sleep(time.Second)

    // Enfin, on capture et rapport le nombre d'opération.
    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)
}
