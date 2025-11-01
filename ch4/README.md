# Chapter 4

## 🧩 Recap

In Chapter 4, I learnt about block, scope, and shadowing in `Go`, understanding how redeclaring a variable within an inner block may lead to confusion or errors.
But thanks to the powerful IDE, GoLand does highlight shadowed variable like:

![](/ch4/ex3/assets/shadowing.png)

The chapter also covered `Go`'s control flow statements, including `if`, `for`, `switch`, `break`, `continue`, and `goto`.
I also learnt how `for range` iterates over arrays, slices, maps, and strings and how labels work with `for` and `goto`.

## 💭 Reflection

- GoLand highlights shadowed variables by default.
- Appreciated that `Go`'s switch statements don't fall through by default, improving clarity.
- `goto` is evil.
- Don't use `goto`.
- There's a slight chance to use `goto`.

## 🪝 References

The standard library [math/rand](https://pkg.go.dev/math/rand) on Go Packages (pkg.go.dev).

[Solution of chapter 4 exercise 1](https://github.com/learning-go-book-2e/ch04/tree/main/exercise_solutions/ex1) on GitHub.

[Solution of chapter 4 exercise 2](https://github.com/learning-go-book-2e/ch04/tree/main/exercise_solutions/ex2) on GitHub.

