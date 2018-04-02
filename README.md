# Challenge

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
  -	GET     /gdrive/search-in-doc/{id}  
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

## Unit Tests
There are some unit tests for the items API on app/src/api/app/items/items_test.go

## Integration Tests
I've wrote some integration tests for the app.
You can find then on integration-tests/integration_test.go