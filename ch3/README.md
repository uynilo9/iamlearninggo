# Chapter 3

## 🧩 Recap

In Chapter 3, I learnt about 4 different data structures in `Go`: `array`, `slice`, `map`, `struct`.

Slices provide dynamic length and capacity with functions such as `append`, `len`, `cap`, `copy`, etc.
Also, it explained that slices share memory with their underlying arrays.

Maps store key-value pairs for fast lookup and updates, requiring when handling missing keys using the `, ok` idiom.

Structs define custom and composite types with support for nested or anonymous fields, promoting organised data representation.
Besides, unlike the above-mentioned data types, structs are defined using the keyword `type`.

## 💭 Reflection

- It's extremely important that slices share memory with their subslices unless user uses `copy`.
- How length and capacity work together is crucial.
- Learnt to use the `, ok` idiom to check for key existence in maps.
- Understood the importance of using `make` to initialise slices and maps properly.

## 🪝 References

[Solution of chapter 3 exercise 1](https://github.com/learning-go-book-2e/ch03/tree/main/exercise_solutions/ex1) on GitHub.

[Solution of chapter 3 exercise 2](https://github.com/learning-go-book-2e/ch03/tree/main/exercise_solutions/ex2) on GitHub.

[Solution of chapter 3 exercise 3](https://github.com/learning-go-book-2e/ch03/tree/main/exercise_solutions/ex3) on GitHub.
