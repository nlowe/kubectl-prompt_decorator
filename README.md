# kubectl prompt-decorator

![](https://github.com/nlowe/kubectl-prompt_decorator/workflows/CI/badge.svg)  [![License](https://img.shields.io/badge/license-MIT-brightgreen)](./LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/nlowe/kubectl-prompt_decorator)](https://goreportcard.com/report/github.com/nlowe/kubectl-prompt_decorator)

Prints your current context and default namespace, useful for adding to your shell's `$PROMPT`

## Usage

On windows this requires Kubectl 1.16+ if you want to install the binary as a `kubectl` plugin.

Install the proper `kubectl-prompt_decorator` binary for your platform somewhere on your `$PATH`. You can now `kubectl prompt-decorator`!

```bash
$ which kubectl-prompt_decorator
~/.bin/kubectl-prompt_decorator
$ kubectl prompt-decorator
[☸ docker-desktop:default]
```

## Building

You need go 1.13+. Build the binary for your platform:

```bash
# linux / unix
$ go build -o kubectl-prompt_decorator main.go

# windows
$ go build -o kubectl-prompt_decorator.exe main.go
```

Then place `kubectl-prompt_decorator` on your `$PATH`.

## TODO

* Publish binaries as GitHub Releases
* Publish as a `krew` plugin
* Better theme support (currently only the default theme is available)

## License

This plugin is published under the MIT License. See [`./LICENSE`](./LICENSE) for details.
