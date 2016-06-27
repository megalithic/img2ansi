//build: unix

package main

import (
	"os"

	"github.com/e11it/gosshold/tree/master/ssh/terminal"
)

func getTermDim() (w, h int, err error) {
	return terminal.GetSize(int(os.Stdout.Fd()))
}
