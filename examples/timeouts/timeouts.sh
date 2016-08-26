# Lancer ce programme montre que la première opération
# timeout et que la seconde réussit.
$ go run timeouts.go 
timeout 1
result 2

# Utiliser ce modèle de timeout avec `select` nécessite 
# de communiquer les résultats à travers les canaux.
# C'est une bonne idée en général, car d'autres
# fonctionnalités importantes de Go sont basées sur les
# canaux et sur `select`.
# Nous allons regarder cela ensuite: timers et tickers.
