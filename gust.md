# Gust Language

Gust is a interpreted, statically-typed language written in Go

The syntax of Gust is based on a combination of Go and Rust
```
fn greet(name: str) -> str {
    return "hello " .. name .. "! Nice to meet you!"
}

let greeting = greet("nick")

println(greeting)
```

```
fn fizzbuzz() {
    for i ;= 1, i != 100, i++ {
        if i % 3 == 0 && i % 5 == 0 {
            println("fizzbuzz")
        } else if i % 3 == 0 {
            println("fizz")
        } else if i % 5 == 0 {
            println("buzz")
        } else {
            println(i)
        }
    }
}

fizzbuzz()
```

```
fn loopFib(n: int) {
    a ;= 0
    b ;= 1
    for i ;= 0, i < n, i++ {
        println(a)
        c ;= a
        a = b
        b ;= c
    }
}

loopFib(10)
```
