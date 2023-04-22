package tests

import (
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
