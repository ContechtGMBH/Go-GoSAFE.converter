**Golang REST API to convert RailML files to graphs.**

Dependencies:
```
$ go get ./...
```
*Requires a Neo4j database on :7474*

How to:

```
$ go run main.go
```
or
```
$ go build -o gosafeconverter
$ ./gosafeconverter
```

Docker:
```
$ docker build --network="host" .
```

Contact:
```
harasymczuk@contecht.eu
```
