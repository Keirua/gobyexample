# Lancer  le programme montre qu'on a exécuté environ
# 3,500,000 d'opérations sur `state`, synchronisées par un
# `mutex`
$ go run mutexes.go
ops: 3598302
state: map[1:38 4:98 2:23 3:85 0:44]

# Ensuite nous verrons comment implémenter la même
# gestion d'état avec uniquement des goroutines et
# des canaux.
