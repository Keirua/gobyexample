# `zeroval` ne change pas le  `i` dans `main`, mais
# `zeroptr` le fait car elle a une référence vers 
# l'adresse mémoire de cette variable.
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
