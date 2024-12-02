package gogger

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type LoggerLevel uint

func (l LoggerLevel) ToReadable() ReadableLevel {
	if l >= Verbose {
		return RVerbose
	} else if l >= Debug {
		return RDebug
	} else if l >= Info {
		return RInfo
	} else if l >= Warn {
		return RWarn
	} else if l >= Error {
		return RError
	} else {
		return ROff
	}
}

type ReadableLevel string

func (l ReadableLevel) ToLevel() (LoggerLevel, error) {
	upper := strings.ToUpper(string(l))
	switch upper {
	case string(RError):
		return Error, nil
	case string(RWarn):
		return Warn, nil
	case string(RInfo):
		return Info, nil
	case string(RDebug):
		return Debug, nil
	case string(RVerbose):
		return Verbose, nil
	case string(ROff):
		return Off, nil
	default:
		return 0, fmt.Errorf("invalid level: %s", l)
	}
}

type LoggerChannel string

func (c LoggerChannel) ToWriter() io.Writer {
	switch c {
	case Stdout:
		return os.Stdout
	case Stderr:
		return os.Stderr
	case Discard:
		fallthrough
	default:
		return io.Discard
	}
}

const (
	Off     LoggerLevel = 0
	Error   LoggerLevel = 10
	Warn    LoggerLevel = 100
	Info    LoggerLevel = 1000
	Debug   LoggerLevel = 10000
	Verbose LoggerLevel = 100000
)

const (
	ROff     ReadableLevel = "OFF"
	RError   ReadableLevel = "ERROR"
	RWarn    ReadableLevel = "WARN"
	RInfo    ReadableLevel = "INFO"
	RDebug   ReadableLevel = "DEBUG"
	RVerbose ReadableLevel = "VERBOSE"
)

const (
	Stdout  LoggerChannel = "stdout"
	Stderr  LoggerChannel = "stderr"
	Discard LoggerChannel = "discard"
)
