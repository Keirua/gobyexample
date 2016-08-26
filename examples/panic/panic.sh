# Exécuter ce programme va amener à une panique, qui 
# affiche un message d'erreur et une trace d'exécution, 
# et qui le programme en renvoyant une valeur de status
# non nulle.
$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# A noter que contrairement à certains langages qui 
# utilisent des exceptions pour la gestion de nombreuses
# erreurs, en Go il est idiomatique d'utiliser des
# valeurs de retour qui indique le status d'erreur dès 
# que possible.
