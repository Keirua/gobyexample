// On veut souvent exécuter du code Go à un certain
// point du futur, ou de manière répétée selon un
// intervalle de temps. Les fonctionnalités builtin
// _timer_ et _ticker_ rendent ces deux tâches faciles.
// Nous allons d'abord regarder les timers, et ensuite
// les [tickers](tickers).

package main

import "time"
import "fmt"

func main() {

    // Les timers représente un événement unique dans le
    // futur. On précise combien de temps on veut
    // attendre, et il founit un canal qui sera notifié
    // à ce moment là. Ce timer attendra 2 secondes.
    timer1 := time.NewTimer(time.Second * 2)

    // Le `<-timer1.C` bloque sur le canal `C` du timer
    // jusqu'à ce qu'il envoie une valeur indiquant que
    // le timer a expiré.
    <-timer1.C
    fmt.Println("Timer 1 expired")

    // Pour simplement attendre, on peut utiliser
    // `time.Sleep`. Mais les timers peuvent être utiles
    // car on peut les stopper avant qu'ils expirent.
    // Voici un exemple de ceci.
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}
