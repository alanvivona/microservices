# Challenge

![Arch diagram](https://raw.githubusercontent.com/alanvivona/microservices-arch/master/docs/microservices-arch.svg?sanitize=true)

## Build :  
  - docker-compose up

This will build and configure 3 docker containers within the same network:
- An nginx reverse proxy accesible on localhost:8080
- A Go API accesible trough the nginx reverse proxy with the following methods:  
  - GET     /item  
  - GET     /item/{id}  
  - POST    /item  
  - DELETE  /item/{id}  
  - GET     /gdrive/auth  
  -	GET     /gdrive/search-in-doc/{id}?word={word}
  - POST    /gdrive/file  
- A MySql container <b>populated with data from a dump on /db/current</b>

## Example API calls

  Get all items:  
  > curl 'http://localhost:8080/item/' -X GET

  Get an item:  
  > curl 'http://localhost:8080/item/1' -X GET

  Create item:  
  > curl 'http://localhost:8080/item/' -X POST -d '{"name":"new item name","description":"new item description"}'  

  Delete item:  
  > curl 'http://localhost:8080/item/1' -X DELETE

  Auth with google (the other two gdrive APIs will redirect here if not authorized yet)
  > curl 'http://localhost:8080/gdrive/auth' -X GET

  Search in doc
  > curl 'http://localhost:8080/gdrive/search-in-doc/1?word=dev' -X GET

  Create a gdrive file
  > curl 'http://localhost:8080/gdrive/file' -X POST -d '{"name": "test gdrive file name", "description": "test gdrive file description"}'

## Unit Tests  
There are some unit tests for the items API on app/src/api/app/items/items_test.go

## Integration Tests  
Integration tests for the app can be found on integration-tests/integration_test.go
