// Les [timers](timers) servent quand on veut faire
// quelquechose dans le futur. Les _tickers_ servent à
// faire quelquechose de manière répétée à intervalles
// réguliers. Voici un exemple d'un ticker qui se répète
// jusqu'à ce qu'on l'arrête.
// intervals. Here's an example of a ticker that ticks
// periodically until we stop it.

package main

import "time"
import "fmt"

func main() {

    // Les tickers utilisent un mécanisme similaire à
    // celui des timers : un canal pour envoyer des
    // valeurs. Ici, nous utilisons la builtin `range`
    // sur le canal pour itérer sur les valeurs
    // lorsqu'elles arrivent, toutes les 500ms.
    ticker := time.NewTicker(time.Millisecond * 500)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()

    // Les tickers peuvent être arrêtés comme les timers.
    // Lorsqu'un ticker est arrêté, il ne recevra plus de
    // valeurs sur son canal. Nous arrêterons le notre
    // après 1600ms.
    time.Sleep(time.Millisecond * 1600)
    ticker.Stop()
    fmt.Println("Ticker stopped")
}
