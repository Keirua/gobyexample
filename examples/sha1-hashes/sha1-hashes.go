// Les [_empreintes SHA1_](http://fr.wikipedia.org/wiki/SHA-1) sont
// fréquemment utilisées pour calculer de courtes chaines
// permettant d'identifier des chaines ou des gros objets
// binaires. Par exemple, le [système de gestion de
// version git system](http://git-scm.com/) utilise SHA1 pour
// identifier les fichiers et répertoires versionnés.
// Voici comment calculer des empreintes SHA1 en Go.

package main

// Go implémente plusieurs fonctions de hachage dans
// plusieurs packages `crypto/*`.
import "crypto/sha1"
import "fmt"

func main() {
    s := "sha1 this string"

    // Le modèle pour générer une empreinte, c'est
    // `sha1.New()`, `sha1.Write(bytes)`, puis
    // `sha1.Sum([]byte{})`.
    // Ici, on commence avec un nouveau hachage.
    h := sha1.New()

    // `Write` attend un tableau d'octets (bytes). Si
    // vous avez une chaîne, utilisez `[]byte(s)`
    // pour forcer le type.
    h.Write([]byte(s))

    //  Ceci récupère le résultat du hachage dans une
    // slice d'octets. L'argument à `Sum` peut être
    // ajouté pour faire l'ajouter à une slice d'octets
    // existante. En général, ce n'est pas nécessaire.
    bs := h.Sum(nil)

    // On affiche souvent les valeurs SHA-1 en
    // hexadécimal, par exemple pour les commits Git.
    // Utilisez le formattage avec `%x` pour convertir
    // un résultat de hachage sous forme de chaîne
    // héxadécimale.
    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}
