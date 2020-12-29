package settings

import (
	"github.com/Netflix/go-env"
	"log"
)

type RegularKey struct {
	Iv  string `env:"CRYPT_IV"`
	Key string `env:"CRYPT_KEY"`
}

func NewRegularKey() *RegularKey {

	newKey := RegularKey{}
	if _, err := env.UnmarshalFromEnviron(&newKey); err != nil {
		log.Printf("Error on crypto vars - %v", err)
	}

	return &newKey
}
