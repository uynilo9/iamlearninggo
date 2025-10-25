# Chapter 2 Exercise 2

## ✨ Question

Write a program where you declare a constant called `value` that can be assigned to both an integer and a floating point variable. Assign it to an integer called `i` and a floating point variable called `f`. Print out `i` and `f`.

## 💡 Solution

First, I declared a constant value with the value `213`.

Constants in `Go` are untyped, which means you can assign a numeric constant to any numeric variable type, as long as it doesn’t cause an overflow.
Therefore, I assigned value to a variable `i` without specifying the type (obviously `213` is `int` by default), and to another variable `f` explicitly typed as `float64`.

Finally, I printed both variables.

## 🪝 References

[Solution of chapter 2 exercise 2](https://github.com/learning-go-book-2e/ch02/tree/main/exercise_solutions/ex2) on GitHub.
