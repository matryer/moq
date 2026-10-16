package registry

import (
	"go/types"
	"os"
	"path/filepath"
	"testing"
)

func TestDestinationPackageLookup(t *testing.T) {
	srcDir := t.TempDir()
	for name, contents := range map[string]string{
		"go.mod":        "module example.com/source\n\ngo 1.26\n",
		"source.go":     "package source\n",
		"mocks/mock.go": "package mocks\n",
	} {
		filename := filepath.Join(srcDir, name)
		if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filename, []byte(contents), 0644); err != nil {
			t.Fatal(err)
		}
	}

	for _, tt := range []struct {
		name string
		want string
	}{
		{"", "example.com/source"},
		{"source", "example.com/source"},
		{"source_test", ""},
		{"mocks", "example.com/source/mocks"},
		{"missing", ""},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r, err := New(srcDir, tt.name)
			if err != nil {
				t.Fatal(err)
			}
			if r.moqPkgPath != tt.want {
				t.Fatalf("destination path = %q, want %q", r.moqPkgPath, tt.want)
			}
			if tt.want != "" && r.AddImport(types.NewPackage(tt.want, tt.name)) != nil {
				t.Fatal("destination package must not import itself")
			}
		})
	}
}
