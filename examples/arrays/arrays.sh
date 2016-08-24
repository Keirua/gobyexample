# A noter que les tableaux apparaissent sous la forme
# `[v1 v2 v3 ...]` lorsqu'on les affiche avec 
# `fmt.Println`.
$ go run arrays.go
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d:  [[0 1 2] [1 2 3]]

# En Go, on utilise les _slices_ bien plus souvent que 
# les tableaux. C'est ce que nous allons voir maintenant.
