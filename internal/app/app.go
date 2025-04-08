package app

import "io"

type cliApp struct {
}

func New() *cliApp {
	return &cliApp{}
}

func (c *cliApp) Args(args ...string) {

}

func (c *cliApp) Output(w io.WriteCloser) {

}
