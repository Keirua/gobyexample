// Go propose un excellent support du formattage façon `printf`.
// Voici quelques exemples de formattages courants.

package main

import "fmt"
import "os"

type point struct {
    x, y int
}

func main() {

    // Go propose plusieurs "verbes" d'affichage, conçus pour formatter des valeurs générales en Go. Par exemple, ceci affiche une instance de notre structure `point`.
    p := point{1, 2}
    fmt.Printf("%v\n", p)

    // Si la valeur est une `struct`, la variable `%+v` va inclure les noms des champs.
    fmt.Printf("%+v\n", p)

    // La variante `%#v` affiche une représentation en syntaxe Go de la valeur, c'est-à-dire le bout de code qui produirait cette valeur.
    fmt.Printf("%#v\n", p)

    // Pour afficher le type d'une valeur, on utilise `%T`.
    fmt.Printf("%T\n", p)

    // Formatter des booléens est simple.
    fmt.Printf("%t\n", true)

    // Il y a beaucoup d'options pour formatter des entiers.
    // `%d` est le formattage standard, en base 10.
    fmt.Printf("%d\n", 123)

    // `%b` fournit une représentation binaire.
    fmt.Printf("%b\n", 14)

    // `%c` affiche le caractère correspondant à l'entier donné.
    fmt.Printf("%c\n", 33)

    // `%x` renvoie l'encodage héxadécimal.
    fmt.Printf("%x\n", 456)

    // Il y a également plusieurs options de formattage pour les float. Pour du formattage basique décimal, utilisez `%f`.
    fmt.Printf("%f\n", 78.9)

    // `%e` et `%E` formattent les float en (des versions légèrement modifiées de) notation scientifique.
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)

    // Pour afficher des chaînes basiques utilisez `%s`.
    fmt.Printf("%s\n", "\"string\"")

    // Pour des chaines avec double-quote comme en code Go, utiliser `%q`.
    fmt.Printf("%q\n", "\"string\"")

    // Comme avec les entiers vus plus tôt `%x` affiche la chaîne en base 16, avec 2 caractères de sortie par octet d'entrée.
    fmt.Printf("%x\n", "hex this")

    // Pour afficher une représentation d'un pointeur, utiliser `%p`.
    fmt.Printf("%p\n", &p)

    // Lorsqu'on formatte des nombres, on veut souvent controller la largeur et la précision du chiffre en sortie. Pour spécifier la largeur d'un entier, on met un nombre après le `%` dans le verbe. Par défaut le résultat est justifié et complété avec des espaces.
    fmt.Printf("|%6d|%6d|\n", 12, 345)

    // On peut aussi préciser la largeur des floats, mais on peut cette fois-ci également préciser la largeur de la précision avec la syntaxe largeur.precision
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

    // Pour justifier à gauche, ajouter `-`.
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

    // Vous pouvez aussi vouloir contrôller la largeur lorsque vous formattez des chaînes, en particulier pour contrôller qu'elles s'alignent comme dans des tableaux. Voici un exemple justifié à droite :
    fmt.Printf("|%6s|%6s|\n", "foo", "b")

    // Pour justifier à gauche, on utilise `-`, comme pour les nombres.
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

    // Pour le moment nous avons vu `Printf`, qui affiche la chaîne dans `os.Stdout`. `Sprintf` formatte et renvoie la chaîne sans l'afficher.
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)

    // On peut formatter et afficher dans des `io.Writers` autres que `os.Stdout` avec `Fprintf`.
    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
