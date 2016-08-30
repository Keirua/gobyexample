// Go supporte le formattage et le parsing de temps via des motifs.

package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    // Voici un exemple basic de formattage de temps selon la  RFC3339, qui utilise la constante correspondante.
    t := time.Now()
    p(t.Format(time.RFC3339))

    // On parse le temps en utilisant la même constante dans `Format`.
    t1, e := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00")
    p(t1)

    // `Format` et `Parse` utilise des motifs d'exemples. En général, vous utiliserez une constante de time pour celà, mais vous pouvez aussi fournir votre propre formattage. Le formattage doit utiliser le temps de référence `Mon Jan 2 15:04:05 MST 2006` pour montrer le motif avec lequel formatter/parser un temps ou une chaîne. 
    // Le temps d'exemples doit être exactement comme ceci : l'année 2006, 15 pour l'heure, lundi pour le jour de la semaine, etc.
    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    p(t2)


    // Pour des représentations purement numériques, vous pouvez également utiliser des fonctions standard de formattage en extrayant les composantes du temps nécessaires.
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

    // `Parse` renvoie une erreur qui explique le problème lorsque le format d'entrée est mal formé will return an error on malformed input.
    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)
}
