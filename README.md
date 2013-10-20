# Golang Web Euler Test

Learning golang by implementing net/http, project euler, and asserts/testing.

Also tests git branching/rebasing.

## Server

* `go build server.go`
* `go run server.go`
* Go to `http://localhost:8080/euler/` in a browser.

### Routes

* `http://localhost:8080/`
* `http://localhost:8080/view/test`
* `http://localhost:8080/euler/[problem #]/`
* `http://localhost:8080/euler/[problem #]/[variable num from problem, e.g. 1000 from #1]/`

## Tests

Uses asserts as well.

`go test` to run tests from the `golang_web_euler` directory.

## References

See:
* http://golang.org/doc/articles/wiki/
* http://golangtutorials.blogspot.com/2011/10/gotest-unit-testing-and-benchmarking-go.html
* https://github.com/coocood/assrt
* https://github.com/eddie/goalg
* http://projecteuler.net/problems
