package samename

import samename "github.com/matryer/moq/pkg/moq/testpackages/samenameimport/samenameimport"

// Example is used to test issues with packages, which import another package with the same name
type Example interface {
	Do(a samename.A) error
}
