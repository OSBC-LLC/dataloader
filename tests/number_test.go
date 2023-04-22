package tests

import (
	"testing"

	"github.com/OSBC-LLC/dataloader/pkg/number"
)

// TestGetRandomInt: Normal Go test that can be run with the following cmd:
// go test ./tests
func TestGetRandomInt(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
	}{
		{

			name: "0 -> 10",
			args: args{min: 0, max: 10},
		},
		{

			name: "-10 -> 0",
			args: args{min: -10, max: 0},
		},
		{

			name: "0 -> 0",
			args: args{min: 0, max: 0},
		},
		{

			name: "0 -> 1",
			args: args{min: 0, max: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := number.GetRandomInt(tt.args.min, tt.args.max); got > tt.args.max || got < tt.args.min {
				t.Errorf("GetRandomInt() = %v | min: %v  max: %v\n", got, tt.args.min, tt.args.max)
			}
		})
	}

}

// FuzzGetRandomInt: Fuzz test that can be executed with the following cmd:
// go test -fuzz=Fuzz -fuzztime=15s ./tests
func FuzzGetRandomInt(f *testing.F) {
	f.Add(0, 10)
	f.Fuzz(func(t *testing.T, min, max int) {
		if min > max {
			min, max = max, min
		} else if min == max {
			max = min + 1
		}
		got := number.GetRandomInt(min, max)
		if got > max || got < min {
			t.Errorf("GetRandomInt() = %v | min: %v  max: %v\n", got, min, max)
		}
	})
}

func TestGetRandomFloat(t *testing.T) {
	type args struct {
		min float64
		max float64
	}
	tests := []struct {
		name string
		args args
	}{
		{

			name: "0 -> 10",
			args: args{min: float64(0), max: float64(10)},
		},
		{

			name: "-10 -> 0",
			args: args{min: float64(-10), max: float64(0)},
		},
		{

			name: "0 -> 0",
			args: args{min: float64(0), max: float64(0)},
		},
		{

			name: "0 -> 1",
			args: args{min: float64(0), max: float64(1)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := number.GetRandomFloat(tt.args.min, tt.args.max); got > tt.args.max || got < tt.args.min {
				t.Errorf("GetRandomFloat() = %v | min: %v  max: %v\n", got, tt.args.min, tt.args.max)
			}
		})
	}

}

func FuzzGetRandomFloat(f *testing.F) {
	f.Add(float64(0), float64(10))
	f.Fuzz(func(t *testing.T, min, max float64) {
		if min > max {
			min, max = max, min
		} else if min == max {
			max = min + 1
		}
		got := number.GetRandomFloat(min, max)
		if got > max || got < min {
			t.Errorf("GetRandomFloat() = %v | min: %v  max: %v\n", got, min, max)
		}
	})
}
