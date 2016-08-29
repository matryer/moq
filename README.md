# moq

Interface mocking tool for go generate.

By [Mat Ryer](https://twitter.com/matryer) and [David Hernandez](https://github.com/dahernan), with ideas lovingly stolen from [Ernesto Jimenez](https://github.com/ernesto-jimenez).

## Usage

```
moq InterfaceName -out mocks_test.go
```

## Install

```
go install github.com/matryer/moq
```

## How to use it

Mocking interfaces is a nice way to write unit tests where you can easily control the behaviour of the mocked object.

Moq creates a struct that has a function field for each method, which you can declare in your test code.

This this example, Moq generated the `EmailSenderMock` type:

```go
func TestCompleteSignup(t *testing.T) {

	called := false
	var sentTo string 

	mockedEmailSender = &EmailSenderMock{
		SendFunc: func(to, subject, body string) error {
			called = true
			sentTo = to
			return nil
		},
	}

	CompleteSignUp("me@email.com", mockedEmailSender)

	if called == false {
		t.Error("Sender.Send expected")
	}
	if sentTo != "me@email.com" {
		t.Errorf("unexpected recipient: %s", sentTo)
	}

}

func CompleteSignUp(to string, sender EmailSender) {
	// TODO: this
}
```

The mocked structure implements the interface, where each method calls the associated function field.