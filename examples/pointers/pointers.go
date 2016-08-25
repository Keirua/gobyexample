// Go supporte les [pointeurs](https://fr.wikipedia.org/wiki/Pointeur_(programmation)), 
// ce qui permet de passer des références à des valeurs
// dans les programmes

package main

import "fmt"

// Nous allons voir comment fonctionnent les pointeurs,
// par opposition aux valeurs, dans deux fonctions : 
// `zeroval` et `zeroptr`.
// `zeroval` a un paramètre, donc les arguments lui sont
// passés par valeur. `zeroval` aura une copie de  `ival`
// distincte de celle de la fonction appellante.
func zeroval(ival int) {
    ival = 0
}


// `zeroptr` a cette fois-ci un paramètre `*int`, 
// c'est-à-dire un pointeur sur un `int`. Le `*iptr` dans
// le corps de la fonction permet de _deréférencer_ le 
// pointeur de son adresse mémoire pour obtenir la valeur 
// à cette adresse. Assigner une valeur à un pointeur 
// déréférence change la valeur à l'adresse référencée.
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    // La syntaxe `&i` donne l'adresse mémoire de  `i`,
    // c'est à dire un pointeur vers `i`.
    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    // On peut afficher les pointeurs également.
    // Pointers can be printed too.
    fmt.Println("pointer:", &i)
}
