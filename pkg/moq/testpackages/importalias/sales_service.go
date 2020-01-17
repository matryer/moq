package importalias

import (
	doucheclient "github.com/matryer/moq/pkg/moq/testpackages/importalias/douche/client"
	niceclient "github.com/matryer/moq/pkg/moq/testpackages/importalias/nice/client"
)

// SalesService is the interface which should be mocked by moq
type SalesService interface {
	PitchEasy(p niceclient.Person)
	PitchHard(p doucheclient.Person)
}
