package middleware

import (
	"testing"

	"github.com/ollama/ollama/envconfig"
)

func setTestHome(t *testing.T, home string) {
	t.Helper()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	envconfig.ReloadServerConfig()
}

// enableCloudForTest sets HOME to a clean temp dir and explicitly enables
// cloud features for the duration of the test. This fork is local-first, so
// cloud is disabled by default; tests that exercise cloud behavior must opt in
// via OLLAMA_NO_CLOUD=0 rather than relying on an empty/unset value.
func enableCloudForTest(t *testing.T) {
	t.Helper()
	t.Setenv("OLLAMA_NO_CLOUD", "0")
	setTestHome(t, t.TempDir())
}
