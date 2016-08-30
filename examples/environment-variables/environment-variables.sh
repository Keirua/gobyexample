# Exécuter le programme montre que nous choisissons une valeur pour FOO, mais que `BAR` est vide.
$ go run environment-variables.go
FOO: 1
BAR: 

# La liste des clés de l'environnement dépendra de votre machine.
TERM_PROGRAM
PATH
SHELL
...

# So nous affectons `BAR` dans l'environnement avant de lancer le programme, il récupérera sa valeur.
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
