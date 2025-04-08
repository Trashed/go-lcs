package main

import (
	"log"
	"os"

	"github.com/Trashed/go-lcs/internal/app"
	"github.com/Trashed/go-lcs/internal/license"
)

func main() {
	cliApp := app.New(license.NewLoader())
	err := cliApp.Args(os.Args[1:]...)
	if err != nil {
		log.Fatalln(err)
	}
	cliApp.Output(os.Stdout)
}
