#TP GOLANG

## SUJET :
Réaliser un programme gérant l’état d’un jeu. Est attendu:

- Une API permettant de récupérer et modifier l’état du jeu, avec au minimum:
- GET pour récupérer l’état du jeu
- POST pour modifier l’état du jeu
- Une gestion non bloquante de l’état, de préférence avec les channels.
- Des tests unitaires pertinents.
- Un script pour tester les modifications concurrentes de l’état du jeu.


##Liens utiles : 
- [Cours du prof](https://docs.google.com/presentation/d/1GBGO9I7otJOt3QNWKtF9RKVE6zefflgh8P_dnK9Em1Q/edit#slide=id.gc6f919934_0_0)
- [Cours sur différents sujets autour du language](https://learnxinyminutes.com/)
- [Cours interatif sur les bases](https://tour.golang.org/welcome/1)
- [Liste d'example](https://gobyexample.com/)
- [Documentation](https://golang.org/)

## Commandes
- go run *.go
- ab -n 1000 -c 40 -T application/json -p data.json http://localhost:8080/api/article/1