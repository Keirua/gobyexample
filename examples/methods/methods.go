// Go supporte des _méthodes_ définies dans 
// les structures.

package main

import "fmt"

type rect struct {
    width, height int
}

// Cette méthode `area` a un _type receveur_ `*rect`.
func (r *rect) area() int {
    return r.width * r.height
}

// Les méthodes peuvent être définies soit pour un 
// pointeur, soit pour une valeur de type receveur. Voici
// un exemple avec une valeur.
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    // Ici on appelle les deux méthodes définies dans
    // notre structure.
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    // Go gère automatiquement la conversion entre 
    // valeur et pointeur pour les appels de méthode. On
    // peut vouloir utiliser un type receveur pointeur 
    // pour éviter la copie lors des appels de méthode ou
    // pour modifier l'objet fourni.
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
