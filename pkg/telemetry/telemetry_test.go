package telemetry

import (
	"testing"
)

func TestSetAndGetVersion(t *testing.T) {
	originalVersion := SyncerVersion
	defer func() {
		SyncerVersion = originalVersion
	}()
	
	testVersion := "v1.2.3"
	SetVersion(testVersion)
	
	if GetVersion() != testVersion {
		t.Errorf("Expected version to be %s, got %s", testVersion, GetVersion())
	}
}

func TestDefaultVersion(t *testing.T) {
	// Reset to default
	SyncerVersion = "dev"
	
	if GetVersion() != "dev" {
		t.Errorf("Expected default version to be 'dev', got %s", GetVersion())
	}
}
