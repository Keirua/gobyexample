// Go propose un support natif pour l'encodage et le
// décodage en [base64](http://fr.wikipedia.org/wiki/Base64).

package main

// Cette syntaxe importe le package `encoding/base64`
// avec le nom `b64` au lieu du nomp par défaut `base64`.
// Cela nous fait gagner un peu de place en dessous.
import b64 "encoding/base64"
import "fmt"

func main() {

    // Voici la chaîne que l'on va encoder et décoder.
    data := "abc123!?$*&()'-=@~"

    // Go support à la fois le base64 standard et celui
    // compatible avec les URLs.
    // Voici comment encoder avec l'encodeur standard. Il
    // a besoin d'un `[]byte` donc on va faire un cast de
    // notre `string` vers ce type.
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    // Le décodage peut renvoyer une erreur, que l'on
    // peut  vérifier si vous ne savez pas déjà si 
    // l'entrée est correctement formée.
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    // Ceci encode et décode selone le format base64
    // compatible avec les URLs.
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
