package app_test

import (
	"testing"

	internal "github.com/Trashed/go-lcs/internal"
	"github.com/Trashed/go-lcs/internal/app"
)

func TestEmptyArgs(t *testing.T) {
	t.Parallel()

	app := app.New()
	err := app.Args()

	if err == nil {
		t.Fatalf("expected err \"%v\", got nil instead", internal.ErrEmptyArgs)
	}
}
