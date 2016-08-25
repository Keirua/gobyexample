// Les _interfaces_ sont des collections nommées de
// signatures de méthodes.

package main

import "fmt"
import "math"

// Voici une interface simple pour une forme géométrique.
type geometry interface {
    area() float64
    perim() float64
}

// Pour notre example, nous allons implémenter cette
// interface sur les types `rect` et `circle`.
type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

// Pour implémenter une interface en Go, il suffit
// simplement d'en implémenter toutes les méthodes.
// Ici, nous implémentons `geometry` chez `rect`s.
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    // Voici maintenant l'implémentation pour les `circle`.
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// Si une variable a une interface, alors on peut appeler
// les méthodes de l'interface. Voici une fonction
// générique qui utilise ceci pour travailler avec
// n'importe quelle `geometry`.
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    // Mes structures `circle` et `rect` implémentent
    // tout deux les interfaces de `geometry`, on peut
    // donc utiliser des instances de ces structures en
    // argument de `measure`.
    measure(r)
    measure(c)
}
