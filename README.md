# Go Gin Boilerplate
> A starter project with Golang, Gin and DynamoDB
> From [credit-url]



Golang Gin boilerplate with DynamoDB resource. Supports multiple configuration environments.

![](header.jpg)

This project use a [DynamoDB base docker image](https://github.com/vsouza/docker-dynamoDB-local).

Setup DynamoDB dependency:  `docker run -d -p 8080:8080 vsouza/dynamo-local --port 8080`

### Boilerplate structure

```
.
├── Makefile
├── Procfile
├── README.md
├── config
│   ├── config.go
│   ├── development.yaml
│   ├── production.yaml
│   └── test.yaml
├── controllers
│   └── user.go
├── db
│   └── db.go
├── forms
│   └── user.go
├── header.jpg
├── main.go
├── middlewares
│   └── auth.go
├── models
│   └── user.go
└── server
    ├── router.go
    └── server.go
```

## Installation

```sh
make deps
```

## Usage example

`curl http://localhost:8888/health`

## Development setup

Running DynamoDB on Docker Image:

check this project: [vsouza/docker-dynamoDB-local](https://github.com/vsouza/docker-dynamoDB-local)

## Release History

* 0.0.1
    * Configuration by environment, Auth and Log middlewares, User entity.

## Meta

Vinicius Souza – [@iamvsouza](https://twitter.com/iamvsouza) – hi@vsouza.com

Distributed under the MIT license. See [License](https://vsouza.mit-license.org) for more information.

[https://github.com/vsouza](https://github.com/vsouza)


[credit-url]: https://github.com/vsouza/go-gin-boilerplate


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fvsouza%2Fgo-gin-boilerplate.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fvsouza%2Fgo-gin-boilerplate?ref=badge_large)
