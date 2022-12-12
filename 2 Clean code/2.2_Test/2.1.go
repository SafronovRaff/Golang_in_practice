package main

import (
	"testing"
)

func Sum1(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func TestSum(t *testing.T) {
	if Sum1() != 0 {
		t.Errorf("Expected Sum() == 0")
	}
}

func TestSumOne1(t *testing.T) {
	if Sum1(1) != 1 {
		t.Errorf("Expected Sum(1) == 1")
	}
}

func TestSumPair1(t *testing.T) {
	if Sum1(1, 2) != 3 {
		t.Errorf("Expected Sum(1, 2) == 3")
	}
}

func TestSumMany1(t *testing.T) {
	if Sum1(1, 2, 3, 4, 5) != 15 {
		t.Errorf("Expected Sum(1, 2, 3, 4, 5) == 15")
	}
}
