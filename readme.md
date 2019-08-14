# Chosen tech stack
[![N|Solid](https://beego.me/static/img/beego_purple.png)](https://beego.me/)

I use beego, because it's simple to use so I can be more productive.

# Infrastucture requirements
* [Go](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
* [Dep](https://github.com/golang/dep) - dep is a dependency management tool for Go. It requires Go 1.9 or newer to compile.

# Directory structure
- conf. All configuration needed.
- controller. All controller will be here.
- database. All migrations will be here.
- models. All model will be here.
- public. All Image will be here.
- routers. All routes will be here.
- tests. All test files will be here.

# Setup instructions
```sh
$ git clone git@bitbucket.org:shunam/pet.git
$ cd pet
$ dep ensure
$ go get github.com/beego/bee
```

# Application deployments

Make sure you are in folder pet.

Run in development
```sh
$ bee run
```

Runing test
```sh
$ cd tests
$ go test
```

Running in prouction
```sh
$ ./pet
```

# Swagger

After running the app, go to localhost:8080/swagger