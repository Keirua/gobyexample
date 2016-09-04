# Pour expérimenter avec les arguments de ligne de
# commande, il vaut mieux d'abord compiler un exécutable
# avec `go build`.
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c

# Ensuite, nous regarderons des techniques plus avancées
# de traitement avec les flags.
