# La chaîne est encodée un peu différente avec les deux 
# encodeurs (`+` à la fin au lieu de `-`), mais elles se
# décodent vers la chaîne de départ comme désiré.
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
