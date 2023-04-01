# GO

## Run go program
- `go run helloWorld.go`

## Performance
- `/usr/bin/time -l -h -p go run dataType.go`

## Variable Declaration

## Double Quote vs Single Quote

## printf - format specifier

## General:
- %v	the value in a default format. when printing structs, the plus flag (%+v) adds field names
- %#v	a Go-syntax representation of the value
- %T	a Go-syntax representation of the type of the value
- %%	a literal percent sign; consumes no value

## Boolean:
- %t	the word true or false

## Integer:
- %b	base 2
- %c	the character represented by the corresponding Unicode code point
- %d	base 10
- %o	base 8
- %O	base 8 with 0o prefix
- %q	a single-quoted character literal safely escaped with Go syntax.
- %x	base 16, with lower-case letters for a-f
- %X	base 16, with upper-case letters for A-F
- %U	Unicode format: U+1234; same as "U+%04X"

## Floating-point and complex constituents:
- %b	decimalless scientific notation with exponent a power of two, in the manner of strconv.FormatFloat with the 'b' format, e.g -123456p-78
- %e	scientific notation, e.g. -1.234456e+78
- %E	scientific notation, e.g. -1.234456E+78
- %f	decimal point but no exponent, e.g. 123.456
- %F	synonym for %f
- %g	%e for large exponents, %f otherwise. Precision is discussed below.
- %G	%E for large exponents, %F otherwise
- %x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
- %X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

## String and slice of bytes (treated equivalently with these verbs):
- %s	the uninterpreted bytes of the string or slice
- %q	a double-quoted string safely escaped with Go syntax
- %x	base 16, lower-case, two characters per byte
- %X	base 16, upper-case, two characters per byte

## Slice:
- %p	address of 0th element in base 16 notation, with leading 0x

## Pointer:
%p	base 16 notation, with leading 0x
The %b, %d, %o, %x and %X verbs also work with pointers,
formatting the value exactly as if it were an integer.
The default format for %v is:

bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p


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
  - GoPls - A Go server developed by Go Team `go install -v golang.org/x/tools/gopls@latest`
  
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
- Example

## Data types
- Int
- Float
- String
- Boolean

## Functions
- Functions as variable
- Functions as argument
- Functions as anonymous

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