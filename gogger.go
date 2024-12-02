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

func newLogger(writer io.Writer, tag, appendix string, flag int) *log.Logger {
	prefix := fmt.Sprintf("%s", tag)
	if appendix == "" {
		prefix += " "
	} else {
		prefix += fmt.Sprintf(".%s ", appendix)
	}
	return log.New(writer, prefix, flag|PresetFlag)
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

	l.error = newLogger(writer, l.Tag, "ERROR", l.Flag)

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

	l.warn = newLogger(writer, l.Tag, "warn", l.Flag)

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

	l.info = newLogger(writer, l.Tag, "", l.Flag)

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

	l.debug = newLogger(writer, l.Tag, "", l.Flag)

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

	l.verbose = newLogger(writer, l.Tag, "", l.Flag)

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
