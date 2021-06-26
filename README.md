# Golang Playground

## The big picture

### Who created GO?

Go was not created by some hot-shot startup, or like a fad! It was created out of necessity to address specific problems with existing programming languages. Its created by some of the industries well known engineers. They have drawan from their experience and working with many languages to create GO.

1. Ken Thompson => the creator of C, Unix and UTF-8
2. Rob pike => Unix UTF8
3. Robert Griesemer => Hotspot JVM

Go is not just a google project through the creators worked at Google at some point. Its community driven.

### Why Create GO?

There wasn't one single programming language which could satisfy all 3 of:

1. Effecient compilation
2. Efficient execution
3. Ease of programming

Go is an effort to fill this void.

### Characteristics of GO

1. **Strong, static type system** - find bugs while you code, compiler helps you figure the type.
2. **C inspired syntax ( well almost)** - You'll see similarities if you come from C# or Java. This was an effort to reduce visual clutter of code. Its a c inspired syntax.
3. **Multi Paradigm ( procedural, OOP)** - Go lets you pick the paradigm based on context rather than imposing it right at the door.
4. **Garbage collected** - You dont have to manage memory on your own. Over years Go has optimised GC that it has no impact on performance
5. **Fully Compiled** - Unlike JS ( interpreted at runtime) or Java ( intermediate Bytecode), Go apps are fully compiled down to executable binary giving best performance.
6. **Rapid compilation** - A boon to developers who are now using TDD where they need to write and run tests while developing features
7. **Single binary output(mostly)** - By default you get a single binary making deployment easy, there is a also plugin eco system and shared lib to share.

### Timeline of GO

- 2007 - Start of design, when the 3 creators came together
- 2009 - Google publicly announced GO
- 2012 - Version 1.0 release
- Feb 2021 - GO 1.16 (latest as of this writing)

Go Has seen a timely release alsmost every 6 months since verion 1 came out in March of 2012. There is a discussion on version 2 which is in discussion phase. There is no confiremed timeline on version 2 yet. The community is focussed on making gsure go solves problems elegeanty, simply and well so the community is not rushing into the next version.

### Who's using GO in production?

Almost any big company you can think of is. Its most well suited for web based applications and APIs. think of companies like GOogle, Amazon, Digital Ocean, Netflix, Facebook, Cloudflare. Docker is actually written in GO.
If you wanna know who in the world is using it check [https://github.com/golang/go/wiki/GoUsers](https://github.com/golang/go/wiki/GoUsers)

### How can GO help you?

The philosophy and values that GO is built around:

- **Simplicity**: All around go the wat its built, the syntax etc is built around simplicity, doesnt mean easy or elementary but easy to approach and use.
- **Network aware and concurrent Apps**: Many of the existing languages like python, Java, C++ were built before we had multi core CPUs and internet as main stream, Go is designed with this out-of-the-box
- **Out-of-box Experience**: You dont have to go to different places to get all you need to develop apps, its right there right at the beginning.
- **Cross platform**: By nature its cross-platform, as more and more we are not constrained by specific platform
- **Backward compatibility**: GO is being activly developed with focus on backward compatibility

Lets look at each of these in some more detail

#### How GO embraces simplicity

Most of the languages start with this but get out of control as more features are added.

if you have used other languages like Java or C++ you would have seen these operators `++i` and `i++` The post increment and pre increment operators.

you'd see something like this while analysing something and then you have to guess what these burried increment operators actually do, and can confuse you leading to subtle bugs.

```c++
println(i++);
println(++i);
```

> So what does GO do with this? in GO the increment and decrement operators are statements and not expressions and not an expression. What this means is this is how Go treats the code.

```Go
i := 1
i++
println(i)
i++
println(i)
```

So in go we just have `i++` This solves the confusion. GO is more verbose and intentional, this helps making the code we write more cearler to communicate.
**Also the looping constructs in Go is simpler**

All you have to remember is just _one_ keyword `for`

The `for` loop used in different forms will work as the same way the while and infinite loops work in other languages

Now this is like any other for loop in other languages, thenormal for loop

```go
for i:=0; i<5;i++
```

If want a while loop kind of behavior we just use

```go
for i < 5
```

and if we need infinite loop we just use the keyword

```go
for
```

for looping over a collection we use

```go
for user := range users
```

> If you are looping in go you just use `for` loop thats it.

#### Network aware and concurrent Apps

Out of the box GO has capability

1.**net and net/http packages**: Core standard library for network aware features

- Create webservers using standard library

2.**goroutines**: light weight treads to manage massive concurrency

- Start thousands of concurrent tasks with minimal resources
- Goroutines are abstraction over processor threads

3.**Channels**: Built in support for communitation with concurrency (CSP) in a threadsafe way

- Safely communicate between concurrent tasks.

#### GO's out of box experience

Many languages require an extinsive amount of setup (external libraries and packages) in order to provide a basic development setup. The GO standard library itself has sufficient support to be up and running

- String Manipulation
- data compression
- File manipulation
- Networking APIs
- Testing is out of the box
- Many many more

The Go CLI which is a simple command line tool packs some powerful features to help you build test and deploy apps

- project initialization
- build
- Code generation
- retreive dependencies
- test
- profiling
- documentation
- report language bugs

There are third party libraries certainly, but the standard library is well equipped to do most of the work.

#### Simplifying cross platform development

With most of the prevalent programming languages , the apps have to me highly customised if it has to run on another platform than originally developed fore.g a C++ App that is built for windows will need extensive changes to if it has to also run on Linux. Same for native mobile apps - Android Vs ios.

with GO its baked in and trivial to impliment. just by changing two paramenters `GOOS` and `GOARCH` we can generate a build ofr an OS.

For e.g for windows

```go
GOOS = windows
GOARCH =amd64
```

For os x

```go
GOOS = darwin
GOARCH =amd64
```

or for android

```go
GOOS = android
GOARCH =arm
```

The key here is you dont need a mac to compile for osx or an android device to compile for android.

Additionally the standard library abstrats a whole lot of otherwise confusing items like file path handeling etc. You just code using one single approach and the compiler take care of it when you build for a particular OS.

#### GO's commitment to backward compatibility

The Go team has commited that if you are using Go version 1, even if version 2 comes out it should still work. However, if there are any security issues the GO team reserves the rights to go back and fix it even if it breaks compatibilty.

So you can rest assured that it will work and it will be secure!

for more (reassurance) checkout [https://golang.org/doc/go1compat](https://golang.org/doc/go1compat)

### GO's primary use cases

So what are the perfect applications of GO?

Top 3 areas where GO shines

- Web services: one of the primary and most wide spread use case, more often you will see API built on Go or GO backend in simple terms
- Web Applications : Similar to web services, web applications are also relying on GO more often
- DevOPs: GO is huge in DevOPs space, Doker is built on GO

Up and coming

- GUI / THick-Client:
- Machine Learning: The high-performance nature and trivial concurrency built with GO is luring more and more to look at it vs Python and R

GO is still fairly a new language and its use cases are growing as the community grows.

## Installing Go

Installing go is very straight forward head to the link below and select your operating system. Just select the defaults on installer and installation should complete in seconds ( after downloading!).

Download latest Go installer from [here](https://golang.org/dl/)

once the installation completes, check if the installation has completed successfuly by enetering this command in your terminal

```shell
go version
```

you should see an output similar to this with the go version installed

```shell
go version go1.16.3 darwin/amd64
```

This means you are reay to GO!

> If you are curious to see what Go tools has to offer just type `go` in the termial and hit return/enter and see what is printed. hint: there are bunch of things can do with just command line. All the critical tools are pre packaged in the go command.

## Setting up the Dev Environment

### Setup your code editor

You can use the editor of your choice, but if you ask me i would say use VS code as it has some pretty good support for go

Download latest version of VS code from [here](https://code.visualstudio.com/download). Ya its Free!

Once you have it installed. open it and on the left hand side go to **Extensions** here search for "go". The first one that pops up ( made by google) is the one you need. install it.

Next hit `cmd + shift + p` or `ctrl +shift + p` to bring up quick access and select/search for "Go: Install/update tools", select it, next select the checkbox on left at the top to select all and click OK. this will install all the required dependencies/packages to make yur GO development enjoyable.

## Sample Applications

A bunch of exercises, demo codelets and simple apps are included to take you from basics to expert

## Hello World

The obligatory hello world App is in the **01_Hello_World** Folder with self explanatory comments.

The intention of this very simple hellow world is well... its a hello world!

## Card Dealing APP

To dig deeper into GO lang concepts we will be building out a simple app that deals a deck of cards

Refer to the source code inside **02_go_cards** it comes with its own README for deeper explanations.
