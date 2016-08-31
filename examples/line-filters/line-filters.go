// Un _filtre de ligne_ est un type de programme courant qui lit des entrées sur stdin, les traite, puis affiche un resultat dérivé sur stdout. `grep` et `sed` sont des exemples de filtres de ligne.

// Voici un filtre de ligne en Go qui met en majuscule tout le texte d'entrée. Vous pouvez utilisez ce modèle pour écrire vos filtres de lignes en Go.
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    // On encadre `os.Stdin` qui n'a pas de buffer, avec un scanner qui en a un. Il nous fournit en outre une méthode pratique `Scan` qui avance la lecture jusqu'à la marque suivant, c'est à dire ici la ligne suivante dans le scanner par défaut.
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        // `Text` renvoie la zone de texte courante, ici une ligne du texte d'entrée.
        ucl := strings.ToUpper(scanner.Text())

        // On affiche la ligne en majuscules.
        fmt.Println(ucl)
    }

    // On teste les erreurs durant `Scan`. Un caractère de fin de fichier est attendu et si on en trouve pas, il est reporté par `Scan` comme une erreur.
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
