package app_test

import (
	"testing"

	golcs "github.com/Trashed/go-lcs"
	"github.com/Trashed/go-lcs/internal/app"
	"github.com/Trashed/go-lcs/license"
)

func TestEmptyArgs(t *testing.T) {
	t.Parallel()

	app := app.New(license.NewLoader())
	err := app.Args()

	if err == nil {
		t.Fatalf("expected err \"%v\", got nil instead", golcs.ErrEmptyArgs)
	}
}
