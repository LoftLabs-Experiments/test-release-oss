package telemetry

var (
	// SyncerVersion holds the current version
	SyncerVersion = "dev"
	// telemetryPrivateKey holds the telemetry private key
	telemetryPrivateKey = ""
)

// SetVersion sets the version
func SetVersion(version string) {
	SyncerVersion = version
}

// GetVersion returns the current version
func GetVersion() string {
	return SyncerVersion
}
