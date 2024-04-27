package model

import "fmt"

type OutputType string

const (
	Text  = OutputType("text")
	Email = OutputType("email")
)

func (o OutputType) Valid() error {
	switch o {
	case Text, Email:
		return nil
	}
	return fmt.Errorf("invalid output type: %s", o)
}
