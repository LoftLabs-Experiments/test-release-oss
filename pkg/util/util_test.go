package util

import (
	"strings"
	"testing"
	"time"
)

func TestGenerateID(t *testing.T) {
	id := GenerateID()
	
	if len(id) == 0 {
		t.Error("Expected non-empty ID")
	}
	
	// UUIDs should be 36 characters long
	if len(id) != 36 {
		t.Errorf("Expected ID length to be 36, got %d", len(id))
	}
}

func TestFormatTimestamp(t *testing.T) {
	testTime := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	formatted := FormatTimestamp(testTime)
	
	expected := "2023-01-01 12:00:00"
	if formatted != expected {
		t.Errorf("Expected %s, got %s", expected, formatted)
	}
}

func TestSanitizeName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test name", "test-name"},
		{"Test Name", "test-name"},
		{"test-name", "test-name"},
		{"TEST_NAME", "test_name"},
	}
	
	for _, test := range tests {
		result := SanitizeName(test.input)
		if result != test.expected {
			t.Errorf("SanitizeName(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestBuildInfo(t *testing.T) {
	version := "v1.0.0"
	buildTime := "2023-01-01"
	
	info := BuildInfo(version, buildTime)
	
	if !strings.Contains(info, version) {
		t.Errorf("Expected build info to contain version %s", version)
	}
	
	if !strings.Contains(info, buildTime) {
		t.Errorf("Expected build info to contain build time %s", buildTime)
	}
}
