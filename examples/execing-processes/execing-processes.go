// Dans l'exemple précédent, nous avions regardé comment [lancer des processus externes](spawning-processes). On fait ça quand on a besoin d'un processus externe accessible à un processus Go qui tourne. Parfois, on veut complètement remplacer le processus Go en cours avec un autre (peut-être non-Go). Pour cela, nous allos utiliser l'implémenter Go de la fonction classique [exec](https://en.wikipedia.org/wiki/Exec_(system_call)).

package main

import "syscall"
import "os"
import "os/exec"

func main() {

    // Pour cet exemple, nous allons lancer `ls`. Go a besoin d'un chemin absolu vers l'exécutable que l'on veut lancer, donc on va utiliser `exec.LookPath` pour le localiser (il est probablement situé dans `/bin/ls`).
    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

    // `Exec` a besoin d'une slice pour les arguments (et non une grosse chaîne).
    // On va donner à `ls` quelques arguments courants.
    // A noter que le premier arguments doit être le nom du programme.
    args := []string{"ls", "-a", "-l", "-h"}

    // `Exec` a également besoin d'un ensemble de [variables d'environnement](environment-variables). Ici, on va lui fournir notre environnement actuel.
    env := os.Environ()

    // Voici enfin l'appel à `syscall.Exec`. Si l'appel réussit, l'exécution de notre processus va s'arrêter et sera remplacé par celui de `/bin/ls -a -l -h`
    // S'il y a une erreur, nous aurons une valeur de retour.
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}
