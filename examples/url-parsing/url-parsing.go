
// Les URLs fournissent une [manière uniforme de localiser des ressources](http://adam.heroku.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// Voici comment analyser les URLs en Go.

package main

import "fmt"
import "net"
import "net/url"

func main() {
    // Nous allons analyser cette URL d'exemple, qui 
    // inclue un protocole (scheme en anglais), des
    // informations d'authentification, un hôte, un port,
    // un chemin, des paramètres de requêtes et un
    // fragment de requête
    s := "postgres://user:pass@host.com:5432/path?k=v#f"

    // On parse l'URL et on s'assure qu'il n'y a pas d'erreurs.
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    // On accède facilement au protocole.
    fmt.Println(u.Scheme)


    // `User` contient toutes les informations d'authentificiation; on appele `Username` et `Password` dessus pour les valeurs individuelles.
    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

    // `Host`contient à la fois  le nom d'hôte et le port s'il est présent. Il faut utiliser `SplitHostPort` pour les extraire.
    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

    // Ici on extraie le chemin (`path`) et le fragment (ce qui est après le `#`).
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    // Pour obtenir les paramètres de requête dans une chaine au format `k=v`, on utilise  `RawQuery`. On peut aussi extraire les paramètres de requête dans une `map`. 
    // Les maps de paramètres de requêtes ont comme clé des string et comme valeur des slices de string, donc il faut accéder à l'index 0 si on veut seulement la première valeur.
    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}
