package utils

import (
	"log"

	"github.com/jaevor/go-nanoid"
	"github.com/mshore-dev/imagebucket/config"
)

var generate func() string

func Setup() {
	var err error

	generate, err = nanoid.Standard(config.Config.Uploads.IDLength)
	if err != nil {
		log.Fatalf("failed to initialize nanoid: %v\n", err)
	}
}

func GenerateID() string {
	return generate()
}
