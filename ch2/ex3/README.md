# Chapter 2 Exercise 3

## ✨ Question

Write a program with three variables, one named `b` of type `byte`, one named `smallI` of type `int32`, and one named `bigI` of type `uint64`. Assign each variable the maximum legal value for its type, then add `1` to each variable. Print out their values.

## 💡 Solution

The type `byte` is essentially an alias of the type `uint8`.

I declared `b`, `smallI`, and `bigI` with the types specified in the question, assigning the maximum legal value to each of them from the standard library `math`.

Next, I added `1` to each variable and printed them out.

As expected, all the variables overflowed to their minimum values.

## 🪝 References

The standard library [math](https://pkg.go.dev/math) on Go Packages (pkg.go.dev).

[Solution of chapter 2 exercise 3](https://github.com/learning-go-book-2e/ch02/tree/main/exercise_solutions/ex3) on GitHub.
