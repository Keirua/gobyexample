# Pour expérimenter le programme de test d'options de ligne de commande, le plus simple est de le compiler puis d'exécuter le binaire ensuite.
$ go build command-line-flags.go

# Essayer le programme en donnant tout d'abord des valeurs pour toutes les options.
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# Notez que si vous ne précisez pas certaines options, elles prennent automatiquement leurs valeurs par défaut.
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# Des arguments restants peuvent être fournis après toutes les options
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# Notez que le package `flag` a besoin que toutes les options apparaissent avant les arguments 'positionnels'. Sinon, les options sont traités commen des arguments positionnels.
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# Utiliseez les options `-h` ou `--help` pour obtenir une aide générée automatiquement au sujet des paramètres de ligne de commande.
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# Si vous fournissez une option qui n'était pas précisée au package `flag`, le programme affichera un message d'erreur puis le texte d'aide à nouveau.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...

# Ensuite nous allons regarder les variables d'environnement, une autre manière courante de paramètrer les programmes.
