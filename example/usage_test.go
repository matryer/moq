package example

import (
	"context"
	"reflect"
	"testing"
)

func TestService_GetPerson(t *testing.T) {
	ctx := context.Background()
	personStore := &PersonStoreMock{
		GetFunc: func(ctx context.Context, id string) (*Person, error) {
			return &Person{
				ID:   "1",
				Name: "John Doe",
			}, nil
		},
	}
	s := NewService(personStore)
	// When
	s.GetPerson(ctx, "1")
	// Then
	actual := personStore.GetCalls()
	expected := []PersonStoreMockGetCalls{
		{
			Ctx: ctx,
			ID:  "1",
		},
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}
