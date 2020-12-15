package syncimport

import (
	stdsync "sync"

	"github.com/matryer/moq/pkg/moq/testpackages/syncimport/sync"
)

type Syncer interface {
	Blah(s sync.Thing, wg *stdsync.WaitGroup)
}
