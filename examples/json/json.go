// Go fournit un support pour l'encodage et le décodage en JSON, à la fois pour les types intégrés mais aussi pour les types sur mesure.

package main

import "encoding/json"
import "fmt"
import "os"

// Nous allons utiliser ces deux structures pour démontrer l'encodage et le décodage vers des types de données custom plus bas.
type Response1 struct {
    Page   int
    Fruits []string
}
type Response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {

    // Nous allons commencer par encoder des types de données de base vers une chaîne JSON. Voici quelques exemples pour des valeurs atomiques.
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

    // Et voici quelques exemples pour les slices et les maps, qui encodent vers des tableaux et objets JSON comme attendu.
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    // Le package JSON peut automatiquement encoder vos types de données custom. 
    // Il incluera seulement les champs exportés dans la sortie encodée et utilisera par défaut ces noms comme clés JSON.
    res1D := &Response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    // Vous pouvez utiliser des tags sur les déclarations de champs de structures pour personnaliser l'encodage des noms de clé JSON. Regardez le définition de `Response2` au-dessus pour voir un exemple de tels tags.
    res2D := &Response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

    // Maintenant regardons le décodage des données JSON vers des valeurs. Voici un exemple pour une structure de donnée générique.
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

    // Nous devons fournir une variable dans laquelle le package JSON peut décoder les données. Cette `map[string]interface{}` contiendra une map avec chaînes comme clé et comme valeur des types arbitraires.
    var dat map[string]interface{}

    // Voici le décodage, accompagné d'un test d'erreurs
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    // Pour utiliser les données dans la map décodée, nous devons les convertir dans le type approprié. Par exemple ici, on convertir `num` en `float64`.
    num := dat["num"].(float64)
    fmt.Println(num)

    // Accéder aux données imbriquées nécessite une série de casts.
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

    // Nous pouvons aussi décoder du JSON dans des types de données custom. Ceci a également l'avantage d'assurer une sécurité au niveau des types de données, et d'éliminer le besoin de vérifier les types lorsqu'on accède aux données décodées.
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := Response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    // Dans les exemples ci-dessus, nous avons toujours utilisé des octets et des chaînes comme données intermédiaires entre les données et leur représentation JSON sur la sortie standard. On peut également streamer l'encodage directement dans un `os.Writer` comme `os.Stdout`, ou même dans le corps d'une réponse HTTP.
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}
