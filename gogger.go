package gogger

import (
	"fmt"
	"io"
	"log"
)

var PresetFlag = log.LstdFlags | log.Lmsgprefix

var DiscardLogger = log.New(io.Discard, "", 0)

var Level = Info

var (
	NormalChannel   = Stdout
	CriticalChannel = Stderr
)

func newLogger(writer io.Writer, l LoggerLevel, tag string, flag int) *log.Logger {
	return log.New(
		writer,
		fmt.Sprintf("%c [%s] ", l.ToReadable()[0], tag),
		flag|PresetFlag,
	)
}

type Logger struct {
	error   *log.Logger
	warn    *log.Logger
	info    *log.Logger
	debug   *log.Logger
	verbose *log.Logger

	Normal   io.Writer
	Critical io.Writer

	Tag  string
	Flag int
}

func (l *Logger) Error() *log.Logger {
	if Level < Error {
		l.error = nil
		return DiscardLogger
	}

	if l.error != nil {
		return l.error
	}

	var writer io.Writer

	if l.Critical != nil {
		writer = l.Critical
	} else if CriticalChannel == Discard {
		return DiscardLogger
	} else {
		writer = CriticalChannel.ToWriter()
	}

	l.error = newLogger(writer, Error, l.Tag, l.Flag)

	return l.error
}

func (l *Logger) Warn() *log.Logger {
	if Level < Warn {
		l.warn = nil
		return DiscardLogger
	}

	if l.warn != nil {
		return l.warn
	}

	var writer io.Writer

	if l.Critical != nil {
		writer = l.Critical
	} else if CriticalChannel == Discard {
		return DiscardLogger
	} else {
		writer = CriticalChannel.ToWriter()
	}

	l.warn = newLogger(writer, Warn, l.Tag, l.Flag)

	return l.warn
}

func (l *Logger) Info() *log.Logger {
	if Level < Info {
		l.info = nil
		return DiscardLogger
	}

	if l.info != nil {
		return l.info
	}

	var writer io.Writer

	if l.Normal != nil {
		writer = l.Normal
	} else {
		writer = NormalChannel.ToWriter()
	}

	l.info = newLogger(writer, Info, l.Tag, l.Flag)

	return l.info
}

func (l *Logger) Debug() *log.Logger {
	if Level < Debug {
		l.debug = nil
		return DiscardLogger
	}

	if l.debug != nil {
		return l.debug
	}

	var writer io.Writer

	if l.Normal != nil {
		writer = l.Normal
	} else {
		writer = NormalChannel.ToWriter()
	}

	l.debug = newLogger(writer, Debug, l.Tag, l.Flag)

	return l.debug
}

func (l *Logger) Verbose() *log.Logger {
	if Level < Verbose {
		l.verbose = nil
		return DiscardLogger
	}

	if l.verbose != nil {
		return l.verbose
	}

	var writer io.Writer

	if l.Normal != nil {
		writer = l.Normal
	} else {
		writer = NormalChannel.ToWriter()
	}

	l.verbose = newLogger(writer, Verbose, l.Tag, l.Flag)

	return l.verbose
}

func NewWithWriter(tag string, flag int, normal, critical io.Writer) *Logger {
	return &Logger{
		Normal:   normal,
		Critical: critical,
		Tag:      tag,
		Flag:     flag,
	}
}

func NewWithFlag(tag string, flag int) *Logger {
	return NewWithWriter(tag, flag, nil, nil)
}

func New(tag string) *Logger {
	return NewWithFlag(tag, 0)
}
