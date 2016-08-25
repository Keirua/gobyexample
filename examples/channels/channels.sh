# Lorsque l'on lance le programme, le message `"ping"` 
# est correctement passé d'une goroutine à l'autre via
# notre canal.
$ go run channels.go 
ping

# Par défaut, l'envoi et la réception de messages bloque
# jusqu'à ce que l'envoyeur et le receveur soient prêts.
# Cette propriété nous a permis d'attendre le message 
# `"ping"` à la fin de notre programme, sans utiliser
# une autre forme de synchronisation.
