
# dump error example

## About this example

This example shows how error dumps works.

## Prepare

```sh
git clone https://github.com/aileron-projects/go-debugger.git
cd go-debugger/examples/DUMPERR/
```

## Run without tag `dumperr`

Run the example without tag.

```sh
go run ./main.go 
```

It show the `VALIDATION_ERROR`.
It does not show the `HTTP_ERROR`.

```txt
2026-08-01 11:21:38 [DEBUGGER][DUMPERR] dump error always
  | Caller: Pkg:main File:main.go Func:main Line:36
  | ┌── Error: VALIDATION_ERROR: parameter must be number
  | (*main.MyError)(VALIDATION_ERROR: parameter must be number)
  | ┌── Stack Trace:
  | goroutine 1 [running]:
~~~ stack trace omitted ~~~
```

## Run with tag `dumperr`

Run the example with tag.

```sh
go run -tags dumperr ./main.go
```

It shows both error dumps of `VALIDATION_ERROR` and `HTTP_ERROR`.

```txt
2026-08-01 11:26:21 [DEBUGGER][DUMPERR] dump error
  | Caller: Pkg:main File:main.go Func:main Line:33
  | ┌── Error: HTTP_ERROR: request failed
  | (*main.MyError)(HTTP_ERROR: request failed)
  | ┌── Stack Trace:
  | goroutine 1 [running]:
~~~ stack trace omitted ~~~

2026-08-01 11:26:21 [DEBUGGER][DUMPERR] dump error always
  | Caller: Pkg:main File:main.go Func:main Line:36
  | ┌── Error: VALIDATION_ERROR: parameter must be number
  | (*main.MyError)(VALIDATION_ERROR: parameter must be number)
  | ┌── Stack Trace:
  | goroutine 1 [running]:
~~~ stack trace omitted ~~~
```
