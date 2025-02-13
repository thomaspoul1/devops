### Compilation
Pour compiler l'application, utilisez la commande suivante depuis le répertoire principal :

```shell
go build -o bin/app .

## L'application requiert un fichier config.yaml dans le dossier /opt pour fonctionner correctement. Structure du fichier de configuration nécessaire :
image: "chemin vers une image en ligne"
