# fener

`fener` is a interpreted ruby/lua-like language. It is a object-oriented programming language with first-class functions.

 - [Fener](#hotshot)
    - [status](#status)
    - [getting started](#getting-started)
    - [testing](#testing)
    - [development](#development)
    - [contribution](#contribution)
 - [Semantics](#semantics)
    - [expressions](#expressions)
    - [variables](#variables)
    - [variables](#comments)
    - [flow-control](#flow-control)
    - [functions](#functions)
    - [builtin](#builtin)
    - [class](#class)
 - [Data Structures](#data-stucture)
    - [list](#list)
    - [maps](#maps)
 - [Import](#import)
    - [import](#import)
    - [stdlib](#stdlib)
 - [Internals](#internals)
    - [lexer](#lexer)
    - [parser](#parser)
    - [eval](#eval)

This is `fener`. A easy language written in Golang.

It's a dynamicaly typed language with object-oriented features.
It is designed similar to languages like `Ruby` or `Python`.

Here's the `FizzBuzz` program written in fener

```go
fn fizzbuzz(n)
    if n % 5 == 0 && n % 3 == 0 then
        print("FizzBuzz")
    elif n % 3 == 0 then
        print("Fizz")
    elif n % 5 == 0 then
        print("Buzz")
    else
        print(n)
    end
end
```

# Status

This language is a hobby project under heavy development.
It's versatile and capable enough for small programs.

It's designed for reading and experimenting.

# Getting Started

To use fener, you need the single binary.

## Go

You can get fener by using the `Go` compiler.

```sh {linenos=false}
go install github.com/pspiagicw/fener@main
```

Or if you use [gox](https://github.com/pspiagicw/gox).

```sh {linenos=false}
gox install github.com/pspiagicw/fener@main
```

## Compile

To compile the project, you only need the `Go` compiler.

```sh {linenos=false}
git clone https://github.com/pspiagicw/fener

cd fener

go build .
```

## Testing

fener has a extensive suite of tests.
It's written in Go's native test runner.

Which depends on the fener binary being compiled inside the project directory.

```
go build .

go test ./...
```

You can run specific tests or get information about all subtests.

```sh
# Run only lexer tests, but show all subtests
go test -v ./lexer
```

# Contribution

This project is under heavy development and contributions are highly appreciated.
A lot of decisions are yet to be taken, and you can be part of them.



# Semantics

`fener` is a expression based language.
Every valid statmeent/expression returns a value.

You can interact with `fener` using the REPL.

> Run the repl using `fener repl`.

## Arithmetic

You can run arithmetic expresions.

```sh {linenos=false}
>>> 1 + 2
int(3)
```

```sh {linenos=false}
>>> 1 + 2 * 6 / 7 - 1
int(1)
```
fener has 3 fundamental data types.

```sh {linenos=false}
>>> 10
int(10)
>>> true
bool(true)
>>> false
bool(false)
>>> "this is a string"
str(this is a string)
```

Complex data types include `classes`, `list` and `maps`.

> `lists` and `maps` are not implemented yet. 

- Lists

```sh {linenos=false}
>>> [1 2 3 4 5]
[int(1) int(2) int(3) int(4) int(5)]
```

- Maps

```sh {linenos=false}
>>> { "name" = "Chris" "surname" = "Pratt" }
{ str(name)->str(Chris) str(surname)->str(Pratt)}
```

> Classes are covered later.

There is a bonus type `null`, which is returned by builtin functions and some statements.
It can't be used by the user.

## Variables

Variables don't have any type, they can hold value of any type.

The assignment expression `<name> = <variables>` returns the value.

```sh {linenos=false}
>>> a = 20
int(20)
>>> b = 10
int(10)
>>> a + b
int(30)
>>> a = "something"
str(something)
>>> a
str(something)
```

## Comments

```lisp {linenos=false}
;; these are comments.
;; Comments are marked by 2 consecutive semicolons.
```

## Flow

`fener` supports if-expressions and while-statements.

## Functions

You can declare functions using the `fn` keyword.

## Builtin


## Class



