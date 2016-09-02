// Les [_hashes SHA1_](http://fr.wikipedia.org/wiki/SHA-1) sont fréquemment utimlisés pour calculer de courtes chaines permettant d'identifier des chaines ou des gros objets binaires. Par exemple, le [système de gestion de version git system](http://git-scm.com/) utilise SHA1 pour identifier les fichiers et répertoires versionnés. Voici comment calculer des hashes SHA1 en Go
uses SHA1s extensively to
// identify versioned files and directories. Here's how to
// compute SHA1 hashes in Go.

package main

// Go implements several hash functions in various
// `crypto/*` packages.
import "crypto/sha1"
import "fmt"

func main() {
    s := "sha1 this string"

    // The pattern for generating a hash is `sha1.New()`,
    // `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
    // Here we start with a new hash.
    h := sha1.New()

    // `Write` expects bytes. If you have a string `s`,
    // use `[]byte(s)` to coerce it to bytes.
    h.Write([]byte(s))

    // This gets the finalized hash result as a byte
    // slice. The argument to `Sum` can be used to append
    // to an existing byte slice: it usually isn't needed.
    bs := h.Sum(nil)

    // SHA1 values are often printed in hex, for example
    // in git commits. Use the `%x` format verb to convert
    // a hash results to a hex string.
    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}
