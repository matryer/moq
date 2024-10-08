package importalias

import (
	srcclient "github.com/matryer/moq/pkg/moq/testpackages/importalias/source/client"
	tgtclient "github.com/matryer/moq/pkg/moq/testpackages/importalias/target/client"
)

// MiddleMan is a test interface.
type MiddleMan interface {
	Connect(src srcclient.Client, tgt tgtclient.Client)
}
