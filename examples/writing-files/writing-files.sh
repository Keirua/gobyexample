# Essayez d'exécuter le code d'écriture dans les fichiers
# Try running the file-writing code.
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# Puis vérifiez le contenu des fichiers écrits.
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# Ensuite, nous verrons comment appliquer certains des 
# idées que nous avons vu pour l'I/O sur les fichiers
# au flux `stdin` et `stdout`.
