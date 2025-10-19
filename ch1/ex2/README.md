# Chapter 1 Exercise 2

## ✨ Question

Add a target to the Makefile called `clean` that removes the `hello_world` binary and any other temporary files created by `go build`. Take a look at the [Go command documentation](https://pkg.go.dev/cmd/go) to find a `go` command to help implement this.

## 💡 Solution

By adding the following code to the `makefile`, I managed to remove binary files:

```makefile
clean:
	go clean
```

## 🪝 References

[Solution of chapter 1 exercise 2](https://github.com/learning-go-book-2e/ch01/blob/main/exercise_solutions/ex2) on GitHub.
