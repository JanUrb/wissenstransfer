# Intro to Go

- [Intro to Go](#intro-to-go)
  - [What is Go and what is its History?](#what-is-go-and-what-is-its-history)
  - [Code Syntax Examples - (Next time you see go code you will be able to read it)](#code-syntax-examples---next-time-you-see-go-code-you-will-be-able-to-read-it)
  - [Concurrency](#concurrency)

## What is Go and what is its History?

* Created by: Google -> Version 1.0 in 2012
* Go is a general-purpose language, designed as system (server) language
* Go Features:
  * List of Features
  * Cross compilation, compiles to binaries -> no VM or Engine needed
  * Static Typing but automatic type inference (x := 1 or y := "hi")
  * std lib provides a lot of functionlity(e.g: web and testing)
  * Simple language with focus on readability (official style guide that is enforced by the compiler)

## Code Syntax Examples - (Next time you see go code you will be able to read it)

* Variable Declaration
* If -> Point to braces style, Talk about gofmt, go vet etc.
* For
* Struct
* Public/Private + Methods

## Concurrency

* Explain the difference between Threads and Goroutines (Scheduler, OS Thread mapping etc), Deadlock
* Concurrency Model: CSP -> Same as Erlang
* Channels: Share by communication -> https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html
* Buffered, unbuffered channel
* Iterating over channel
* Select (Random nature of select)
