package main

import "fmt"

// ex1: Test go.mod golang version. By using generics that was introduced in go1.18,
// we can see that the go.mod file is set to go1.18.
// When we change it to go1.17, we get an error:
// # ex1
// ./main.go:7:2: embedding interface element int|float32|float64|complex64|complex128 requires go1.18 or later (-lang was set to go1.17; check go.mod)
// ./main.go:11:10: type parameter requires go1.18 or later (-lang was set to go1.17; check go.mod)
// ./main.go:14:3: invalid operation: operator + not defined on sum (variable of type T constrained by Numbers)
// ./main.go:20:30: implicit function instantiation requires go1.18 or later (-lang was set to go1.17; check go.mod)
// ./main.go:21:34: implicit function instantiation requires go1.18 or later (-lang was set to go1.17; check go.mod)
// ./main.go:22:33: implicit function instantiation requires go1.18 or later (-lang was set to go1.17; check go.mod)
// ./main.go:23:34: implicit function instantiation requires go1.18 or later (-lang was set to go1.17; check go.mod)

// Numbers is a type that can be added
type Numbers interface {
	int | float32 | float64 | complex64 | complex128
}

// Add adds all the numbers in the slice
func Add[T Numbers](n []T) T {
	var sum T
	for _, v := range n {
		sum += v
	}
	return sum
}

// Main is the entry point
func main() {
	fmt.Printf("ints: %d\n", Add([]int{1, 2, 3}))
	fmt.Printf("floats32: %f\n", Add([]float32{1.1, 2.2, 3.3}))
	fmt.Printf("complex: %f\n", Add([]complex64{1 + 1i, 2 + 2i, 3 + 3i}))
	fmt.Printf("floats62: %f\n", Add([]float64{1.1, 2.2, 3.3}))
}
