// Parser des nombres à partir d'une chaîne est une tâche simple mais courante dans de nombreux programmes. Voici comment le faire en Go.

package main

// Le package `strconv` de la librairie standard fournit ces fonctionnalités.
import "strconv"
import "fmt"

func main() {

    // Avec `ParseFloat`, le `64` indique combien de bits de précision utiliser.
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    // 
    // Pour `ParseInt`, le `0` signifie qu'on cherche à déduire la base à partir de la chaine. `64` demande à faire rentrer le résultat dans 64 bits.
    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    // `ParseInt` reconnait les nombres formattés en hexadécimal.
    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    // Une fonction `ParseUint` est également disponible.
    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

    // `Atoi` est une fonction pratique pour extraire facilement un int en base 10.
    k, _ := strconv.Atoi("135")
    fmt.Println(k)


    // Les fonctions de parsing renvoient une erreur lorsque les données d'entrée ne conviennent pas.
    _, e := strconv.Atoi("wat")
    fmt.Println(e)
}
