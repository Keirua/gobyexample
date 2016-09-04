// Go fournit un support complet des fonctions usuelles
// de gestion du temps et des durées. Voici quelques
// exemples.

package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    // Commençons pour récupérer l'heure actuelle
    // We'll start by getting the current time.
    now := time.Now()
    p(now)

    // On peut construire un structure `time` en fournissant année, mois, jour, etc. Les `time`s sont toujours associés à un lieu, c'est à dire  à un fuseau horaire.
    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    // On peut extraire les différentes valeurs des composantes du temps comme on s'y attend.
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    // Une méthode `Weekday` est également disponible pour connaitre le jour de la semaine (Monday-Sunday)
    p(then.Weekday())

    // Ces méthodes comparent deux instances de `time`, pour savoir si la première est arrivée respectivement avant, après ou au même moment que la seconde.
    // These methods compare two times, testing if the
    // first occurs before, after, or at the same time
    // as the second, respectively.
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    // La méthode `Sub` renvoie une `Duration`, qui représente l'intervalle écoulé entre les dates contenues dans deux isntances de `time`.
    diff := now.Sub(then)
    p(diff)

    // On peut calculer la longueur de la durée dans plusieurs unités.
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    // On peut aussi utiliser `Add` pour avancer à une date selon une certaine durée, ou reculer avec `-`.
    p(then.Add(diff))
    p(then.Add(-diff))
}
