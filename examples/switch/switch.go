// Les expressions _Switch réalisent des conditions
// à travers plusieurs branches

package main

import "fmt"
import "time"

func main() {

    // Voici un `switch` basique.
    i := 2
    fmt.Print("write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    // On peut utiliser des virugules pour séparer
    // plusieurs expressions dans un même `case`. Nous
    // utilisons également le cas `default` optionnel
    // dans cet exemple.
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("it's the weekend")
    default:
        fmt.Println("it's a weekday")
    }

    // `switch` sans expression est une autre manière de
    // faire un if/else. Vous pouvez aussi voir que les
    // expressions dans `case` peuvent être
    // non-constantes.
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("it's before noon")
    default:
        fmt.Println("it's after noon")
    }

    // Un switch sur les types `switch` compare des types et non des valeurs. Vous pouvez utiliser cela pour découvrir le type correspondant à la valeur d'une interface. Dans cet exemple, la variable `t` aura comme valeur le type de la variable i.
    whatAmI := func(i interface{}) string {
        switch t := i.(type) {
        case bool:
            return "I am a bool"
        case int:
            return "I am an int"
        default:
            return fmt.Sprintf("Can't handle type %T", t)
        }
    }
    fmt.Println(whatAmI(1))
    fmt.Println(whatAmI(true))
    fmt.Println(whatAmI("hey"))
}
