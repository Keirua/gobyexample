# Lancer notre programme montre que la gestion d'état par
# goroutines réalise environ 800,000 operations par
# secondes.
$ go run stateful-goroutines.go
ops: 807434

# Pour ce cas particulier, l'approche avec les goroutines
# était un peu plus complexe qu'avec les mutexes.
# Elle peut être utile dans certains cas néanmoins, par
# exemple lorsque vous avez d'autres canaux impliqués ou 
# quand la gestion de tels mutexes serait source
# d'erreurs.
# Vous devriez utiliser l'approche qui vous parait la
# plus naturelle, et qui vous permet d'écrire le
# programme le plus juste possible.
