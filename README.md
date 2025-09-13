# Golang module template

## Description

Ce projet est un template d'API REST en Go.

## Technologies utilisées

* [Go](https://golang.org/)
* [Swag](https://github.com/swaggo/swag)
* [docker](https://www.docker.com/)
* [docker-compose](https://docs.docker.com/compose/)

## Initialisation du projet

1. Clonez le dépôt :
```bash
git clone https://github.com/LordPax/golang-module-template.git
```

2. Accédez au répertoire du projet :
```bash
cd golang-module-template
```

3. Lancer les conteneur docker
```bash
docker-compose up
```

## Installation back

1. Créez un fichier `.env` à la racine du répertoire `back` et ajoutez les variables d'environnement suivantes :
```bash
NAME='Golang Api'
DOMAIN=localhost:8080
PORT=:8080
GIN_MODE=debug
ALLOWED_ORIGINS='*'

DB_HOST=localhost
DB_USER=root
DB_PASSWORD=root
DB_NAME=golang-app
DB_PORT=5432

COOKIE_SECURE=false
JWT_SECRET_KEY=secret

BREVO_API_KEY=
BREVO_SENDER=noreply@example.fr

OS_CLOUD=openstack
```

2. Intaller les dépendances :
```bash
go mod download
go mod vendor
swag init
```

3. Build le projet :
```bash
go build
```
