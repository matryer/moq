// +build go1.13

package nested

// See https://github.com/matryer/moq/issues/103#issuecomment-533772114
// Keeping a valid Go file for the go tool to find and detect the package.
//
// 'go test' fails if no files are present in the folder on Go v1.13:
//
// 	$ go test ./...
// 	...
// 	--- FAIL: TestModulesNestedPackage (0.05s)
// 		moq_modules_test.go:108: moq.New: Couldn't load mock package: go [list -e -json -compiled=true -test=false -export=false -deps=false -find=true --]: exit status 1: build .: cannot find module for path .
// 		FAIL
// 	...
