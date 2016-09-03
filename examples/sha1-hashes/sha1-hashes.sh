# Lancer le programme calcule l'empreinte et l'affiche
# de manière lisible en hexadécimal.
$ go run sha1-hashes.go
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94

# Vous pouvez calculer des empreintes de la même manière
# qu'ici avec d'autres méthodes de hachage. Par exemple,
# pour utiliser MD5, il faut importer `crypto/md5` et
# utiliser `md5.New()`.

# Si vous avez besoin d'empreintes sûre pour de la
# cryptographie, vous devriez vous renseigner au sujet de
# la [force des différentes algorithmes](https://fr.wikipedia.org/wiki/Fonction_de_hachage_cryptographique)!
