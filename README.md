# GO

## Run go program
- `go run helloWorld.go`

## Performance
- `/usr/bin/time -l -h -p go run dataType.go`

## Variable Declaration

## Double Quote vs Single Quote

## printf - format specifier
- %v = formats value in default format
- %d = for decimal integers
- %T =  print dataType of var
- %t = true or false
- %c = print character
- %.2f = floating point numbers upto 2 decimals


## What extensions should installed?
  - Gocode - a helper tool which is intended to be integrated with your source code editor
  - Gopkgs - a tool that provides list of available Go packages that can be imported. This is an alternative to `go list all`
  - GoOutline - a utility for extracting a JSON representation of the declarations in a Go source file. 
  - GoSymbols - a utility for extracting a JSON representation of the package symbols from a go source tree.
  - Guru - a tool for answering questions about Go source code.
  - GoRename - a command performs precise type-safe renaming of identifiers in Go source code.
  - Delve - a debugger for the Go programming language
  - GoCode - an autocompletion daemon for the Go programming language
  - GoReturns - adds zero-value return values to incomplete Go return statements, to save you time when writing Go.
  - GoLint - lints the Go source files named on its command line.
  
## How can I compile Go Code in my Terminal?
Open the terminal and navigate to the directory containing your file. For example, if your code file is HelloWorld.go, run the following command to compile.
```sh
go build HelloWorld.go
```
This will produce and output HelloWorld which is an executable. To run the file, execute the following command.

```sh
./HelloWorld
```

**To compile and run this file :**
1) Click on the ▶️ in the top right of the IDE to compile and run the go code. 

  More information on compiling code in VSCode can be found here:  https://code.visualstudio.com/docs/languages/go

## How can I Debug the code?
You can set break points by clicking to the left of line number where you want to break and investigate.
You can use the debugger in VSCode IDE using the icon on the left panel.

# Table of content

## Variable declaration
- EXAMPLE

## Data types
- INT
- FLOAT
- STRING

## Functions
- Functions as variable
- Functions as argument
- Function as anonymous

## Create your own workspace for multiple GO modules
```
go work init
go work use ./path-to-module  ./path-to-module2
```

Example
```
go work init
go work use ./BuildingModernWebApplicationsWithGo/31BuildingMoreComplexTemplateCache
```

# Reference
- [W3School](https://www.w3schools.com/go/index.php)
- [Workspace](https://stackoverflow.com/a/58524450/14010585)
- [Use JSON in GO](https://www.digitalocean.com/community/tutorials/how-to-use-json-in-go)