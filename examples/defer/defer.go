// _Defer_ est utilisé pour s'assurer qu'un appel de
// fonction est réalisé en dernier à la fin de
// l'exécution d'un programme, en général pour libérer
// les ressources.
// On utilise souvent `defer` là où on aurait utilisé
// `ensure` et `finally` dans d'autres langages.

package main

import "fmt"
import "os"

// Supposons que nous voulions créer un fichier, écrire
// dedans et le fermer quand on a terminé. Voici comment
// nous pourrions faire celà avec `defer`.
func main() {

    // Immédiatement après avoir obtenu un objet fichier
    // avec `createFile`, on repousse la fermetture de ce
    // fichier avec `closeFile`. Cette fonction sera
    // exécutée à la fin de la fonction où se trouve
    // l'appel (ici, `main`), c'est à dire après que
    // `writeFile` ait terminé.
    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
    fmt.Println("closing")
    f.Close()
}
