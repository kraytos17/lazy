package main

import (
	"fmt"

	lazy "github.com/kraytos17/lazy/lazy"
)

func main() {
	fmt.Println("Eager sum:\t", lazy.Sum(func() int { return 10 + 5 }, func() int { return 20 })())

	lazySum := func(a, b lazy.Lazy[int]) lazy.Lazy[int] {
		return lazy.Sum(a, b)
	}
	fmt.Println("Lazy sum:\t", lazySum(func() int { return 10 + 5 }, func() int { return 20 })())

	fmt.Println("Lazy First: ", lazy.First(func() int { return 10 }, func() int { return lazy.Hang().(int) })())

	fmt.Println("\n==============================")

	fmt.Println("false && false ==", lazy.And(lazy.Trace(func() bool { return false }, "L"),
		lazy.Trace(func() bool { return false }, "R"))())
	fmt.Println("true && false ==", lazy.And(lazy.Trace(func() bool { return true }, "L"),
		lazy.Trace(func() bool { return false }, "R"))())
	fmt.Println("true && true ==", lazy.And(lazy.Trace(func() bool { return true }, "L"),
		lazy.Trace(func() bool { return true }, "R"))())
	fmt.Println("false && true ==", lazy.And(lazy.Trace(func() bool { return false }, "L"),
		lazy.Trace(func() bool { return true }, "R"))())

	fmt.Println("---")
	fmt.Println("false || false ==", lazy.Or(lazy.Trace(func() bool { return false }, "L"),
		lazy.Trace(func() bool { return false }, "R"))())
	fmt.Println("true || false ==", lazy.Or(lazy.Trace(func() bool { return true }, "L"),
		lazy.Trace(func() bool { return false }, "R"))())
	fmt.Println("true || true ==", lazy.Or(lazy.Trace(func() bool { return true }, "L"),
		lazy.Trace(func() bool { return true }, "R"))())
	fmt.Println("false || true ==", lazy.Or(lazy.Trace(func() bool { return false }, "L"),
		lazy.Trace(func() bool { return true }, "R"))())

	fmt.Println("\n==============================")

	fmt.Println(lazy.ToList([]int{1, 2, 3}))
	fmt.Println(lazy.ToList([]int{1, 2, 3})())
	fmt.Println(lazy.ToList([]int{1, 2, 3})().Head())
	fmt.Println(lazy.ToList([]int{1, 2, 3})().Tail().Head())
	fmt.Println(lazy.ToList([]int{1, 2, 3})().Tail().Tail().Head())
	fmt.Println(lazy.ToList([]int{1, 2, 3})().Tail().Tail().Tail())

	fmt.Println("---")

	fmt.Println(lazy.Range(func() int { return 3 }))
	fmt.Println(lazy.Range(func() int { return 3 })())
	fmt.Println(lazy.Range(func() int { return 3 })().Head())
	fmt.Println(lazy.Range(func() int { return 3 })().Tail().Head())
	fmt.Println(lazy.Range(func() int { return 3 })().Tail().Tail().Head())
	fmt.Println(lazy.Range(func() int { return 3 })().Tail().Tail().Tail().Head())
	fmt.Println(lazy.Range(func() int { return 3 })().Tail().Tail().Tail().Tail().Head())

	fmt.Println("---")

	lazy.PrintList(lazy.ToList([]int{1, 2, 3, 4, 5}))

	fmt.Println("---")

	lazy.PrintList(lazy.Take(func() int { return 10 }, lazy.Range(func() int { return 3 })))

	fmt.Println("---")

	lazy.PrintList(
		lazy.Take(func() int { return 10 },
			lazy.Filter(func(x int) bool { return x%2 == 0 },
				lazy.Range(func() int { return 1 }))))

	fmt.Println("\n==============================")

	prime := lazy.Sieve(lazy.Range(func() int { return 2 }))

	lazy.PrintList(lazy.Take(func() int { return 10 }, prime))
}
