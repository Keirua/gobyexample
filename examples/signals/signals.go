// Parfois, on voudrait que les programmes en Go gèrent intelligemment les [signaux Unix](http://en.wikipedia.org/wiki/Unix_signal).
// Par exemple, on peut vouloir qu'un serveur s'éteigne gracieusement lorsqu'il reçoit le signal `SIGTERM`, ou un qu'outil en ligne de commande arrête de traiter les entrées lorsuq'il reçoit un `SIGINT`.
// Voici comment gérer les signaux en Go avec des canaux.

package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {

    // Go signale les travaux de notification en envoyer des valeurs `os.Signal` sur un canal. Nous allons créer un canal pour recevoir ces notifications. Nous allons aussi en faire un pour nous notifier quand le programme peut se terminer.
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    // `signal.Notify` enregistre le canal fournit pour recevoir les notifications sur les signaux précisés.
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    // Cette goroutine exécute une réception bloquante des signaux. Quand elle se termine, elle l'affichera et notifiera le programme qu'il peut terminer.
    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    // Le programme va attendre ici jusqu'à ce qu'il reçoive le signal attendu (comme indiqué par la goroutine au dessus, qui envoit une valeur sur `done`), puis termine.
    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
