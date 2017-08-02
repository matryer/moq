package moq

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestMoq(t *testing.T) {
	m, err := New("testpackages/example", "")
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
		"panic(\"moq: PersonStoreMock.CreateFunc is nil but PersonStore.Create was just called\")",
		"panic(\"moq: PersonStoreMock.GetFunc is nil but PersonStore.Get was just called\")",
		"lockPersonStoreMockGet.Lock()",
		"mock.calls.Get = append(mock.calls.Get, callInfo)",
		"lockPersonStoreMockGet.Unlock()",
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}

func TestMoqExplicitPackage(t *testing.T) {
	m, err := New("testpackages/example", "different")
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

// TestVeradicArguments tests to ensure variadic work as
// expected.
// see https://github.com/matryer/moq/issues/5
func TestVariadicArguments(t *testing.T) {
	m, err := New("testpackages/variadic", "")
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

func TestNothingToReturn(t *testing.T) {
	m, err := New("testpackages/example", "")
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
	m, err := New("testpackages/channels", "")
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
	m, err := New("testpackages/imports/two", "")
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

func TestTemplateFuncs(t *testing.T) {
	fn := templateFuncs["Exported"].(func(string) string)
	if fn("var") != "Var" {
		t.Errorf("exported didn't work: %s", fn("var"))
	}
}

func TestVendoredPackages(t *testing.T) {
	m, err := New("testpackages/vendoring/user", "")
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
	m, err := New(".", "moqtest_test")
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "Service")
	if err != nil {
		t.Errorf("mock error: %s", err)
	}
	s := buf.String()
	if !strings.Contains(s, `/moq/pkg/moq/testpackages/dotimport"`) {
		t.Error("contains invalid dot import")
	}
}

func TestEmptyInterface(t *testing.T) {
	m, err := New("testpackages/emptyinterface", "")
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
