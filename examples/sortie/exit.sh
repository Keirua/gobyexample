# Si vous lancez `exit.go` en utilisant `go run`, le code
# de status sera récupéré et affiché par `go`.
$ go run exit.go
exit status 3

# En compilant et en exécutant le binaire, vous pouvez 
# voir le statut dans le terminal.
$ go build exit.go
$ ./exit
$ echo $?
3

# Notez que le `!` de notre programme n'a jamais été 
# affiché.
