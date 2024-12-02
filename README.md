# log.New wrapper

## Installation

```shell
go get -u github.com/allape/gogger
```

## Example

```go
package main

import (
	"github.com/allape/gogger"
	"io"
	"log"
	"os"
)

func main() {
	err := gogger.InitFromEnv()
	if err != nil {
		panic(err)
	}

	l := gogger.New("main")
	l.Error().Println("error")

	file, err := os.Create("main.log")
	if err != nil {
		panic(err)
	}

	stdout := io.MultiWriter(os.Stdout, file)
	stderr := io.MultiWriter(os.Stderr, file)

	fileLogger := gogger.NewWithWriter("main", log.LstdFlags, stdout, stderr)
	fileLogger.Error().Println("error")
}

```

### Environment variables

- `GOGGER_LEVEL`: log level, case insensitive, default `info`.
    - `off`
    - `error`
    - `warn`
    - `info`
    - `debug`
    - `verbose`
- `GOGGER_FLAG`: log flag, default `log.LstdFlags|log.Lmsgprefix` == `67`.
    - See https://golang.org/pkg/log/#pkg-constants for more information.
- `GOGGER_NORMAL_CHANNEL`: normal channel, for levels except `error` and `warn`, default `stdout`.
    - See `GOGGER_CRITICAL_CHANNEL`.
- `GOGGER_CRITICAL_CHANNEL`: critical channel, for `error` and `warn`, default `stderr`.
    - `stdout`: `os.Stdout`
    - `stderr`: `os.Stderr`
    - `discard`: `io.Discard`
