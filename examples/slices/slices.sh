# A noter que, bien que les slices soient un type
# différent des array, elles s'affichent de la même
# manière avec `fmt.Println`.
$ go run slices.go
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]

# Vous pouvez lire ce [bon article](http://blog.golang.org/2011/01/go-slices-usage-and-internals.html)
# de l'équipe Go pour plus de détails sur le design et
# l'implémentation des slices en Go.

# Maintenant que nous avons vu les arrays et les slices,
# nous allons regarder une autre structure de donnée
# clé : les maps
