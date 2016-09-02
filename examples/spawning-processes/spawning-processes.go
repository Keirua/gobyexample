// Parfois nos programmes Go doivent lancer d'autres processus non-Go. Par exemple, la mise en forme de syntaxe de ce site est implémentée par un processus [`pygmentize`](http://pygments.org/) lancé depuis un programme Go. Regardons quelques exemples de lancer de programmes depuis Go.

package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {

    // Nous allons commencer par une simple commande qui ne prend ni arguments ni entrée, et qui affiche simplement quelque chose sur la sortie standard. La fonction helper `exec.Command` crée un objet pour représenter ce processus externe.
    dateCmd := exec.Command("date")

    //
    // `.Output` est un autre helper qui gère les cas courants d'exécution de commande, en attendant qu'elle se termine et en récupérant sa sortie. S'il n'y a pas eu d'erreurs, `dateOut` contiendra les octets avec les informations de date.
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))

    // Ensuite nous allons regarder un exemple un peu plus complexe, où on transfère des données au processus externe sur son entrée standard `stdin`, et collecte les résultats de sa sortie `stdout`.
    grepCmd := exec.Command("grep", "hello")

    // Ici, on récupère explicitement les canaux d'entrée/sortie, démarre le processus, envoie des données d'entrée, lit la sortie correspondante, et finalement attend que le processus se termine.
    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd.Wait()

    // Nous n'avons pas fait les contrôles d'erreurs dans les exemple ci-dessus, mais vous pouvez utiliser l'habituel motif `if err != nil` pour chacun d'entre eux. De plus, on ne collectre les résultats qu'à travers `StdoutPipe`, mais on pourrait utiliser la sortie d'erreur (`stderr`) de la même manière via `StderrPipe`.
    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

    // A noter que quand on lance des commandes, nous devons fournir un commande et un tableau d'arguments explicitement délimitée (et non une unique chaîne comme sur la ligne de commande).
    // Si vous voulez lancer une commande avec une chaîne, vous pouvez utiliser l'option -c de `bash` :
    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}
