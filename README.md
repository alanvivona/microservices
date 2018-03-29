# challenge

Build:
  docker-compose up

This will build and configure 3 docker containers:
- An nginx reverse proxy accesible on localhost:8080
- A Go API accesible by the nginx proxy with the following methods:
    GET     /item
    GET     /item/{id}
    POST    /item
    DELETE  /item/{id}
- A MySql container (automatically loaded with data from a script with the last dump present on /db/current)
