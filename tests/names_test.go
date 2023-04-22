package tests

import (
	"strings"
	"testing"

	pkgstr "github.com/OSBC-LLC/dataloader/pkg/string"
)

func TestRandomMaleFirstName(t *testing.T) {
	for i := 0; i < 50; i++ {
		name, err := pkgstr.GetRandomMaleFirstName()
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
		if name == "" {
			t.Errorf("Expected a non-empty string, got an empty string")
		}
		if strings.ContainsAny(name, " \t\n") {
			t.Errorf("Expected a string without whitespace, got: %q", name)
		}
	}
}

func TestGetRandomFemaleFirstName(t *testing.T) {
	for i := 0; i < 50; i++ {
		name, err := pkgstr.GetRandomFemaleFirstName()
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
		if name == "" {
			t.Errorf("Expected a non-empty string, got an empty string")
		}
		if strings.ContainsAny(name, " \t\n") {
			t.Errorf("Expected a string without whitespace, got: %q", name)
		}
	}
}

func TestGetRandomMaleOrFemaleFirstName(t *testing.T) {
	for i := 0; i < 50; i++ {
		name, err := pkgstr.GetRandomMaleOrFemaleFirstName()
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
		if name == "" {
			t.Errorf("Expected a non-empty string, got an empty string")
		}
		if strings.ContainsAny(name, " \t\n") {
			t.Errorf("Expected a string without whitespace, got: %q", name)
		}
	}
}

func TestGetRandomLastName(t *testing.T) {
	for i := 0; i < 50; i++ {
		name, err := pkgstr.GetRandomLastName()
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
		if name == "" {
			t.Errorf("Expected a non-empty string, got an empty string")
		}
		if strings.ContainsAny(name, " \t\n") {
			t.Errorf("Expected a string without whitespace, got: %q", name)
		}
	}
}
