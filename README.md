# go-api-task

`Install dep`

```sh
$ go get -v github.com/golang/dep/cmd/dep
```

`Install the project's dependencies`

```sh
$ dep ensure
```

`Build API`

```sh
$ go build -o build/api
```

`Start API`

```sh
$ build/api
```

`Build docker`

```sh
docker build --no-cache -t img-task-go .
```