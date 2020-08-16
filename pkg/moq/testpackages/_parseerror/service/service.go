package service

import "github.com/matryer/notexist"

// Service does something good with computers.
type Service interface {
	DoSomething(notexist.SomeType) error
}
