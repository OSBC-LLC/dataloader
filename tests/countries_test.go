package tests

import (
	"strings"
	"testing"

	pkgstr "github.com/OSBC-LLC/dataloader/pkg/string"
)

func TestGetRandomCountry(t *testing.T) {
	for i := 0; i < 50; i++ {
		name, err := pkgstr.GetRandomCountry()
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
		if name == "" {
			t.Errorf("Expected a non-empty string, got an empty string")
		}
	}
}

func TestGetRandomCountryCode(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "alpha 2 code",
			args: args{l: 2},
			want: 2,
		},
		{
			name: "alpha 3 code",
			args: args{l: 3},
			want: 3,
		},
		{
			name: "negative test w/ 5",
			args: args{l: 5},
			want: 3,
		},
		{
			name: "negative test w/ 0",
			args: args{l: 0},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pkgstr.GetRandomCountryCode(tt.args.l)
			if err != nil {
				t.Errorf("GetRandomCountryCode(%v) = %v\n", tt.args.l, err)
			}
			if len(got) != tt.want {
				t.Errorf("GetRandomCountryCode(%v) = %v\n", tt.args.l, got)
			}
			if got == "" {
				t.Errorf("Expected a non-empty string, got an empty string")
			}
			if strings.ContainsAny(got, " \t\n") {
				t.Errorf("Expected a string without whitespace, got: %q", got)
			}
		})
	}
}
