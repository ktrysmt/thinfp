package testing

import (
	"reflect"
	"testing"

	fp "github.com/ktrysmt/thinfp"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		f        func(int) string
		expected []string
	}{
		{
			name: "int to string",
			list: []int{1, 2, 3},
			f:    func(i int) string { return "Num" },
			expected: []string{
				"Num", "Num", "Num",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := fp.Map(tt.list, tt.f)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Map() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		list      []int
		predicate func(int) bool
		expected  []int
	}{
		{
			name: "even numbers",
			list: []int{1, 2, 3, 4},
			predicate: func(i int) bool {
				return i%2 == 0
			},
			expected: []int{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := fp.Filter(tt.list, tt.predicate)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Filter() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name        string
		list        []int
		initial     int
		accumulator func(int, int) int
		expected    int
	}{
		{
			name:    "sum of numbers",
			list:    []int{1, 2, 3},
			initial: 0,
			accumulator: func(acc int, val int) int {
				return acc + val
			},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := fp.Reduce(tt.list, tt.initial, tt.accumulator)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Reduce() = %v, want %v", actual, tt.expected)
			}
		})
	}
}
