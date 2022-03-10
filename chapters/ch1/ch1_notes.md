# Chapter 1: Setting Up Your Go Environment

#### 1.0 Running a go script:

To run a go file as a script:
`$ go run hello.go`
This creates a binary file in a temp directory and deletes it when the program finishes.

To build a binary for later use:
`$ go build hello.go`
This creates an executable with the same name as the go file in the same directory. To change the name of the executable:
`$ go build -o hello_world hello.go`

This binary can be directly run in terminal:
```
$ chmod +x hello_world
$ ./hello_world
```

#### 1.1 Installing Third-Party Go Tools

Third-party packages are installed directly from source code repositories (rather than a centrally hosted service):
`$ go install <repo_url>`

For example,
`$ go install github.com/rakyll/hey@latest`

This downloads the hey package and installs it in the `$GOPATH/bin` dir.

#### 1.2 Formatting Code

Go is strict with how it is formatted to both avoid "format wars" and to allow for greater efficiency when writing code.

To reformat code (fixes indentation, lining up fields in a struct, operator spacing):
`$ go fmt`

A further too called `goimports` takes this further by tidying imports:
`$ goimports -l -w`


#### 1.3 Linting and Vetting

The `golint` tool ensures that all go files adhere (syntactically) to the strict style of Go.

To install:
`$ go install golang.org/x/lint/golint@latest`

To run:
`$ golint ./`

The `vet` tool checks for further errors in syntactially correct code. It examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string.

`$ go vet ./`

#### 1.4 Makefiles
Makefiles specify build steps. The make utility requires a file, Makefile (or makefile), which defines a set of tasks to be executed. You may have used make to compile a program from source code.

To run:
`$ make`

When run in `ch1` dir, it will output:

```
$ go fmt ./...
$ go vet ./...
$ go build -o hello_world hello.go
```
