// Dans l'exemple précédent nous avons vu comment gérer des compteurs d'état simples avec des opérations atomiques. Pour des états plus compliqués, on peut utiliser un _[mutex](http://en.wikipedia.org/wiki/Mutual_exclusion)_ pour accéder de  manière sûre à des données à travers plusieurs goroutines.

package main

import (
    "fmt"
    "math/rand"
    "runtime"
    "sync"
    "sync/atomic"
    "time"
)

func main() {

    // Pour notre exemple, l'état `state` sera une map.
    var state = make(map[int]int)

    // Ce `mutex` va synchroniser l'accès à `state`.
    var mutex = &sync.Mutex{}

    // Pour comparer l'approche avec des mutexes avec une autre que nous verrons plus tard, `ops` va compter combien d'opérations nous réalisons avec l'état.
    var ops int64 = 0

    // Ici on lance 100 goroutines pour exécuter des lectures répétées sur l'état.
    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {

                // A chaque lecture, on sélectionne une
                // clé à laquelle on souhaite accéder,
                // on bloque le mutex avec `Lock()` pour
                // s'assurer un accès exclusif à l'état,
                // on lit la valeur de la clé choisie, on
                // débloque le mutex, puis on incrémente
                // le compteur.
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)

                // Pour nous assurer que cette goroutine
                // ne prend pas toutes les ressources, on
                // rend la main explicitement après
                // chaque opération avec
                // `runtime.Gosched()`. Le programmateur
                // gère normalement automatiquement ceci,
                // par ex. après les opérations sur les
                // canaux et pour les appels bloquants
                // comme `time.Sleep`, mais dans ce cas
                // on doit le faire manuellement.
                runtime.Gosched()
            }
        }()
    }

    // On démarre également 10 goroutines pour simuler
    // des écritures, de la même manière que pour les
    // écritures.
    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)
                runtime.Gosched()
            }
        }()
    }

    // On fait travailler les 10 goroutines sur les
    // `state` et `mutex` pendant une seconde.
    time.Sleep(time.Second)

    // On rapporte le nombre total d'opérations
    // réalisées.
    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)

    // Avec un bloquage final de `state`, on peut
    // connaitre l'état final.
    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}
