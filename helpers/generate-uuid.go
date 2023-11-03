package helpers

import "github.com/google/uuid"

type GeneratorInterface interface {
	GenerateUUID()string
}

type generator struct{}

func NewGenerator() GeneratorInterface{
	return &generator{}
}

func (g generator) GenerateUUID() string {
	var result = uuid.NewString()

	return result
}