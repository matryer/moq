package moq

// taken from https://github.com/ernesto-jimenez/gogen
// Copyright (c) 2015 Ernesto Jiménez

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type customImporter struct {
	source        string
	imported      map[string]*types.Package
	base          types.Importer
	skipTestFiles bool
}

func (i *customImporter) Import(path string) (*types.Package, error) {
	var err error
	if path == "" || path[0] == '.' {
		path, err = filepath.Abs(filepath.Clean(path))
		if err != nil {
			return nil, err
		}
		path = stripGopath(path)
	}
	if pkg, ok := i.imported[path]; ok {
		return pkg, nil
	}
	pkg, err := i.fsPkg(path)
	if err != nil {
		return nil, err
	}
	i.imported[path] = pkg
	return pkg, nil
}

func gopathDir(source, pkg string) (string, error) {
	// check vendor directory
	vendorPath, found := vendorPath(source, pkg)
	if found {
		return vendorPath, nil
	}
	for _, gopath := range gopaths() {
		absPath, err := filepath.Abs(path.Join(gopath, "src", pkg))
		if err != nil {
			return "", err
		}
		if dir, err := os.Stat(absPath); err == nil && dir.IsDir() {
			return absPath, nil
		}
	}
	return "", fmt.Errorf("%s not in $GOPATH or %s", pkg, path.Join(source, "vendor"))
}

func vendorPath(source, pkg string) (string, bool) {
	for {
		if isGopath(source) {
			return "", false
		}
		var err error
		source, err = filepath.Abs(source)
		if err != nil {
			return "", false
		}
		vendorPath, err := filepath.Abs(path.Join(source, "vendor", pkg))
		if err != nil {
			return "", false
		}
		if dir, err := os.Stat(vendorPath); err == nil && dir.IsDir() {
			return vendorPath, true
		}
		source = filepath.Dir(source)
	}
}

func removeGopath(p string) string {
	for _, gopath := range gopaths() {
		p = strings.Replace(p, path.Join(gopath, "src")+"/", "", 1)
	}
	return p
}

func gopaths() []string {
	return strings.Split(os.Getenv("GOPATH"), string(filepath.ListSeparator))
}

func isGopath(path string) bool {
	for _, p := range gopaths() {
		if p == path {
			return true
		}
	}
	return false
}

func (i *customImporter) fsPkg(pkg string) (*types.Package, error) {
	dir, err := gopathDir(i.source, pkg)
	if err != nil {
		return importOrErr(i.base, pkg, err)
	}

	dirFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		return importOrErr(i.base, pkg, err)
	}

	fset := token.NewFileSet()
	var files []*ast.File
	for _, fileInfo := range dirFiles {
		if fileInfo.IsDir() {
			continue
		}
		n := fileInfo.Name()
		if path.Ext(fileInfo.Name()) != ".go" {
			continue
		}
		if i.skipTestFiles && strings.Contains(fileInfo.Name(), "_test.go") {
			continue
		}
		file := path.Join(dir, n)
		src, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		f, err := parser.ParseFile(fset, file, src, 0)
		if err != nil {
			return nil, err
		}
		files = append(files, f)
	}
	conf := types.Config{
		Importer: i,
	}
	p, err := conf.Check(pkg, fset, files, nil)

	if err != nil {
		return importOrErr(i.base, pkg, err)
	}
	return p, nil
}

func importOrErr(base types.Importer, pkg string, err error) (*types.Package, error) {
	p, impErr := base.Import(pkg)
	if impErr != nil {
		return nil, err
	}
	return p, nil
}

// newImporter returns an importer that will try to import code from gopath before using go/importer.Default and skipping test files
func newImporter(source string) types.Importer {
	return &customImporter{
		source:        source,
		imported:      make(map[string]*types.Package),
		base:          importer.Default(),
		skipTestFiles: true,
	}
}

// // DefaultWithTestFiles same as Default but it parses test files too
// func DefaultWithTestFiles() types.Importer {
// 	return &customImporter{
// 		imported:      make(map[string]*types.Package),
// 		base:          importer.Default(),
// 		skipTestFiles: false,
// 	}
// }

// stripGopath teks the directory to a package and remove the gopath to get the
// canonical package name.
func stripGopath(p string) string {
	for _, gopath := range gopaths() {
		p = strings.TrimPrefix(p, path.Join(gopath, "src")+"/")
	}
	return p
}
