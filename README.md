# challenge

Build:
  docker-compose up

This will build and configure 3 docker containers:
- An nginx reverse proxy accesible on localhost:8080
- A Go API accesible trough the nginx reverse proxy with the following methods:
    GET     /item
    GET     /item/{id}
    POST    /item
    DELETE  /item/{id}
    GET     /gdrive/auth
		GET     /gdrive/search-in-doc/{id}
		POST    /gdrive/file
- A MySql container with data populated using the last dump present on /db/current
