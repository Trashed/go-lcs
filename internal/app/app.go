package app

import (
	"io"

	internal "github.com/Trashed/go-lcs/internal"
	license "github.com/Trashed/go-lcs/internal/license"
)

type App struct {
	License *internal.License

	loader *license.Loader
}

func New(loader *license.Loader) *App {
	return &App{
		loader: loader,
	}
}

func (app *App) Args(args ...string) error {
	if len(args) == 0 {
		return internal.ErrEmptyArgs
	}

	var err error
	app.License, err = app.loader.Fetch(args[0], args[1:]...)

	if err != nil {
		return err
	}

	return nil
}

func (app *App) Output(w io.WriteCloser) {

}
