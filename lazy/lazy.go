package lazy

import "fmt"

type Lazy[T any] func() T

// LazyList represents a lazy list.
type LazyList[T any] struct {
	Head Lazy[T]
	Tail Lazy[*LazyList[T]]
}

// Sum performs a sum of two lazy numbers.
func Sum(a, b Lazy[int]) Lazy[int] {
	return func() int {
		return a() + b()
	}
}

// First returns the first of two lazy values.
func First[T any](a, b Lazy[T]) Lazy[T] {
	return a
}

// And performs a logical AND operation on two lazy booleans.
func And(a, b Lazy[bool]) Lazy[bool] {
	return func() bool {
		if !a() {
			return false
		}

		return b()
	}
}

// Or performs a logical OR operation on two lazy booleans.
func Or(a, b Lazy[bool]) Lazy[bool] {
	return func() bool {
		if a() {
			return true
		}

		return b()
	}
}

// Trace wraps a lazy computation with a log message.
func Trace[T any](x Lazy[T], message string) Lazy[T] {
	return func() T {
		fmt.Println(message)
		return x()
	}
}

// ToList creates a lazy list from a slice.
func ToList[T any](xs []T) Lazy[*LazyList[T]] {
	return func() *LazyList[T] {
		if len(xs) == 0 {
			return nil
		}

		return &LazyList[T]{
			Head: func() T { return xs[0] },
			Tail: ToList(xs[1:]),
		}
	}
}

// Range creates an infinite lazy list of numbers starting from a given number.
func Range(begin Lazy[int]) Lazy[*LazyList[int]] {
	return func() *LazyList[int] {
		x := begin()
		return &LazyList[int]{
			Head: func() int { return x },
			Tail: Range(func() int { return x + 1 }),
		}
	}
}

// PrintList prints the elements of a lazy list.
func PrintList[T any](xs Lazy[*LazyList[T]]) {
	pair := xs()
	for pair != nil {
		fmt.Println(pair.Head())
		pair = pair.Tail()
	}
}

// Take creates a new lazy list by taking the first n elements of the input list.
func Take[T any](n Lazy[int], xs Lazy[*LazyList[T]]) Lazy[*LazyList[T]] {
	return func() *LazyList[T] {
		m := n()
		pair := xs()
		if m > 0 && pair != nil {
			return &LazyList[T]{
				Head: pair.Head,
				Tail: func() *LazyList[T] {
					return Take(func() int { return m - 1 }, func() *LazyList[T] {
						return pair.Tail()
					})()
				},
			}
		}
		return nil
	}
}

// Filter creates a new lazy list by filtering the input list with a predicate.
func Filter[T any](f func(T) bool, xs Lazy[*LazyList[T]]) Lazy[*LazyList[T]] {
	return func() *LazyList[T] {
		pair := xs()
		if pair == nil {
			return nil
		}

		x := pair.Head()
		if f(x) {
			return &LazyList[T]{
				Head: func() T { return x },
				Tail: func() *LazyList[T] {
					return Filter(f, func() *LazyList[T] {
						return pair.Tail()
					})()
				},
			}
		}
		
		return Filter(f, func() *LazyList[T] {
			return pair.Tail()
		})()
	}
}

// Sieve creates a new lazy list of prime numbers.
func Sieve(xs Lazy[*LazyList[int]]) Lazy[*LazyList[int]] {
	return func() *LazyList[int] {
		pair := xs()
		if pair == nil {
			return nil
		}

		y := pair.Head()
		return &LazyList[int]{
			Head: func() int { return y },
			Tail: func() *LazyList[int] {
				return Sieve(Filter(func(x int) bool { return x%y != 0 }, func() *LazyList[int] {
					return pair.Tail()
				}))()
			},
		}
	}
}

var Hang func() any

func Init() {
	Hang = func() any {
		return Hang()
	}
}
