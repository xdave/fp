package fp_test

import (
	"fmt"
	"strings"
	"testing"
)

import (
	"github.com/xdave/fp"
)

func TestMapWrongTypePanics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Did not panic as expected")
		}
	}()
	input := "someOddType"
	fp.Map(input, func() {})
}

func TestReduceWrongTypePanics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Did not panic as expected")
		}
	}()
	input := "someOddType"
	fp.Reduce(input, func() {})
}

func ExampleMapSliceString() {
	input := []string{"Hello", "World"}
	fmt.Println(fp.Map(input, func(item string, i int) string {
		return strings.ToUpper(item)
	}))
	// Output: [HELLO WORLD]
}

func ExampleMapSliceInt() {
	input := []int{1, 2, 3, 4, 5}
	fmt.Println(fp.Map(input, func(item, i int) int {
		return item * 5
	}))
	// Output: [5 10 15 20 25]
}

func ExampleMapMapStringInt() {
	input := map[string]int{"two": 1, "four": 2, "six": 3, "eight": 4}
	result := fp.Map(input, func(item int, i string) int {
		return item * 2
	}).(map[string]int)
	fmt.Println(result["two"], result["four"], result["six"], result["eight"])
	// Output: 2 4 6 8
}

func ExampleReduceSlice() {
	input := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println(fp.Reduce(input, func(a, b float64) float64 {
		return a + b
	}))
	// Output: 15
}

func ExampleReduceMap() {
	input := map[string]int{"foo": 5, "bar": 6, "baz": 7}
	fmt.Println(fp.Reduce(input, func(a, b int) int {
		return a * b
	}))
	// Output: 210
}

func ExampleMapReduceSlice() {
	input := []int{1, 2, 3, 4}
	fmt.Println(fp.Reduce(fp.Map(input, func(item, i int) int {
		return item * 5
	}), func(a, b int) int {
		return a + b
	}))
	// Output: 50
}
