package main

import (
	"os"

	"github.com/Trashed/go-lcs/internal/app"
)

func main() {
	cliApp := app.New()
	cliApp.Args(os.Args[1:]...)
	cliApp.Output(os.Stdout)
}
