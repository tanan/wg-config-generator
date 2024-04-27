package util

import "fmt"

type KeyError struct {
	kind  KeyType
	inner error
}

type KeyType string

const (
	Private   = KeyType("Private")
	Public    = KeyType("Public")
	PreShared = KeyType("PreShared")
)

var (
	ErrPrivateKey   = NewKeyError(Private, nil)
	ErrPublicKey    = NewKeyError(Public, nil)
	ErrPreSharedKey = NewKeyError(PreShared, nil)
)

func NewKeyError(kind KeyType, err error) KeyError {
	return KeyError{
		kind:  kind,
		inner: err,
	}
}

func (err KeyError) Error() string {
	return fmt.Sprintf("failed to generate key. type: %s, error: %s", err.kind, err.inner.Error())
}

func (err KeyError) Is(target error) bool {
	t, ok := target.(KeyError)
	return ok && err.kind == t.kind
}

func (e KeyError) Unwrap() error {
	return e.inner
}
