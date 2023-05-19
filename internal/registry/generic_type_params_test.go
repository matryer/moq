package registry

import "testing"

func TestPackageUtilities(t *testing.T) {
	testCases := []struct {
		Name             string
		Input            string
		ExpectedPkgName  string
		ExpectedPkgPath  string
		ExpectedTypeName string
	}{
		{
			Name:             "external package",
			Input:            "github.com/matryer/moq/fakepkg.MyType",
			ExpectedPkgName:  "fakepkg",
			ExpectedPkgPath:  "github.com/matryer/moq/fakepkg",
			ExpectedTypeName: "fakepkg.MyType",
		},
		{
			Name:             "internal package with ptr",
			Input:            "*os/fs.FileInfo",
			ExpectedPkgName:  "fs",
			ExpectedPkgPath:  "os/fs",
			ExpectedTypeName: "*fs.FileInfo",
		},
		{
			Name:             "internal package",
			Input:            "os/fs.FileInfo",
			ExpectedPkgName:  "fs",
			ExpectedPkgPath:  "os/fs",
			ExpectedTypeName: "fs.FileInfo",
		},
		{
			Name:             "external package with ptr",
			Input:            "*github.com/matryer/moq/fakepkg.MyType",
			ExpectedPkgName:  "fakepkg",
			ExpectedPkgPath:  "github.com/matryer/moq/fakepkg",
			ExpectedTypeName: "*fakepkg.MyType",
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			if res := getPkgName(test.Input); res != test.ExpectedPkgName {
				t.Fatalf("Got unexpected package name, Expected: '%s' Got: '%s'\n", test.ExpectedPkgName, res)
			}

			if res := getPackagePath(test.Input); res != test.ExpectedPkgPath {
				t.Fatalf("Got unexpected package path, Expected: '%s' Got: '%s'\n", test.ExpectedPkgPath, res)
			}

			if res := getName(test.Input); res != test.ExpectedTypeName {
				t.Fatalf("Got unexpected type name, Expected: '%s' Got: '%s'\n", test.ExpectedTypeName, res)
			}
		})
	}
}

func TestConstraintContainsPkg(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected bool
	}{
		{Input: "github.com/matryer/moq/pkg.SomeType", Expected: true},
		{Input: "*os/fs.T", Expected: true},
		{Input: "os.T", Expected: false},
		{Input: "os/fs", Expected: false},
		{Input: "os/fs.", Expected: false},
		{Input: "os/fs./os", Expected: false},
	}

	for _, test := range testCases {
		if res := ConstraintAppearsImported(test.Input); res != test.Expected {
			t.Fatalf("Got unexpected result, Expected: '%v' Got: '%v' for string: '%s'\n", test.Expected, res, test.Input)
		}
	}
}
