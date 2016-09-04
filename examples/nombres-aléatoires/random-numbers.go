// Le package `math/rand` de Go permet la génération de nombres [pseudo-aléatoires](https://fr.wikipedia.org/wiki/G%C3%A9n%C3%A9rateur_de_nombres_pseudo-al%C3%A9atoires)

package main

import "time"
import "fmt"
import "math/rand"

func main() {

    // Par exemple, `rand.Intn` renvoie un `int` n, tel que
    // `0 <= n < 100`.
    fmt.Print(rand.Intn(100), ",")
    fmt.Print(rand.Intn(100))
    fmt.Println()

    // `rand.Float64` renvoie un `float64` `f`, tel que
    // `0.0 <= f < 1.0`.
    fmt.Println(rand.Float64())

    // On peut s'en servir pour générer des float aléatoires dans d'autres intervalles, par exemple
    // `5.0 <= f' < 10.0`.
    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)
    fmt.Println()

    // Le générateur de nombres par défaut est déterministe, donc il produit la même séquence de nombres à chaque fois par défaut. Pour produire des séquences variées, il faut lui fournir une `graine` qui change. A noter que ce n'est pas une technique sûre pour les nombres aléatoires qui doivent être secrets : utiliser `crypto/rand` pour ceux-là.
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    // on appelle les méthodes de l'instance de `rand.Rand` obtenue comme les autres méthodes du package.
    fmt.Print(r1.Intn(100), ",")
    fmt.Print(r1.Intn(100))
    fmt.Println()

    // Si on fournit deux fois la même graine à deux générateurs, il produiront la même séquence de nombre aléatoires..
    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.Intn(100), ",")
    fmt.Print(r2.Intn(100))
    fmt.Println()
    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    fmt.Print(r3.Intn(100), ",")
    fmt.Print(r3.Intn(100))
}
