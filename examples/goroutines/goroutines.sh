# Lorsque l'on lance ce programme, on voit la sortie 
# de la fonction bloquante en premier, puis les 
# exécutions croisées des deux goroutines. 
# Cet entrelacement reflète le fait que les deux
# goroutines sont exécutées de manière concurrente.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
<enter>
done

# Ensuite nous allons regarder un complément aux 
# goroutines : les channels