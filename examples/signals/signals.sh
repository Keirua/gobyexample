# Quand on lance ce programme, il va bloquer, en attendant un signal. En faisant `ctrl-C` (que le terminal montre comme `^C`), on peut envoyer un signal `SIGINT`, qui amène le programme à afficher `interrupt` et puis termine.
$ go run signals.go
awaiting signal
^C
interrupt
exiting
