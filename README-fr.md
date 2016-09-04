## Le Go par l'exemple

Contenu et outils de génération du site [le Go par l'exemple](http://le-go-par-l-exemple.keiruaprod.fr), un site qui enseigne le Go à travers des programmes d'exemple annotés. Il s'agit de la traduction française du site [Go by Example](https://gobyexample.com)


### Présentation

Le site Le Go par l'exemple est construit en extrayant le code et les exemples des fichiers du répertoire exemples. Des pages statiques sont générées dans le répertoire `public` à partir des modèles du répertoire `templates`. Les programmes qui implémentent ce processus de génération sont dans `tools`, accompagné des dépendances nécessaires dans le répertoire `vendor`.

Le répertoire `public` généré peut être servi par n'importe quel système gérant un contenu statique. Le site original utilise Amazon S3 et CloudFront.


### Génération

Pour générer le site, Go et Python doivent être installés. Lancez :

```console
$ go get github.com/russross/blackfriday
$ tools/build
$ open public/index.html
```

Pour compiler continuellement dans une boucle :

```console
$ tools/build-loop
```


### License

Cette traduction est la propriété de Clément Camin, publié sous license 
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

Le travail original est la propriété de Mark McGranaghan, publié sous license 
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

Go Gopher est la propriété de [Renée French](http://reneefrench.blogspot.com/) publié sous licence
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).


### Traductions

Des traductions réalisées par des contributeurs du site Go by Example sont disponibles en

* [Chinois](http://gobyexample.everyx.in/) par [everyx](https://github.com/everyx)
* [Espagnol](http://goconejemplos.com) par la [communauté Go de Mexico](https://github.com/dabit/gobyexample)
* [Italien](http://gobyexample.it) par la [communauté italienne Golang](https://github.com/golangit/gobyexample-it)

### Remerciements

Merci à [Jeremy Ashkenas](https://github.com/jashkenas) pour [Docco](http://jashkenas.github.com/docco/), qui a inspiré ce projet.
