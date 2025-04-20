package license

import (
	"embed"
	"io"
	"io/fs"

	golcs "github.com/Trashed/go-lcs"
)

//go:embed templates/*
var embedFS embed.FS

type loader struct {
	// TODO: Add reference to fs, this also enables better testing

	fs fs.FS
}

func NewLoader() *loader {
	return &loader{fs: embedFS}
}

// Fetch reads the license template from disk by name. Additional arguments could be used to replace values in the template.
func (l loader) Fetch(name string, additionalArgs ...string) (*golcs.License, error) {

	f, err := l.fs.Open("templates/" + name + ".lcs")
	if err != nil {
		return nil, err
	}

	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	license := golcs.License{
		Name:    name,
		Content: content,
	}

	return &license, nil
}
