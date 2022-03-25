# iter [![Go Reference](https://pkg.go.dev/badge/github.com/rodaine/iter.svg)](https://pkg.go.dev/github.com/rodaine/iter)

_A generic Go iterator library inspired by Rust's [std::iter::Iterator](https://doc.rust-lang.org/std/iter/trait.Iterator.html)_

```go
isEven := func(n int) bool { return n%2 == 0 }

add := func(a, b int) int { return a + b }

sum, ok := CountUpBy(0, 3).   // 0, 3, 6, 9, 12, ...
             Filter(isEven).  // 0, 6, 12, 18, ...
             Take(5).         // 0, 6, 12, 18, 24
             Reduce(add)      // 60
```

## Installation

**Requires:** Go 1.18+

```sh
go get github.com/rodaine/iter
```

## Limitations

Go's generics are powerful, but still limited. This affects the ergonomics of some usecases. The common thread in all cases is that methods must be fully instantiated/monomorphized at compile-time.  

### Methods cannot have type parameters

_Associated Issue:_ [golang/go#49085](https://github.com/golang/go/issues/49085)

Because methods cannot have type parameters (beyond the types defined on the receiver type), operations that would map types cannot be implemented as methods. This makes `Map`, `Fold`, and others not currently chainable; however, they have been implemented as free functions:

```go
// WANT
// func (iter Iterator[A]) Map[B any](fn MapFunc[A, B]) Iterator[B] { ... }

FromItems(1, 2, 3).
  Map(strconv.Itoa).
  ToSlice() // ["1", "2", "3"]
  
// GOT
// func Map[A, B any](iter Iterator[A], fn MapFunc[A, B]) Iterator[B] { ... }

Map(
  FromItems(1, 2, 3), 
  strconv.Itoa,
).ToSlice() // ["1", "2", "3"]
```

### Types/Methods cannot cause an instantiation cycle

_Associated Issue:_ [golang/go#50215](https://github.com/golang/go/issues/50215)

When a type is instantiated in Go, all its methods must also be instantiated at compile time, even if that method is not necessarily used. This is required as methods may be called dynamically via interface assertions or reflection. For `Enumerate`, which maps an `Iterator[E]` to `Iterator[Pair[int, E]]`, defining this as a method would require for the compiler to also instantiate `Iterator[Pair[int, Pair[int, E]]]`, and so on infinitely. This does work as a free function, however.

```go
// WANT
// func (iter Iterator[A]) Enumerate() Iterator[Pair[A, B]] { ... }

FromItems('a', 'b', 'c').
  Enumerate().
  ToSlice() // [{0, 'a'}, {1, 'b'}, {2, 'c'}]
  
// GOT
// func Enumerate[A](iter Iterator[A]) Iterator[Pair[int, A]] { ... }

Enumerate(
  FromItems('a', 'b', 'c'),
).ToSlice() // [{0, 'a'}, {1, 'b'}, {2, 'c'}]
```

### Specialization is not supported

Certain functionality, such as `Sum` and `Equality`, require specializations or narrowing of the generic types to work. Go generics currently do not support specialization of methods, but it can be achieved with free functions.

```go
// WANT
// func (iter Iterator[N Number]) Sum() N { ... }

FromItems(1, 2, 3).
  Sum() // 6

// GOT
// func Sum[N Number](iter Iterator[N]) N { ... }

Sum(
  FromItems(1, 2, 3),
) // 6
```

## Performance

The performance characteristics of this module have not yet been evaluated. 