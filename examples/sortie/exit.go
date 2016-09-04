// On utilise `os.Exit` pour quitter immédiatement avec
// un status donné.

package main

import "fmt"
import "os"

func main() {

    // Les `defer`s ne seront  _pas_ exécuté lorsque l'on
    // utilise `os.Exit`, donc ce `fmt.Println` ne sera
    // jamais appelé.
    defer fmt.Println("!")

    // On quitte avec un statut 3.
    // Exit with status 3.
    os.Exit(3)
}

// Contrairement au C par exemple, Go n'utilise pas un
// entier comme valeur de retour de main pour indiquer
// le statut de sortie.
// Si vous voulez indiquer une valeur de sortie non-nulle
// il faut utiliser `os.Exit`.
