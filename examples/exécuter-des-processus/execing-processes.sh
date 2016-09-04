# Quand on lance notre programme, il est remplacé par ls.
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# Notez que Go ne propose pas la fonction Unix classique
# `fork`. En général, ce n'est pas un problème, car
# lancer des goroutines, des processus et exécuter des
# processus couvre la plupart des cas d'usage de `fork`.
