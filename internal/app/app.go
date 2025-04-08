package app

import (
	"io"
	"log"

	internal "github.com/Trashed/go-lcs/internal"
)

type App struct {
	License *internal.License

	loader LicenseFetcher
}

type LicenseFetcher interface {
	Fetch(name string, additionalArgs ...string) (*internal.License, error)
}

func New(loader LicenseFetcher) *App {
	return &App{
		loader: loader,
	}
}

func (app *App) Args(args ...string) error {

	log.Printf("args: %+v\n", args)

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

func (app *App) Output(w io.WriteCloser) error {
	defer w.Close()

	_, err := w.Write(app.License.Content)
	if err != nil {
		return err
	}

	return nil
}
