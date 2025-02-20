# go-unit-test-coverage-tool
Its a sample repo to test go unit test coverage tool

Commands

* tree -L 2

Show test coverage

    $ go test -cover

Save Cover profile

    $ go test -coverprofile=cover.txt

Analyse Cover profile

    $ go tool cover -html=cover.txt
    $ go tool cover -func=cover.txt