// Le package `strings` de la librairie standard fournit de nombreuses fonctions utiles pour la manipulations des chaînes de caractères. En voici un aperçu.

package main

import s "strings"
import "fmt"

// On crée un alias sur `fmt.Println` vers un nom plus court car on va beaucoup s'en service ensuite.
var p = fmt.Println

func main() {

    // Voici un exemple des fonctions disponibles dans
    // `strings`. Notez que ce sont toutes des fonctions du package, pas des méthodes sur l'objet string lui-même.
    // Cela signifne que vous devez passer la chaîne en question commen premier argument des fonctions.
    p("Contains:  ", s.Contains("test", "es"))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
    p()

    // Vous pourrez trouver plus de fonctions dans la documentation du package [`strings`](http://golang.org/pkg/strings/)

    // Cela ne fait pas partie de `strings`, mais il est utile de mentionner qu'on peut obtenir la longueur d'une chaîne avec `len` et qu'on peut accéder à un caractère en particulier via son index.
    p("Len: ", len("hello"))
    p("Char:", "hello"[1])
}
