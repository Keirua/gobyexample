# Nous recevons les valeurs `"one"` puis `"two"`
# comme attendu.
$ time go run select.go 
received one
received two

# A noter que le temps total d'exécution est seulement
# de ~2 secondes, car les deux `Sleeps` s'exécutent de 
# manière concurrente.
real	0m2.245s
