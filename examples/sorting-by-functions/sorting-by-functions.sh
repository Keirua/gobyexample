# Faire tourner notre programme montre une liste triée 
# par longueur de chaîne, comme désiré.
$ go run sorting-by-functions.go 
[kiwi peach banana]

# En suivant ce modèle, qui consiste à créer un type 
# custom, implémenter les trois méthodes de l'interface 
# sur ce type, et ensuite appeler sort.Sort sur une
# collection de ce type custom, nous pouvons trier des
# slices selon n'importe quelle manière.
