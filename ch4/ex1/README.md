# Chapter 4 Exercise 1

## ✨ Question

Write a `for` loop that puts 100 random numbers between 0 and 100 into an `int` slice.

## 💡 Solution

By creating a slice and using a `for` loop, I made the program generate a random number between 0 and 100 with the `math.Intn()` function, and append it to the slice a hundred times.

By the way, the question says "between," which means 100 should be inclusive. Therefore, I called the `math.Intn()` function with the argument `101`, since it generates numbers randomly in the range `[0, n)` (`n` is exclusive) [^1].

## 🪝 References

The standard library [math/rand](https://pkg.go.dev/math/rand) on Go Packages (pkg.go.dev).

[Solution of chapter 4 exercise 1](https://github.com/learning-go-book-2e/ch04/tree/main/exercise_solutions/ex1) on GitHub.

[^1]: https://pkg.go.dev/math/rand#Intn#:~:text=%5B0%2Cn%29
