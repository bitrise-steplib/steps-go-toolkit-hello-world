# steps-go-toolkit-hello-world

Bitrise "Hello world" step which uses the [Bitrise CLI](https://github.com/bitrise-io/bitrise.io) Go toolkit


## Technical notes for Go-Toolkit Steps

- `vendor` in your dependencies, e.g. with [godep](https://github.com/tools/godep),
  Bitrise CLI will not run `go get`!
- Go-Toolkit Steps are regular Go projects/packages, you can `go install`, `go build`, `go test`, ...
  just like you would with any other Go project, given that you do have a `GOPATH` prepared
  and the Step's code is at the right location (following Go's package location
  requirements).
    - Note: if all you want to do with the step is to run it through the [Bitrise CLI](https://github.com/bitrise-io/bitrise.io),
      then you don't have to clone the step's repository into a `GOPATH`,
      and actually `GOPATH` is not required at all. The Bitrise CLI will prepare
      everything for the step before it would execute it.
