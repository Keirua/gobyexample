// Go supporte les _constantes_ de caractères, chaines de
// caractères, booléens et valeurs numériques

package main

import "fmt"
import "math"

// `const` declare une valeur constante.
const s string = "constant"

func main() {
    fmt.Println(s)

    // Le mot clé `const` peut apparaître à chaque
    // endroit où l'on peut mettre le mot clé `var`
    const n = 500000000

    // Les expressions constantes réalisents les
    // opérations arithmétiques avec une précision
    // arbitraire.
    const d = 3e20 / n
    fmt.Println(d)

    // Une constante numérique n'a pas de type
    // jusqu'à ce qu'on lui en donne un, par exemple
    // via un cast explicite
    fmt.Println(int64(d))

    // On peut donner un type à un nombre en l'utilisant
    // dans un contexte qui en requiert un, tel qu'un
    // assignement ou un appel de fonction.
    // Par exemple ici, `math.Sin` attends un `float64`.
    fmt.Println(math.Sin(n))
}
