# Notre programme montre, au cours de son exécution, les
# 9 jobs qui s'exécutent, traités par les différents
# workers. Le programme prend seulement environ 3
# secondes à s'exécuter, bien qu'il traite environ 9
# secondes de travail car il y a 3 workers qui
# travaillent de manière concurrente.
$ time go run worker-pools.go 
worker 1 processing job 1
worker 2 processing job 2
worker 3 processing job 3
worker 1 processing job 4
worker 2 processing job 5
worker 3 processing job 6
worker 1 processing job 7
worker 2 processing job 8
worker 3 processing job 9

real	0m3.149s
