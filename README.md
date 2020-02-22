Some Go Tutorials 

# Getting Started
This code is my playground based on examples at https://tour.golang.org/list

To run the `hello` entrypoint:
#### Simple way
`go run hello.go`

#### As PATH executable
Follow below steps:
* set GOPATH and place on PATH

```
export GOPATH=/home/alex/Code/go
export PATH=$PATH:$GOPATH/bin
```

* run `go install hello.go`
* run entrypoint via `hello`

# Using modules and packages
Source is https://golang.org/doc/code.html

# Key concepts

## Env vars
```go
// Set GOPATH
export GOPATH=$HOME/go
export PATH="$PATH:${GOPATH//://bin:}/bin"
```

## Modules
A repository contains one or more modules. A module is a collection of related Go packages that are released together. A Go repository typically contains only one module, located at the root of the repository. A file named go.mod there declares the module path: the import path prefix for all packages within the module. The module contains the packages in the directory containing its go.mod file as well as subdirectories of that directory, up to the next subdirectory containing another go.mod file (if any).

```go
// Build compiles the packages named by the import paths, along with their dependencies, but it does not install the results
go build hello.go
```

## Packages
Go programs are organized into packages. A package is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.

## Sublime Packages Installed
* Gofmt
* Golang Build

## Getting `Goimports` to run automatically in Sublime Text
Thank you https://www.alexedwards.net/blog/streamline-your-sublime-text-and-go-workflow!

* Install Gofmt Sublime packages
* Install Goimports from command line: `go get golang.org/x/tools/cmd/goimports`
* In Gofmt Settings-User, add:
```
{
  "cmds": [
    ["goimports"]
  ],
  "format_on_save": true
}
```
* Restart Sublime

```go
// Fetch external package via CLI
go get golang.org/x/tour

// Compile and install packages named by import paths in a go file
go install hello.go

// Run package
hello
```
