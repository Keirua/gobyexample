// Les _slices_ sont un type de donnée clé en Go.
// Ils sont bien plus puissant que les array pour
// manipuler les séquences de données

package main

import "fmt"

func main() {
    // Contrairement aux array, les slices sont typées
    // uniquement par les éléments qu'elles contiennent
    // (pas par le nombre d'éléments). Pour créer une
    // slice vide de longueur non nulle, on utilise la
    // builtin `make`. Ici, on crée une slice de
    // `string`s
    // de longueur `3` (initiallement de valeurs
    // nulles).
    s := make([]string, 3)
    fmt.Println("emp:", s)

    // On peut affecter et retrouver les valeurs comme
    // avec les array
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // `len` renvoie la longueur de la slice comme
    // prévu.
    fmt.Println("len:", len(s))

    // En plus de ces opérations basiques, les slices en
    // supportent plusieurs qui les rendent plus riches
    // que les array. L'une d'elle est la builtin
    // `append`, qui renvoie une slice avec une ou
    // plusieurs nouvelles valeurs. A noter que la
    //  nouvelle slice est renvoyée par append : la
    // slice
    //  originale, passée en paramètres, n'est pas
    // modifiée.
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // On peut également également copier les slices.
    // Ici, on crée une slice `c` vide, de la même
    // longueur que `s` et on copie les valeurs de `s`
    // dans `c`
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // Les slices possède un opérateur "slice" avec la
    // syntaxe `slice[debut:fin]`. Par exemple,
    // l'exemple
    // suivant crée une slice avec les éléments `s[2]`,
    // `s[3]`, et `s[4]`.
    l := s[2:5]
    fmt.Println("sl1:", l)

    // Ceci crée une slice jusqu'à (mais sans) `s[5]`.
    l = s[:5]
    fmt.Println("sl2:", l)

    // Et ceci crée une slice à partir de (et incluant)
    // `s[2]`.
    l = s[2:]
    fmt.Println("sl3:", l)

    // On peut déclarer et initialiser une varibale pour
    // une slice en une ligne également. A la différence
    // des array, on ne précise pas la longueur dans [].
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // Les slices peuvent être composées en structures
    // de
    // données multidimensionnelles. La longueur des
    // slices internes peut varier, contrairement aux
    // tableaux multidimensionnels.
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
