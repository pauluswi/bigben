# bigben
An e-wallet service.
Built on Go Lang.



# Installation Guide

## Docker way  (recommended)
### System requirements
- [Docker](https://www.docker.com/) - Docker is an open source platform for building, deploying, and managing containerized applications. Learn about containers, how they compare to VMs, and why Docker is so widely adopted and used.
- [docker-compose](https://docs.docker.com/compose/) - Compose is a tool for defining and running multi-container Docker applications. With Compose, you use a YAML file to configure your application’s services. Then, with a single command, you create and start all the services from your configuration

### Setup (dockerize)

note : make sure docker and docker-compose installed
- move to project dir
- run cleanup docker-compose command 
- run build command 


```sh
$ cd bigben
$ docker-compose rm
$ docker-compose up --build
```


- open another terminal
- run docker bash 
- run seeder for seeding dummy data
```sh
$ docker-compose exec app bash 
$ ./main seed CustomerSeed AccountSeed
$ go test -v controller/*.go -race -coverprofile=coverage.out -covermode=atomic
```

- after that can open the page [localhost:3000](http://127.0.0.1:3000)
- for test http request can use http-request.http file at root project
- postman collection (https://www.getpostman.com/collections/a59464ce90293b484794)

## Manual
### (manual installation)

For manual installation requires 
- MySQL Database 
- Go v1.6 above

modify .env files and configure database, 
this app provide 4 commands.

- migrating database schema
```sh
$ go run main.go migrate up 
```

- drop database schema
```sh
$ go run main.go migrate down 
```

- seed dummy data 
```sh
$ go run main.go seed [ArgSeederFunc...] 
```

- serving app

```sh
$ go run main.go serve
```

- after configuring the database, next is to run the command 

```sh
$ go mod download
$ go mod tidy
$ go run main.go migrate up 
$ go run main.go seed CustomerSeed AccountSeed
$ go run main.go serve 
```

- after that can open the page [localhost:3000](http://127.0.0.1:3000)
- for test http request can use http-request.http file at root project
- postman collection (https://www.getpostman.com/collections/a59464ce90293b484794)
- unit testing
```sh
$ go test -v controller/*.go -race -coverprofile=coverage.out -covermode=atomic
```
![](doc/Screen Shot 2021-09-10 at 13.55.27.png)

feedback
