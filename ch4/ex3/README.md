# Chapter 4 Exercise 3

## ✨ Question

Start a new program.In `main`, declare an `int` variable called `total`.
Write a `for` loop that uses a variable named `i` to iterate from 0 (inclusive) to 10 (exclusive).
The body of the `for` loop should be as follows:
```go
total := total + i
fmt.Println(total)
```
After the `for` loop, print out the value of `total`.

What is printed out?

What is the likely bug in this code?

## ✏️ Answer

It printed the numbers from 0 to 9 separately.

The `:=` shadows the `total` on the left hand side as a new variable in the `for` loop scope, since `:=` is able to redeclare variables. [^1]

By the way, this behaviour is extremely important and mentioned in Chapter 2.

Still by the way, GoLand highlights shadowed variables. See the picture below:

![](/ch4/ex3/assets/shadowing.png)

## 💡 Solution

By replacing `:=` with `=`, I made the program correctly modify the `total` variable declared outside the `for` loop (adding `i` to it) and avoided shadowing.

## 🪝 References

[Solution of chapter 4 exercise 3](https://github.com/learning-go-book-2e/ch04/tree/main/exercise_solutions/ex3) on GitHub.

The [Short variable declarations](https://go.dev/ref/spec#Short_variable_declarations) on The Go Programming Language Specification (go.dev/ref/spec).

[^1]: https://go.dev/ref/spec#Short_variable_declarations#:~:text=a%20short%20variable%20declaration%20may%20redeclare%20variables%20provided%20they%20were%20originally%20declared%20earlier%20in%20the%20same%20block
