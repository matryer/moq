package moq

import (
	"bytes"
	"flag"
	"fmt"
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
	var strs = []string{
		"package example",
		"type PersonStoreMock struct",
		"CreateFunc func(ctx context.Context, person *Person, confirm bool) error",
		"GetFunc func(ctx context.Context, id string) (*Person, error)",
		"func (mock *PersonStoreMock) Create(ctx context.Context, person *Person, confirm bool) error",
		"func (mock *PersonStoreMock) Get(ctx context.Context, id string) (*Person, error)",
		"panic(\"PersonStoreMock.CreateFunc: method is nil but PersonStore.Create was just called\")",
		"panic(\"PersonStoreMock.GetFunc: method is nil but PersonStore.Get was just called\")",
		"lockPersonStoreMockGet.Lock()",
		"mock.calls.Get = append(mock.calls.Get, callInfo)",
		"lockPersonStoreMockGet.Unlock()",
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
	var strs = []string{
		"package example",
		"var _ PersonStore = &PersonStoreMock{}",
		"type PersonStoreMock struct",
		"CreateFunc func(ctx context.Context, person *Person, confirm bool) error",
		"GetFunc func(ctx context.Context, id string) (*Person, error)",
		"func (mock *PersonStoreMock) Create(ctx context.Context, person *Person, confirm bool) error",
		"func (mock *PersonStoreMock) Get(ctx context.Context, id string) (*Person, error)",
		"panic(\"PersonStoreMock.CreateFunc: method is nil but PersonStore.Create was just called\")",
		"panic(\"PersonStoreMock.GetFunc: method is nil but PersonStore.Get was just called\")",
		"lockPersonStoreMockGet.Lock()",
		"mock.calls.Get = append(mock.calls.Get, callInfo)",
		"lockPersonStoreMockGet.Unlock()",
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
	var strs = []string{
		"package example",
		"type AnotherPersonStoreMock struct",
		"CreateFunc func(ctx context.Context, person *Person, confirm bool) error",
		"GetFunc func(ctx context.Context, id string) (*Person, error)",
		"func (mock *AnotherPersonStoreMock) Create(ctx context.Context, person *Person, confirm bool) error",
		"func (mock *AnotherPersonStoreMock) Get(ctx context.Context, id string) (*Person, error)",
		"panic(\"AnotherPersonStoreMock.CreateFunc: method is nil but PersonStore.Create was just called\")",
		"panic(\"AnotherPersonStoreMock.GetFunc: method is nil but PersonStore.Get was just called\")",
		"lockAnotherPersonStoreMockGet.Lock()",
		"mock.calls.Get = append(mock.calls.Get, callInfo)",
		"lockAnotherPersonStoreMockGet.Unlock()",
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
	var strs = []string{
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
	var strs = []string{
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
// see https://github.com/matryer/moq/issues/5
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
	var strs = []string{
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

// TestSliceResult tests to ensure slice return data type works as
// expected.
// see https://github.com/matryer/moq/issues/124
func TestSliceResult(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/variadic"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}

	var buf bytes.Buffer
	if err = m.Mock(&buf, "Echoer"); err != nil {
		t.Errorf("m.Mock: %s", err)
	}

	golden := filepath.Join("testpackages/variadic/testdata", "echoer.golden.go")
	if err := matchGoldenFile(golden, buf.Bytes()); err != nil {
		t.Errorf("check golden file: %s", err)
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
	var strs = []string{
		"mock.ClearCacheFunc(id)",
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}

func TestChannelNames(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/channels"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "Queuer")
	if err != nil {
		t.Errorf("m.Mock: %s", err)
	}
	s := buf.String()
	var strs = []string{
		"func (mock *QueuerMock) Sub(topic string) (<-chan Queue, error)",
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
	var strs = []string{
		`	"sync"`,
		`	"github.com/matryer/moq/pkg/moq/testpackages/imports/one"`,
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

func TestFormatter(t *testing.T) {
	cases := []struct {
		name string
		conf Config
	}{
		{name: "gofmt", conf: Config{SrcDir: "testpackages/imports/two"}},
		{name: "goimports", conf: Config{SrcDir: "testpackages/imports/two", Formatter: "goimports"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m, err := New(tc.conf)
			if err != nil {
				t.Fatalf("moq.New: %s", err)
			}
			var buf bytes.Buffer
			err = m.Mock(&buf, "DoSomething")
			if err != nil {
				t.Errorf("m.Mock: %s", err)
			}

			golden := filepath.Join("testpackages/imports/testdata", tc.name+".golden.go")
			if err := matchGoldenFile(golden, buf.Bytes()); err != nil {
				t.Errorf("check golden file: %s", err)
			}
		})
	}
}

func matchGoldenFile(goldenFile string, actual []byte) error {
	// To update golden files, run the following:
	// go test -v -run '^<Test-Name>$' github.com/matryer/moq/pkg/moq -update
	if *update {
		if err := ioutil.WriteFile(goldenFile, actual, 0644); err != nil {
			return fmt.Errorf("write: %s: %s", goldenFile, err)
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

func TestTemplateFuncs(t *testing.T) {
	fn := templateFuncs["Exported"].(func(string) string)
	if fn("var") != "Var" {
		t.Errorf("exported didn't work: %s", fn("var"))
	}
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
	var strs = []string{
		`"github.com/matryer/somerepo"`,
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}

func TestVendoredInterface(t *testing.T) {
	m, err := New(Config{
		SrcDir:  "testpackages/vendoring/vendor/github.com/matryer/somerepo",
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
	var strs = []string{
		`"github.com/matryer/somerepo"`,
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
	incorrectImport := `"github.com/matryer/moq/pkg/moq/testpackages/vendoring/vendor/github.com/matryer/somerepo"`
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
	var strs = []string{
		`"github.com/matryer/buildconstraints"`,
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}

// TestDotImports tests for https://github.com/matryer/moq/issues/21.
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

func helperExec(t *testing.T, args []string, dir string) []byte {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Exec: %s; Output:\n%s\n", err, output)
	}
	return output
}

func TestGoGenerateVendoredPackages(t *testing.T) {
	output := helperExec(t, []string{"go", "generate", "./..."}, "testpackages/gogenvendoring")
	if bytes.Contains(output, []byte(`vendor/`)) {
		t.Error("contains vendor directory in import path")
	}
}

func TestOverrlappingImports(t *testing.T) {
	m, err := New(Config{
		SrcDir:  "testpackages/overlappingimports",
		PkgName: "template",
	})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "Example", "Example2")
	if err != nil {
		t.Errorf("mock error: %s", err)
	}
	s := buf.String()
	// assertions of things that should be mentioned
	var strs = []string{
		`template1 "github.com/matryer/moq/pkg/moq/testpackages/overlappingimports"`,
		`"github.com/matryer/moq/pkg/moq/testpackages/overlappingimports/one"`,
		`one1 "github.com/matryer/moq/pkg/moq/testpackages/overlappingimports/one/one"`,
		`"html/template"`,
		`custom "text/template"`,
		`HTML() template.HTML`,
		`Text() custom.Template`,
		`Self() template1.Example`,
		`MultiFile() one.Example`,
		`Example2MultiFile() one1.Example`,
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
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

func TestParamShadowsImport(t *testing.T) {
	m, err := New(Config{SrcDir: "testpackages/shadow", PkgName: "gen"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}

	genDir := "./testpackages/shadow/gen"
	err = os.MkdirAll(genDir, 0777)
	if err != nil {
		t.Fatalf("mkdir: %s", err)
	}
	defer os.Remove(genDir)

	f, err := os.Create(filepath.Join(genDir, "shadowmock.go"))
	if err != nil {
		t.Fatalf("failed to create shadowmock: %s", err)
	}
	defer f.Close()
	defer os.Remove(f.Name())

	err = m.Mock(f, "Shadower")
	if err != nil {
		t.Fatalf("mock error: %s", err)
	}

	helperExec(t, []string{"go", "vet", genDir}, "")
}
