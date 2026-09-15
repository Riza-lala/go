//go:build !solution

package counter

// New возвращает функцию, выдающую арифметическую последовательность.
func New(start, step int) func() int {
	_ = step

	return func() int {
		return start
	}
}
