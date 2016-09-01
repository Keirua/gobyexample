// Go supporte de base les _valeurs de retour multiples_.
// On s'en sert souvent en Go idiomatique, par exemple
// pour renvoyer le résultat et la valeur d'erreur d'une
// fonction.

package main

import "fmt"

// Le `(int, int)` dans cette signature de fonction
// montre que la fonction renvoie deux `int`s.
func vals() (int, int) {
    return 3, 7
}

// Les valeurs de retour peuvent également être nommées.
// Cela peut être utile pour aider à document le but de
// chaque valeur de retour. Ces noms peuvent être
// référencés dans le corps de la fonction, comme
// n'importe quelle autre variable. Un retour nu avec un
// simple `return` utilisera les valeurs de ces variables
// comme résultat.
func splitPrice(price float32) (dollars, cents int) {
    dollars = int(price)
    cents = int((price - float32(dollars)) * 100.0)
    return
}

// Un cas d'usage, c'est pour modifier une valeur de
// retour à l'intérieur d'une clause `defer`
func getFileSize() (file_size int64, had_error bool) {
    f, err := os.Open("/tmp/dat")
    if err != nil {
        return 0, true
    }
    defer func() {
        if err := f.Close(); err != nil {
            had_error = true
        }
    }()

    fi, err := f.Stat()
    if err != nil {
        return 0, true
    }
    return fi.Size(), false
}

func main() {

    // Ici nous utilisons les deux valeurs de retour
    // différentes de l'appel avec un _assignement
    // multiple_.
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // Si vous voulez seulement un sous-ensemble des
    // valeurs de retour, vous pouvez utiliser
    // l'identifieur blanc `_`.
    _, c := vals()
    fmt.Println(c)

    dollars, cents := splitPrice(12.42)
    fmt.Println(dollars, cents)

    file_size, had_error := getFileSize()
    fmt.Println(file_size, had_error)
}
