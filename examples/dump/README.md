# dump example

## About this example

This example shows how object dumps works.

## Prepare

```sh
git clone https://github.com/aileron-projects/go-debugger.git
cd go-debugger/examples/dump/
```

## Run without tag `dump`

Run the example without tag.

```sh
go run ./main.go 
```

It show the profile of `bob`.
It does not show the profile of `alice`.

```txt
2026-08-01 11:13:16 [DEBUGGER][DUMP] dump always profile
  | Caller: Pkg:main File:main.go Func:main Line:31
  | ┌── args[0]
  | (*main.profile)({
  |  name: (string) (len=3) "bob",
  |  age: (int) 20,
  |  favorites: ([]string) (len=2) {
  |   (string) (len=5) "apple",
  |   (string) (len=10) "strawberry"
  |  },
  |  experience: (map[string]int) (len=4) {
  |   (string) (len=1) "C": (int) 6,
  |   (string) (len=2) "Go": (int) 3,
  |   (string) (len=4) "Java": (int) 1,
  |   (string) (len=4) "Rust": (int) 2
  |  }
  | })
```

## Run with tag `dump`

Run the example with tag.

```sh
go run -tags dump ./main.go
```

It shows both profiles of `alice` and `bob`.

```txt
2026-08-01 11:11:30 [DEBUGGER][DUMP] dump profile
  | Caller: Pkg:main File:main.go Func:main Line:28
  | ┌── args[0]
  | (*main.profile)({
  |  name: (string) (len=5) "alice",
  |  age: (int) 20,
  |  favorites: ([]string) (len=2) {
  |   (string) (len=5) "apple",
  |   (string) (len=6) "orange"
  |  },
  |  experience: (map[string]int) (len=3) {
  |   (string) (len=3) "C++": (int) 5,
  |   (string) (len=2) "Go": (int) 3,
  |   (string) (len=4) "Java": (int) 1
  |  }
  | })
2026-08-01 11:11:30 [DEBUGGER][DUMP] dump always profile
  | Caller: Pkg:main File:main.go Func:main Line:31
  | ┌── args[0]
  | (*main.profile)({
  |  name: (string) (len=3) "bob",
  |  age: (int) 20,
  |  favorites: ([]string) (len=2) {
  |   (string) (len=5) "apple",
  |   (string) (len=10) "strawberry"
  |  },
  |  experience: (map[string]int) (len=4) {
  |   (string) (len=1) "C": (int) 6,
  |   (string) (len=2) "Go": (int) 3,
  |   (string) (len=4) "Java": (int) 1,
  |   (string) (len=4) "Rust": (int) 2
  |  }
  | })
```
