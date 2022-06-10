package main

import (
	"github.com/emirpasic/gods/sets/hashset"
)

func main() {
	set := hashset.New()   // empty
	set.Add(1)             // 1
	set.Add(2, 2, 3, 4, 5) // 3, 1, 2, 4, 5 (random order, duplicates ignored)
	set.Clear()            // empty
	set.Empty()            // true
	set.Size()             // 0
	println(gcd(10, 100))
}
