package main

import (
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func Test1(t *testing.T) {
	t.Errorf("1 + 2 = %d; want 3", Add(1, 2))
}
