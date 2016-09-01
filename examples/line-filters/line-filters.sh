# Pour essayer notre filtre de ligne, commencez par créer
# un fichier avec quelques lignes en minuscule.
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# Puis utilisez le filtre de ligne pour les mettre en
# majuscule.
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER
