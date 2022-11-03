package moq

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pmezard/go-difflib/difflib"
)

var update = flag.Bool("update", false, "Update golden files.")

func TestMoq(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/example"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "PersonStore")
	if err != nil {
		t.Errorf("m.Mock: %s", err)
	}
	s := buf.String()
	// assertions of things that should be mentioned
	strs := []string{
		"package example",
		"type PersonStoreMock struct",
		"CreateFunc func(ctx context.Context, person *Person, confirm bool) error",
		"GetFunc func(ctx context.Context, id string) (*Person, error)",
		"func (mock *PersonStoreMock) Create(ctx context.Context, person *Person, confirm bool) error",
		"func (mock *PersonStoreMock) Get(ctx context.Context, id string) (*Person, error)",
		"panic(\"PersonStoreMock.CreateFunc: method is nil but PersonStore.Create was just called\")",
		"panic(\"PersonStoreMock.GetFunc: method is nil but PersonStore.Get was just called\")",
		"mock.lockGet.Lock()",
		"mock.calls.Get = append(mock.calls.Get, callInfo)",
		"mock.lockGet.Unlock()",
		"// ID is the id argument value",
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}

func TestMoqWithStaticCheck(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/example"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "PersonStore")
	if err != nil {
		t.Errorf("m.Mock: %s", err)
	}
	s := buf.String()
	// assertions of things that should be mentioned
	strs := []string{
		"package example",
		"var _ PersonStore = &PersonStoreMock{}",
		"type PersonStoreMock struct",
		"CreateFunc func(ctx context.Context, person *Person, confirm bool) error",
		"GetFunc func(ctx context.Context, id string) (*Person, error)",
		"func (mock *PersonStoreMock) Create(ctx context.Context, person *Person, confirm bool) error",
		"func (mock *PersonStoreMock) Get(ctx context.Context, id string) (*Person, error)",
		"panic(\"PersonStoreMock.CreateFunc: method is nil but PersonStore.Create was just called\")",
		"panic(\"PersonStoreMock.GetFunc: method is nil but PersonStore.Get was just called\")",
		"mock.lockGet.Lock()",
		"mock.calls.Get = append(mock.calls.Get, callInfo)",
		"mock.lockGet.Unlock()",
		"// ID is the id argument value",
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}

func TestMoqWithAlias(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/example"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "PersonStore:AnotherPersonStoreMock")
	if err != nil {
		t.Errorf("m.Mock: %s", err)
	}
	s := buf.String()
	// assertions of things that should be mentioned
	strs := []string{
		"package example",
		"type AnotherPersonStoreMock struct",
		"CreateFunc func(ctx context.Context, person *Person, confirm bool) error",
		"GetFunc func(ctx context.Context, id string) (*Person, error)",
		"func (mock *AnotherPersonStoreMock) Create(ctx context.Context, person *Person, confirm bool) error",
		"func (mock *AnotherPersonStoreMock) Get(ctx context.Context, id string) (*Person, error)",
		"panic(\"AnotherPersonStoreMock.CreateFunc: method is nil but PersonStore.Create was just called\")",
		"panic(\"AnotherPersonStoreMock.GetFunc: method is nil but PersonStore.Get was just called\")",
		"mock.lockGet.Lock()",
		"mock.calls.Get = append(mock.calls.Get, callInfo)",
		"mock.lockGet.Unlock()",
		"// ID is the id argument value",
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}

func TestMoqExplicitPackage(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/example", PkgName: "different"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "PersonStore")
	if err != nil {
		t.Errorf("m.Mock: %s", err)
	}
	s := buf.String()
	// assertions of things that should be mentioned
	strs := []string{
		"package different",
		"type PersonStoreMock struct",
		"CreateFunc func(ctx context.Context, person *example.Person, confirm bool) error",
		"GetFunc func(ctx context.Context, id string) (*example.Person, error)",
		"func (mock *PersonStoreMock) Create(ctx context.Context, person *example.Person, confirm bool) error",
		"func (mock *PersonStoreMock) Get(ctx context.Context, id string) (*example.Person, error)",
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}

func TestMoqExplicitPackageWithStaticCheck(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/example", PkgName: "different"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "PersonStore")
	if err != nil {
		t.Errorf("m.Mock: %s", err)
	}
	s := buf.String()
	// assertions of things that should be mentioned
	strs := []string{
		"package different",
		"var _ example.PersonStore = &PersonStoreMock{}",
		"type PersonStoreMock struct",
		"CreateFunc func(ctx context.Context, person *example.Person, confirm bool) error",
		"GetFunc func(ctx context.Context, id string) (*example.Person, error)",
		"func (mock *PersonStoreMock) Create(ctx context.Context, person *example.Person, confirm bool) error",
		"func (mock *PersonStoreMock) Get(ctx context.Context, id string) (*example.Person, error)",
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}

func TestMoqSkipEnsure(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/example", PkgName: "different", SkipEnsure: true})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "PersonStore")
	if err != nil {
		t.Errorf("m.Mock: %s", err)
	}
	s := buf.String()
	// assertions of things that should be mentioned
	strs := []string{
		"package different",
		"type PersonStoreMock struct",
		"CreateFunc func(ctx context.Context, person *example.Person, confirm bool) error",
		"GetFunc func(ctx context.Context, id string) (*example.Person, error)",
		"func (mock *PersonStoreMock) Create(ctx context.Context, person *example.Person, confirm bool) error",
		"func (mock *PersonStoreMock) Get(ctx context.Context, id string) (*example.Person, error)",
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}

func TestMoqWithResets(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/example", PkgName: "different", WithResets: true})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "PersonStore")
	if err != nil {
		t.Errorf("m.Mock: %s", err)
	}
	s := buf.String()
	// assertions of things that should be mentioned
	strs := []string{
		"package different",
		"type PersonStoreMock struct",
		"CreateFunc func(ctx context.Context, person *example.Person, confirm bool) error",
		"GetFunc func(ctx context.Context, id string) (*example.Person, error)",
		"func (mock *PersonStoreMock) CreateResetCalls()",
		"func (mock *PersonStoreMock) ResetCalls()",
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}

func TestNotCreatingEmptyDirWhenPkgIsGiven(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/example", PkgName: "different"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "PersonStore")
	if err != nil {
		t.Errorf("m.Mock: %s", err)
	}
	s := buf.String()
	if len(s) == 0 {
		t.Fatalf("mock should be generated")
	}
	if _, err := os.Stat("testpackages/example/different"); !os.IsNotExist(err) {
		t.Fatalf("no empty dir should be created by moq")
	}
}

// TestVariadicArguments tests to ensure variadic work as
// expected.
// see https://github.com/rewardStyle/moq/issues/5
func TestVariadicArguments(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/variadic"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "Greeter")
	if err != nil {
		t.Errorf("m.Mock: %s", err)
	}
	s := buf.String()
	// assertions of things that should be mentioned
	strs := []string{
		"package variadic",
		"type GreeterMock struct",
		"GreetFunc func(ctx context.Context, names ...string) string",
		"return mock.GreetFunc(ctx, names...)",
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}

func TestNothingToReturn(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/example"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "PersonStore")
	if err != nil {
		t.Errorf("m.Mock: %s", err)
	}
	s := buf.String()
	if strings.Contains(s, `return mock.ClearCacheFunc(id)`) {
		t.Errorf("should not have return for items that have no return arguments")
	}
	// assertions of things that should be mentioned
	strs := []string{
		"mock.ClearCacheFunc(id)",
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}

func TestImports(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/imports/two"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "DoSomething")
	if err != nil {
		t.Errorf("m.Mock: %s", err)
	}
	s := buf.String()
	strs := []string{
		`	"sync"`,
		`	"github.com/rewardStyle/moq/pkg/moq/testpackages/imports/one"`,
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
		if len(strings.Split(s, str)) > 2 {
			t.Errorf("more than one: \"%s\"", str)
		}
	}
}

func TestMockGolden(t *testing.T) {
	cases := []struct {
		name       string
		cfg        Config
		interfaces []string
		goldenFile string
	}{
		{
			// Tests to ensure slice return data type works as expected.
			// See https://github.com/rewardStyle/moq/issues/124
			name:       "SliceResult",
			cfg:        Config{SrcDir: "testpackages/variadic"},
			interfaces: []string{"Echoer"},
			goldenFile: filepath.Join("testpackages/variadic", "echoer.golden.go"),
		},
		{
			// Tests generation of mock where a method on the interface uses a
			// blank identifier.
			// See https://github.com/rewardStyle/moq/issues/70
			name:       "BlankID",
			cfg:        Config{SrcDir: "testpackages/blankid"},
			interfaces: []string{"Swallower"},
			goldenFile: filepath.Join("testpackages/blankid", "swallower.golden.go"),
		},
		{
			name:       "ChannelNames",
			cfg:        Config{SrcDir: "testpackages/channels", StubImpl: true},
			interfaces: []string{"Queuer"},
			goldenFile: filepath.Join("testpackages/channels", "queuer_moq.golden.go"),
		},
		{
			// Tests generation of mock when the interface imports a different
			// package by the same name as it's own.
			// See https://github.com/rewardStyle/moq/issues/94
			name:       "PkgShadow",
			cfg:        Config{SrcDir: "testpackages/shadow/http", PkgName: "mock"},
			interfaces: []string{"Thing"},
			goldenFile: filepath.Join("testpackages/shadow/mock", "thing_moq.golden.go"),
		},
		{
			// Tests generation of mock when a method parameter shadows an
			// imported package name.
			name:       "ParamShadow",
			cfg:        Config{SrcDir: "testpackages/shadow"},
			interfaces: []string{"Shadower"},
			goldenFile: filepath.Join("testpackages/shadow", "shadower_moq.golden.go"),
		},
		{
			name:       "ImportAlias",
			cfg:        Config{SrcDir: "testpackages/importalias"},
			interfaces: []string{"MiddleMan"},
			goldenFile: filepath.Join("testpackages/importalias", "middleman_moq.golden.go"),
		},
		{
			// Tests conflict resolution for generated names of method
			// parameters.
			name:       "ParamNameConflict",
			cfg:        Config{SrcDir: "testpackages/paramconflict"},
			interfaces: []string{"Interface"},
			goldenFile: filepath.Join("testpackages/paramconflict", "iface_moq.golden.go"),
		},
		{
			// Tests generation of names for unnamed method parameters.
			name:       "GenerateParamNames",
			cfg:        Config{SrcDir: "testpackages/genparamname"},
			interfaces: []string{"Interface"},
			goldenFile: filepath.Join("testpackages/genparamname", "iface_moq.golden.go"),
		},
		{
			name:       "SyncImport",
			cfg:        Config{SrcDir: "testpackages/syncimport"},
			interfaces: []string{"Syncer"},
			goldenFile: filepath.Join("testpackages/syncimport", "syncer_moq.golden.go"),
		},
		{
			// Tests anonymous imports are not included in the generated mock.
			name:       "AnonymousImport",
			cfg:        Config{SrcDir: "testpackages/anonimport"},
			interfaces: []string{"Example"},
			goldenFile: filepath.Join("testpackages/anonimport", "iface_moq.golden.go"),
		},
		{
			name:       "ShadowTypes",
			cfg:        Config{SrcDir: "testpackages/shadowtypes"},
			interfaces: []string{"ShadowTypes"},
			goldenFile: filepath.Join("testpackages/shadowtypes", "shadowtypes_moq.golden.go"),
		},
		{
			name:       "Generics",
			cfg:        Config{SrcDir: "testpackages/generics"},
			interfaces: []string{"GenericStore1", "GenericStore2", "AliasStore"},
			goldenFile: filepath.Join("testpackages/generics", "generics_moq.golden.go"),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m, err := New(tc.cfg)
			if err != nil {
				t.Fatalf("moq.New: %s", err)
			}

			var buf bytes.Buffer
			if err = m.Mock(&buf, tc.interfaces...); err != nil {
				t.Errorf("m.Mock: %s", err)
				return
			}

			if err := matchGoldenFile(tc.goldenFile, buf.Bytes()); err != nil {
				t.Errorf("check golden file: %s", err)
			}
		})
	}
}

func TestFormatter(t *testing.T) {
	cases := []struct {
		name string
		conf Config
	}{
		{name: "gofmt", conf: Config{SrcDir: "testpackages/imports/two"}},
		{name: "goimports", conf: Config{SrcDir: "testpackages/imports/two", Formatter: "goimports"}},
		{name: "noop", conf: Config{SrcDir: "testpackages/imports/two", Formatter: "noop"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m, err := New(tc.conf)
			if err != nil {
				t.Fatalf("moq.New: %s", err)
			}
			var buf bytes.Buffer
			err = m.Mock(&buf, "DoSomething:"+tc.name+"Mock")
			if err != nil {
				t.Errorf("m.Mock: %s", err)
			}

			golden := filepath.Join("testpackages/imports/two", tc.name+".golden.go")
			if err := matchGoldenFile(golden, buf.Bytes()); err != nil {
				t.Errorf("check golden file: %s", err)
			}
		})
	}
}

func matchGoldenFile(goldenFile string, actual []byte) error {
	// To update golden files, run the following:
	// go test -v -run '^<Test-Name>$' github.com/rewardStyle/moq/pkg/moq -update
	if *update {
		if err := os.MkdirAll(filepath.Dir(goldenFile), 0o750); err != nil {
			return fmt.Errorf("create dir: %s", err)
		}
		if err := ioutil.WriteFile(goldenFile, actual, 0o600); err != nil {
			return fmt.Errorf("write: %s", err)
		}

		return nil
	}

	expected, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		return fmt.Errorf("read: %s: %s", goldenFile, err)
	}

	// Normalise newlines
	actual, expected = normalize(actual), normalize(expected)
	if !bytes.Equal(expected, actual) {
		diff, err := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
			A:        difflib.SplitLines(string(expected)),
			B:        difflib.SplitLines(string(actual)),
			FromFile: "Expected",
			ToFile:   "Actual",
			Context:  1,
		})
		if err != nil {
			return fmt.Errorf("diff: %s", err)
		}
		return fmt.Errorf("match: %s:\n%s", goldenFile, diff)
	}

	return nil
}

func TestVendoredPackages(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/vendoring/user"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "Service")
	if err != nil {
		t.Errorf("mock error: %s", err)
	}
	s := buf.String()
	// assertions of things that should be mentioned
	strs := []string{
		`"github.com/sudo-suhas/moq-test-pkgs/somerepo"`,
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}

func TestVendoredInterface(t *testing.T) {
	m, err := New(Config{
		SrcDir:  "testpackages/vendoring/vendor/github.com/sudo-suhas/moq-test-pkgs/somerepo",
		PkgName: "someother",
	})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "SomeService")
	if err != nil {
		t.Errorf("mock error: %s", err)
	}
	s := buf.String()
	// assertions of things that should be mentioned
	strs := []string{
		`"github.com/sudo-suhas/moq-test-pkgs/somerepo"`,
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
	incorrectImport := `"github.com/rewardStyle/moq/pkg/moq/testpackages/vendoring/vendor/github.com/sudo-suhas/moq-test-pkgs/somerepo"`
	if strings.Contains(s, incorrectImport) {
		t.Errorf("unexpected import: %s", incorrectImport)
	}
}

func TestVendoredBuildConstraints(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/buildconstraints/user"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "Service")
	if err != nil {
		t.Errorf("mock error: %s", err)
	}
	s := buf.String()
	// assertions of things that should be mentioned
	strs := []string{
		`"github.com/sudo-suhas/moq-test-pkgs/buildconstraints"`,
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}

// TestDotImports tests for https://github.com/rewardStyle/moq/issues/21.
func TestDotImports(t *testing.T) {
	preDir, err := os.Getwd()
	if err != nil {
		t.Errorf("Getwd: %s", err)
	}
	err = os.Chdir("testpackages/dotimport")
	if err != nil {
		t.Errorf("Chdir: %s", err)
	}
	defer func() {
		err := os.Chdir(preDir)
		if err != nil {
			t.Errorf("Chdir back: %s", err)
		}
	}()
	m, err := New(Config{SrcDir: ".", PkgName: "moqtest_test"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "Service")
	if err != nil {
		t.Errorf("mock error: %s", err)
	}
	s := buf.String()
	if strings.Contains(s, `"."`) {
		t.Error("contains invalid dot import")
	}
}

func TestEmptyInterface(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/emptyinterface"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "Empty")
	if err != nil {
		t.Errorf("mock error: %s", err)
	}
	s := buf.String()
	if strings.Contains(s, `"sync"`) {
		t.Error("contains sync import, although this package isn't used")
	}
}

func TestGoGenerateVendoredPackages(t *testing.T) {
	cmd := exec.Command("go", "generate", "./...")
	cmd.Dir = "testpackages/gogenvendoring"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Errorf("StdoutPipe: %s", err)
	}
	defer stdout.Close()
	err = cmd.Start()
	if err != nil {
		t.Errorf("Start: %s", err)
	}
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, stdout)
	err = cmd.Wait()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			t.Errorf("Wait: %s %s", exitErr, string(exitErr.Stderr))
		} else {
			t.Errorf("Wait: %s", err)
		}
	}
	s := buf.String()
	if strings.Contains(s, `vendor/`) {
		t.Error("contains vendor directory in import path")
	}
}

func TestImportedPackageWithSameName(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/samenameimport"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "Example")
	if err != nil {
		t.Errorf("mock error: %s", err)
	}
	s := buf.String()
	if !strings.Contains(s, `a samename.A`) {
		t.Error("missing samename.A to address the struct A from the external package samename")
	}
}

func TestParseError(t *testing.T) {
	_, err := New(Config{SrcDir: "testpackages/_parseerror/service"})
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}
	if !strings.Contains(err.Error(), `could not import github.com/matryer/notexist (invalid package name: "")`) {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestMockError(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/example"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	cases := []struct {
		name     string
		namePair string
		wantErr  string
	}{
		{
			name:     "TypeNotFound",
			namePair: "DoesNotExist",
			wantErr:  "interface not found: DoesNotExist",
		},
		{
			name:     "UnexpectedType",
			namePair: "Person",
			wantErr:  "Person (github.com/rewardStyle/moq/pkg/moq/testpackages/example.Person) is not an interface",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := m.Mock(ioutil.Discard, tc.namePair)
			if err == nil {
				t.Errorf("expected error but got nil")
				return
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Errorf("unexpected error: %s", err.Error())
			}
		})
	}
}

// normalize normalizes \r\n (windows) and \r (mac)
// into \n (unix)
func normalize(d []byte) []byte {
	// Source: https://www.programming-books.io/essential/go/normalize-newlines-1d3abcf6f17c4186bb9617fa14074e48
	// replace CR LF \r\n (windows) with LF \n (unix)
	d = bytes.Replace(d, []byte{13, 10}, []byte{10}, -1)
	// replace CF \r (mac) with LF \n (unix)
	d = bytes.Replace(d, []byte{13}, []byte{10}, -1)
	return d
}
