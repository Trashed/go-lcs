package license_test

import (
	"strings"
	"testing"

	"github.com/Trashed/go-lcs/internal/license"
)

func TestFetch(t *testing.T) {
	t.Parallel()

	const LICENSE = "mit"

	l := license.NewLoader()

	license, err := l.Fetch(LICENSE)

	if err != nil {
		t.Fatalf("fetching license failed with an err: %v\n", err)
	}

	if license == nil || (license.Name != LICENSE && len(license.Content) == 0 || !strings.Contains(string(license.Content), "MIT License")) {
		t.Fatalf("failed to fetch the license \"%s\"", LICENSE)
	}
}
