# GO Lang Learningss

This repo contains source code examples and notes from the go language exploration tutorials

## 01_Introduction

Go lang is a high Performance language. It is

- Open Source, Maintained By Google
- Provides vast functionality support via Standard Libraries
- is Statically Typed
- runs fast, because it gets compiled to machine languages

## 02_Why Choose Go Language

Following are five reasons to choose go language

- Quick Build Time : Applications using go gets build quickly
- Fast Startup : Applications will have fast startup time
- Performance and Efficiency : Go Apps use system resources efficiently and with great performance
- Concurrency: Go lang have goroutines (lightweight Threads) which effectively uses multi-threading
- Static Typing and Compilation: Supports Startic typing, and handling compile time errors at development time

## 03_Installation

Installation is simple just go to go website and download the binary files suitable for the OS you are running.
after installation run

> go version

this should output go version installed.

## 04_Hello_World

- first we define the package name
- then we define the main function using func keyword
- then we use Println method from fmt package to print the information

- to build the program

```bash
go build main.go

// outputs main.exe file
```

- to run directly

```bash
go run main.go // shows output directly
```

## 05_Simple Values

Go language supports following basic datatypes
- Integer
- Strings
- Floats
- Booleans

## 06_Variables

Variables in go can be defined in following ways

```go
var name string // declaration
name = "Om Pant" // Assignment

var name2 string = "Prateek" // initialization

var name3 = "Ajay" // initialization with type inference

name4 := "Himanshu" // shorthand initialization
```

## 07_Constants

We can define constant values in go using const keyword.

```go
const name = "Yash"

name = "Abc" // Will throw Error
```

- all rules applies to const declaration as we declared variables
- we can also define const and var's outside function, but the shorthand notation will not work outside function.

> Constant Groupings
to define multiple constants, we can use constant grouings

```go
const (
    port = 5500
    host = "localhost"
)

fmt.Println(host, port)
```