package uuid

import "github.com/google/uuid"

type GenerateFn func() string

func Generate() string {
	return uuid.New().String()
}
