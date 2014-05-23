# Gosick

Scheme interpreter implemented by Golang.  
This is started as [a programming project](https://github.com/k0kubun/gosick/blob/master/project.md) for newcomers of my laboratory.

## Installation

```bash
$ go get github.com/k0kubun/gosick
```

## Usage

```bash
# Invoke interactive shell
$ gosick

# Excecute scheme source code
$ gosick source.scm

# One liner
$ gosick -e "(+ 1 2)"

# Dump AST of input source code
$ gosick -a

# Show help
$ gosick -h
```

## Implemented syntax and functions
### Done
- +, -, *, /, =, <, <=, >, >=
- cons, car, cdr, list, length, last, append, set-car!, set-cdr!
- if, cond, and, or, not, begin, do
- null?, number?, boolean?, procedure?, pair?, list?, symbol?, string?
- string-append, symbol->string, string->symbol, string->number, number->string
- lambda, define, set!
- load

### Pending
| Name | Description |
|:-----|:------------|
| let, let*, letrec | eval order is incorrect, and they are the same implementation |
| memq, eq?, neq?, equal? | only Number and Symbol are supported |
| write, print | second argument (choosing output port) is not implemented |

### To be done
- define-macro

## Pending features
- Tail Call Optimization
- Non-integer number

## License

Gosick is released under the [MIT License](http://opensource.org/licenses/MIT).
