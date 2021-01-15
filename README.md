# go-topic

## Topic List
* data types
* collection

## Go-Lang command
* `$ go mod init github.com/hendry-ref/golang-topic` to initialize go.mod file as deps management.
* `$ go install github.com/hendry-ref/golang-topic`
  * to install the `main.go` into `$GOBIN` directory
* `$ go env`
  * display all go environment variables
* `$ go env -w GOBIN=/somewhere/else/bin`
  * modify the go environment, in this case the `$GOBIN`
* `$ go env -u GOBIN`
  * unset the previously set variable from `-w` switch
* `$ go build`
  * saves the complied package in the local build cache
* `$ go run`
  * build and immediately run the compiled binary. only when the file is under `main` package with `func main() {...}`


* `$ go install, go build, go run` 
  * will automatically download all the dependencies and record in go.mod file
* `$ go clean -modcache`
  * Deps are automatically downloaded to `pkg/mod` under `$GOPATH` 
  * this command use to remove all downloaded modules
