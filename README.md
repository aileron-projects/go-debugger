# go-debugger

**Go debugging utilities.**

<div align="center">

[![GoDoc](https://godoc.org/github.com/aileron-projects/go-debugger?status.svg)](http://godoc.org/github.com/aileron-projects/go-debugger)
[![Test](https://github.com/aileron-projects/go-debugger/actions/workflows/test.yaml/badge.svg?branch=main)](https://github.com/aileron-projects/go-debugger/actions/workflows/test.yaml?query=branch%3Amain)
[![License](https://img.shields.io/badge/License-Apache%202.0-yellow.svg)](./LICENSE)

[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/aileron-projects/go-debugger)
[![OpenSourceInsight](https://badgen.net/badge/open%2Fsource%2F/insight/cyan)](https://deps.dev/go/github.com%2Faileron-projects%2Fgo-debugger)
[![OSS Insight](https://badgen.net/badge/OSS/Insight/orange)](https://ossinsight.io/analyze/aileron-projects/go-debugger)

</div>

## Features

- Dump objects. On/off dumping with tag `-tags dump`.
- Dump errors. On/off dumping with tag `-tags dumperr`.
- Output destination can be cahnged to stdout, stderr, discard and files.
  - Object dumps with environmental variable `GO_DEBUGGER_DUMP_OUTPUT`.
  - Error dumps with environemtnal variable `GO_DEBUGGER_DUMP_OUTPUT`.
- Easy stack frames manipulation.

## Tested Environments

Operating System:

- `Linux`: [ubuntu-latest](https://github.com/actions/runner-images)
- `Windows`: [windows-latest](https://github.com/actions/runner-images)
- `macOS`: [macos-latest](https://github.com/actions/runner-images)

Architecture (Using QEMU on linux):

- x86: `amd64`, `386`
- arm: `arm/v5`, `arm/v6`, `arm/v7`, `arm64`
- risc: `riscv64`, `loong64`
- ppc: `ppc64`, `ppc64le`
- mips: `mips`, `mips64`, `mips64le`, `mipsle`
- ibm: `s390x`

## Release Cycle

- Releases are made as needed.
- [Semantic Versioning](https://semver.org/) `vX.Y.Z` is used.

## License

[Apache-2.0](LICENSE)

## Usage

### Object dumps

2 Options to output object dumps.
The function `Dump` works with build tag so it can work only when debugging.

- `Dump` and `DumpTo` works with build tag `-tags dump`.
- `DumpAlways` and `DumpAlwaysTo` works without any build tag.

By default, `Dump` and `DumpAlways` output dumps to stdout.
It can be changed by the environment variable `GO_DEBUGGER_DUMP_OUTPUT`.
`GO_DEBUGGER_DUMP_OUTPUT` can take one of these values.

- `stdout`: standard output
- `stderr`: standard error output
- `file`: file output (files created in system's temp directory)
- `discard`: discard all output

```go
val := struct {
    foo int
    bar string
}{
    foo: 123,
    bar: "bar",
}

debugger.Dump("this is an example.", val)

// Example output:
// 
// 2026-08-01 11:38:47 [DUMP] this is an example.
//   | Caller: Pkg:github.com/aileron-projects/go-debugger_test File:example_test.go Func:Example Line:42
//   | ┌── args[0]
//   | (struct { foo int; bar string }) {
//   |  foo: (int) 123,
//   |  bar: (string) (len=3) "bar"
//   | }
```

### Error dumps

Error dumps works just like object dumps.
Use following function for error dumps.

- `DumpErr` and `DumpErrTo` works with build tag `-tags dumperr`.
- `DumpErrAlways` and `DumpErrAlwaysTo` works without any build tag.

And use `GO_DEBUGGER_DUMPERR_OUTPUT` to change dump output destination.

```go
debugger.DumpErr("this is an example.", io.EOF)

// Example output:
// 
// 2026-08-01 11:49:40 [DEBUGGER][DUMPERR] this is an example.
//   | Caller: Pkg:github.com/aileron-projects/go-debugger_test File:example_test.go Func:Example Line:36
//   | ┌── Error: EOF
//   | (*errors.errorString)(EOF)
//   | ┌── Stack Trace:
//   | goroutine 1 [running]:
//   | github.com/aileron-projects/go-debugger.dumpErr({0x7ff6afbde608, 0x1615158b2058}, {0x7ff6afa4598e, 0x13}, {0x1615158d7a48, 0x1, 0x1615158c63f0?})
// ~~ stack trace omitted ~~~
```

## Build Tags

- `dump`: enables object dump output to work.
- `dumperr`: enables error dump output to work.

## Enviromental Variables

- `GO_DEBUGGER_DUMP_OUTPUT`: optionaly specifies object dump output destination. `stdout`, `stderr`, `discard` or `file`.
- `GO_DEBUGGER_DUMP_PACKAGES`: optionaly filters go packages to output object dumps.
- `GO_DEBUGGER_DUMPERR_OUTPUT`: optionaly specifies error dump output destination. `stdout`, `stderr`, `discard` or `file`.
- `GO_DEBUGGER_DUMPERR_PACKAGES`: optionaly filters go packages to output error dumps.
