# FS Go

FS Go is a tiny CLI tool in order to spawn a local file server.

### REQUIREMENTS

- [Go](https://golang.org) >= **1.13**

### BUILD

In order to build the executable run the following command:

```
$ go build
```

### START THE SERVER

Starting the server is easy:

```
$ ./fs-go -addr 8080 -dir .
```

Only two parameters are needed:

- **dir**: The directory that will be served. (**default**: `.`)
- **addr**: The address that the server will listen to. (**default**: `:8080`)

### INSTALL GLOBALLY

If you only want to use this CLI tool as a binary then you can do that
easily:

```
$ go install
```

Now you can use:

```
$ fs-go
```

From anywhere in your system.

### BUILD AND RUN WITH DOCKER

Build image
```
docker build -t fs:local .
```

Run it
```
docker run --rm -v $(pwd):/data -p 8080:8080 fs:local
```

Run in different port
```
docker run --rm -v $(pwd):/data -p 9000:9000 fs:local -addr :9000
```
