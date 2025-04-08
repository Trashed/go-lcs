package app

import (
	"io"

	internal "github.com/Trashed/go-lcs/internal"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (c *App) Args(args ...string) error {
	if len(args) == 0 {
		return internal.ErrEmptyArgs
	}

	// TODO: Parse arguments and prepare the license

	return nil
}

func (c *App) Output(w io.WriteCloser) {

}
