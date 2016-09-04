// En Go, il est idiomatique de communiquer les erreurs
// via une valeur de retour explicite et séparrée. Cela
// contraste avec les exceptions utilisées dans des
// languages comme Java ou Ruby, et ce qu'on retrouve
// en C, où l'on renvoit parfait une valeur, et parfois
// un code d'erreur. L'approche en Go fait qu'il est
// facile de voir quelles fonctions renvoient un code
// d'erreur et de les gérer en utilisant les mêmes
// constructions du langage utilisées pour n'importe
// quelle autre tâche sans erreur.

package main

import "errors"
import "fmt"

// Par convention, les erreurs sont la dernière valeur de
// retour et ont pour type `error`, une interface du
// langage.
func f1(arg int) (int, error) {
    if arg == 42 {

        // `errors.New` construit une valeur d'erreur
        // basique avec le message d'erreur donné.
        return -1, errors.New("can't work with 42")

    }

    // La valeur nil pour l'erreur indique qu'il n'y a
    // pas eu d'erreurs.
    return arg + 3, nil
}

// Il est possible d'utiliser des types d'erreurs sur
// mesure en implémentant la méthode  `Error()` sur ces
// types. Voici une variante de l'exemple ci dessus qui
// utilise un type sur mesure pour représenter
// explicitement une erreur d'argument.
type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {
        // Dans ce cas, nous utilisons la syntaxe
        // `&argError` pour construire l'object et
        // fournir les valeurs des deux champs `arg`
        // et `prob`.
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {
    // Les deux boucles ci-dessous testent chacunes de
    // nos fonctions qui renvoient des erreurs. Notez que
    // l'utilisation d'un test d'erreur inline dans le
    // `if` est idiomatique en code Go.
    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    // Si vous voulez utilisez les données d'une erreur
    // custom, vous devrez récupérer l'erreur comme une
    // instance du type d'erreur sur mesure en
    // explicitant le type à utiliser.
    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
