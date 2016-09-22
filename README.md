![moq logo](moq-logo-small.png)

Interface mocking tool for go generate.

By [Mat Ryer](https://twitter.com/matryer) and [David Hernandez](https://github.com/dahernan), with ideas lovingly stolen from [Ernesto Jimenez](https://github.com/ernesto-jimenez).

### What is moq?

Moq is a tool that generates a struct from any interface. The struct can be used in test code as a mock of the interface.


### Installing

To start using moq, just run go get:
```
$ go get github.com/matryer/moq
```

### Usage

```
moq [flags] destination interface [interface2 [interface3 [...]]]
  -out string
    	output file (default stdout)
  -pkg string
    	package name (default will infer)
```

In a command line:

```
$ moq -out mocks_test.go . MyInterface
```

In code (for go generate):

```go
package my

//go:generate moq -out myinterface_moq_test.go . MyInterface

type MyInterface interface {
	Method1() error
	Method2(i int)
}
```

Then run `go generate` for your package.

### Install

```
go install github.com/matryer/moq
```

### How to use it

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

## Tips

* Keep mocked logic inside the test that is using it
* Only mock the fields you need - it will panic if a nil function gets called
* Use closured variables inside your test function to capture details about the calls to the methods
* Use `go:generate` to invoke the `moq` command

## License

The moq command (and all code) is licensed under the [MIT License](LICENSE).

The moq logo was created by [Chris Ryer](http://chrisryer.co.uk) and is licensed under the [Creative Commons Attribution 3.0 License](https://creativecommons.org/licenses/by/3.0/).
