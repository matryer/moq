package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/matryer/moq/pkg/moq"
)

// Version is the command version, injected at build time.
var Version string = "dev"

type userFlags struct {
	outFile    string
	outDir     string
	pkgName    string
	formatter  string
	stubImpl   bool
	skipEnsure bool
	withResets bool
	remove     bool
	args       []string
}

func main() {
	var flags userFlags
	flag.StringVar(&flags.outFile, "out", "", "output file (default stdout)")
	flag.StringVar(&flags.outDir, "out-dir", "", "output dir (exclusive with -out)")
	flag.StringVar(&flags.pkgName, "pkg", "", "package name (default will infer)")
	flag.StringVar(&flags.formatter, "fmt", "", "go pretty-printer: gofmt, goimports or noop (default gofmt)")
	flag.BoolVar(&flags.stubImpl, "stub", false,
		"return zero values when no mock implementation is provided, do not panic")
	printVersion := flag.Bool("version", false, "show the version for moq")
	flag.BoolVar(&flags.skipEnsure, "skip-ensure", false,
		"suppress mock implementation check, avoid import cycle if mocks generated outside of the tested package")
	flag.BoolVar(&flags.remove, "rm", false, "first remove output file, if it exists")
	flag.BoolVar(&flags.withResets, "with-resets", false,
		"generate functions to facilitate resetting calls made to a mock")

	flag.Usage = func() {
		fmt.Println(`moq [flags] source-dir interface [interface2 [interface3 [...]]]`)
		flag.PrintDefaults()
		fmt.Println(`Specifying an alias for the mock is also supported with the format 'interface:alias'`)
		fmt.Println(`Ex: moq -pkg different . MyInterface:MyMock`)
	}

	flag.Parse()
	flags.args = flag.Args()

	if *printVersion {
		fmt.Printf("moq version %s\n", Version)
		os.Exit(0)
	}

	if err := run(flags); err != nil {
		fmt.Fprintln(os.Stderr, err)
		flag.Usage()
		os.Exit(1)
	}
}

func run(flags userFlags) error {
	if len(flags.args) < 2 {
		return errors.New("not enough arguments")
	}

	if flags.remove && flags.outFile != "" {
		if err := os.Remove(flags.outFile); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
		}
	}

	srcDir, args := flags.args[0], flags.args[1:]
	m, err := moq.New(moq.Config{
		SrcDir:     srcDir,
		PkgName:    flags.pkgName,
		Formatter:  flags.formatter,
		StubImpl:   flags.stubImpl,
		SkipEnsure: flags.skipEnsure,
		WithResets: flags.withResets,
	})
	if err != nil {
		return err
	}

	switch {
	case flags.outDir != "" && flags.outFile != "":
		return errors.New("use only one from -out and -out-dir arguments")
	case flags.outDir != "":
		return mockToDir(m, flags.outDir, args...)
	case flags.outFile != "":
		return mockToFile(m, flags.outFile, args...)
	default:
		// mock to stdout
		return m.Mock(os.Stdout, args...)
	}
}

func mockToDir(m *moq.Mocker, outDir string, args ...string) error {
	if err := os.MkdirAll(outDir, 0o750); err != nil {
		return err
	}

	var buf bytes.Buffer
	for _, arg := range args {
		if err := m.Mock(&buf, arg); err != nil {
			return err
		}

		filename := filepath.Join(outDir, m.FileMockName(arg))
		if err := ioutil.WriteFile(filename, buf.Bytes(), 0o600); err != nil {
			return err
		}

		buf.Reset()
	}

	return nil
}

func mockToFile(m *moq.Mocker, outFile string, args ...string) error {
	var buf bytes.Buffer
	var out io.Writer = &buf

	if err := m.Mock(out, args...); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(outFile), 0o750); err != nil {
		return err
	}

	return ioutil.WriteFile(outFile, buf.Bytes(), 0o600)
}
