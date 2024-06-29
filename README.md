# The Testing Package
The `testing.T` type offers methods to control test execution.

## Errors and Logs
- `t.Error*` mark a test as failed an continue
- `t.Log*` log during a test run
- `t.Fatal*` fast fail

## Parallel Tests
All tests calling the function `t.Parallel()` are executed in parallel.

## Skipping Tests
The `t.Skip()` method in conjunction with the `testing.Short()` method
allows to separate unit tests from integration tests or
short tests from long-running tests in general.

- `testing.Short()` marks a test as a short one
- `go test -test.short` runs only short tests

## Test Teardown and Cleanup
The `t.Cleanup()` function is executed after each (sub)test.

Calling the `t.Helper()` function in a helper function ensures 
that the reported line number points to the line in the test
that uses the helper function. See `divide_test.go` as an example.

The `t.TempDir()` function creates a temporary directory and
automatically deletes the directory when the test is completed.

# Coverage Tests
```
go test -cover
go test ./... -coverpkg./...  # use all packages (not only test ones)
go test -coverprofile=report.file  # for automatic check
go test cover -html=report.html  # human readable report
go test -cover -covermode set|count|atomic
```

- set: coverage is based on statements (default)
- count: count how many times a statement was run
- atomic: as count but for parallel tests

# Benchmark Tests
Use `perf/cmd` to analyse results with `benchstat` and `benchsave`.
```
go test -benchmem -bench
```

# Writing Fuzz Tests
The test function:
- accepts `testing.F` as the unique parameter;
- defines initial values (`seed corpus`), with the `f.Add()` method;
- defines a `fuzz target`.

See also: https://go.dev/doc/tutorial/fuzz
