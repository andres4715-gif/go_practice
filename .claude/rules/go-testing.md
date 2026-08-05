# Rule: Go testing with testify

Applies when writing or editing code and tests in this repository.

## Project style

- Code comments, variable names, and explanations are in **Spanish**, simple and didactic.
- Run `gofmt -w .` before finishing. Code must pass `go build ./...` and `go vet ./...`.

## Test framework

- Use the standard `testing` package as the base and **testify** (`github.com/stretchr/testify`) for assertions.
- Prefer `require` when a failure should stop the test (e.g. an `err` that invalidates the rest) and `assert` when you want to keep checking.

```go
import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestCarDataConstructor(t *testing.T) {
    c, err := carDataConstructor("3", "Gran turing", "Mazda", 2017, 30000)
    require.NoError(t, err)                 // if it fails, stop here
    assert.Equal(t, "Mazda", c.Brand)       // keep checking even on failure
    assert.Equal(t, "3 Gran turing", c.Model)
}
```

## Conventions

- Test file: `<file>_test.go`, in the **same package** as the code under test.
- Test functions: `func TestXxx(t *testing.T)`.
- Prefer **table-driven tests** for multiple cases:

```go
func TestCarDataConstructor_Invalid(t *testing.T) {
    cases := []struct {
        name  string
        model string
        price int64
    }{
        {"no model", "", 30000},
        {"negative price", "Corolla", -1},
    }
    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            _, err := carDataConstructor(c.model, "v", "Toyota", 2020, c.price)
            require.Error(t, err)   // here we only care that it fails
        })
    }
}
```

## Running the tests

```bash
go test ./...                              # whole module
go test ./basic_learning/black_hole/...    # a single package/folder
go test -run TestCarDataConstructor ./...  # a specific test
go test -v ./...                           # verbose output
```
