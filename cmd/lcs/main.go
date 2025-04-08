package main

import (
	"log"
	"os"

	"github.com/Trashed/go-lcs/internal/app"
	"github.com/Trashed/go-lcs/internal/license"
)

func main() {
	cliApp := app.New(&license.Loader{})
	log.Fatalln(cliApp.Args(os.Args[1:]...))
	cliApp.Output(os.Stdout)
}
