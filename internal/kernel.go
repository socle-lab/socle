package internal

import (
	"os"

	socle "github.com/socle-lab/core"
)

func Boot(module string) (*socle.Socle, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	// init socle
	socleApp := &socle.Socle{}
	err = socleApp.New(path, module)
	if err != nil {
		return nil, err
	}

	return socleApp, nil

}
