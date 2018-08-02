# Command line application template for Go

Implement CLI application by editing [main.go](main.go).  
You may add new files to keep your code clean, if it is allowed in your challenge.

## How to get input parameters
You can get arguments as `args` in [main.go](main.go) file where the `run` method is defined.  

``` go
func run(args []string) {
  // code to run
}
```

`args` is simply came from `os.Args`, passed by `main` function.
It passes command line arguments without its script name.

## How to output result
You can use `fmt.Println` methods to output your results.

``` go
fmt.Println(args)
```

## Install External Libraries
If you want to use external libraries, do the following:

- Add the following lines to [codecheck.yml](codecheck.yml), before the `go build` line  
(You can have multiple libraries by adding more lines)

``` yaml
build:
  - go get namespace.of/some/library
```
