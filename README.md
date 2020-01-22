Some Go Tutorials

# Using modules and packages
Source is https://golang.org/doc/code.html

# Key concepts
Set GOPATH
```
export GOPATH=$HOME/go
export PATH="$PATH:${GOPATH//://bin:}/bin"
```

Build module
```
go build hello.go
```

Install package
```
go install hello.go
```

Run package
```
hello
```
  