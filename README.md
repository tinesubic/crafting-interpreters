# GLOX

This repository contains my code for following along with Crafting Interpreters book by Bob Nystrom: [link](https://www.craftinginterpreters.com/introduction.html#design-note)

* Implementation in Go
* To run, use [`just`](https://github.com/casey/just) task runner:
  * REPL: `just glox-repl`
  * Standard: `just glox <file>`

## Chapter 1: Introduction
### Challenges: 
* There are at least six domain-specific languages used in the little system I cobbled together to write and publish this book. What are they
  * Make
  * Markdown
  * XML
  * Dart
  * YAML
  * HTML, CSS
* [x] Get a “Hello, world!” program written and running in Java. Set up whatever makefiles or IDE projects you need to get it working. If you have a debugger, get comfortable with it and step through your program as it runs.

## Chapter 2: Map of the Territory
### Challenges:
* Pick an open source implementation of a language you like. Download the source code and poke around in it. Try to find the code that implements the scanner and parser. Are they handwritten, or generated using tools like Lex and Yacc? (.l or .y files usually imply the latter.). Go: seems like handwritten
* Just-in-time compilation tends to be the fastest way to implement dynamically typed languages, but not all of them use it. What reasons are there to not JIT? Avoid recompilation on target - performance overhead on startup.
* Most Lisp implementations that compile to C also contain an interpreter that lets them execute Lisp code on the fly as well. Why? Can be used as REPL, compiled can be deployed to other targets. For development?


