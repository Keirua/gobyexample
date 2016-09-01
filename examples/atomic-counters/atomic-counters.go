// Le mécanisme principal pour gérer les états en Go,
// c'est la communication à travers les canaux. Nous
// avons vu ceci par exemple avec les [worker pools](worker-pools).
// Il y a quelques autres options pour gérer les états
// ceci dit. Ici, nous allons utiliser le package
// `sync/atomic` pour des _compteurs atomiques_ auxquels
// accèdent plusieurs goroutines.

package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

func main() {

    // Nous utiliserons un entier non-signé pour
    // représenter notre compteur (toujours positif).
    var ops uint64 = 0

    // Pour simuler des mises à jour concurrentes, nous
    // commençons 50 goroutines qui incrémentent le
    // compteur environ une fois par milliseconde.
    for i := 0; i < 50; i++ {
        go func() {
            for {
                // Pour incrémenter atomiquement le
                // compteur, nous utilisons `AddUint64`,
                // en luis donnant l'adresse mémoire de
                // notre compteur `ops` avec la syntaxe
                // `&`.
                atomic.AddUint64(&ops, 1)

                // On fait continuer la goroutine.
                runtime.Gosched()
            }
        }()
    }

    // On attend une seconde pour permettre aux
    // goroutines de travailler, et au compteur de
    // grandir en valeur.
    time.Sleep(time.Second)

    // Afin d'utiliser le compteur de manière sûre alors
    // qu'il est toujours mis à jour par les autres
    // goroutines, on extraie une copie de la valeur
    // courante dans `opsFinal` via `LoadUint64`.
    // Comme plus haut, nous devons donner à cette
    // fonction l'adresse mémoire `&ops` depuis laquelle
    // chercher la valeur.
    opsFinal := atomic.LoadUint64(&ops)
    fmt.Println("ops:", opsFinal)
}
