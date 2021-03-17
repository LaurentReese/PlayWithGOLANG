Très bien expliqué ici :
https://dpp.st/blog/golang-modules/

Pour construire le module
=========================
écrire un package (différent de main) dans laurent.go, avec une fonction commençant par une majuscule
Le nom du package dans laurent.go doit être le nom du module (et du répertoire) : theModule

go mod init github.com/LaurentReese/PlayWithGOLANG/theModule

git add .

git commit

git push


Plus tard utiliser des tags pour donner des versions à ce module, genre :
(git tag v1.0.0)
(git push --tags)

Pour utiliser le module
=======================
écrire un programme essai.go qui importe le module en question, et qui utilise les fonctions du module
(si le module a été mis à jour et que je veux la dernière version
effacer go.mod, effacer go.sum, effacer essai
go get -u github.com/LaurentReese/PlayWithGOLANG/theModule
)  => le -u est très important

(Si c'est la première fois qu'on utilise le module
go mod init theModule
un fichier go.sum est crée et aussi go.mod
)
go build essai.go
