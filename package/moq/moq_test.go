package moq_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/matryer/moq/package/moq"
)

func TestMoq(t *testing.T) {

	m, err := moq.New("../../example")
	if err != nil {
		t.Errorf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "PersonStore")
	if err != nil {
		t.Errorf("m.Mock: %s", err)
	}
	log.Println(buf.String())

}
