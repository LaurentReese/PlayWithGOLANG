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
écrire un programme essai.go qui importe le module en question, et qui utilise une fonction du module
go mod init theModule
go build essai.go
un fichier go.sum est crée et aussi go.mod

TROP PRATIQUE !
