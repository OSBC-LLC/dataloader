package tests

import (
	"testing"

	"github.com/OSBC-LLC/dataloader/pkg/string"
)

func TestRandomMaleFirstName(t *testing.T) {
	for i := 0; i < 50; i++ {
		name, err := string.GetRandomMaleFirstName()
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
		if name == "" {
			t.Errorf("Expected a non-empty string, got an empty string")
		}
	}
}

func TestGetRandomFemaleFirstName(t *testing.T) {
	for i := 0; i < 50; i++ {
		name, err := string.GetRandomFemaleFirstName()
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
		if name == "" {
			t.Errorf("Expected a non-empty string, got an empty string")
		}
	}
}
