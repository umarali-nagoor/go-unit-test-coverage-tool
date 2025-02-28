# go-unit-test-coverage-tool
Its a sample repo to test go unit test coverage tool

Commands

* tree -L 2

Show test coverage

    $ go test -cover

Save Cover profile

    $ go test -coverprofile=cover.txt ./...

Analyse Cover profile

    $ go tool cover -html=cover.txt
    $ go tool cover -func=cover.txt

Common go test Flags
```
go test ./...	                       Runs tests in all subdirectories.
go test -v	                           Runs tests in verbose mode (shows all test names).
go test -cover	                       Shows test coverage percentage.
go test -coverprofile=coverage.out	   Saves coverage data to a file.
go test -run TestFunctionName	       Runs a specific test.
go test -bench .	                   Runs benchmarks.
go test -timeout 30s	               Sets a test timeout (default: 10m).
```
