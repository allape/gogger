package gogger

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestReadableLevelAndLevel(t *testing.T) {
	levels := []LoggerLevel{Verbose, Debug, Info, Warn, Error, 0}

	for _, level := range levels {
		readable := level.ToReadable()
		lvl, err := readable.ToLevel()
		if err != nil {
			t.Fatal(err)
		}
		if level != lvl {
			t.Fatalf("Expected: %v, Got: %v", level, lvl)
		}
	}
}

func TestEnv(t *testing.T) {
	err := os.Setenv(EnvLevel, string(RDebug))
	if err != nil {
		t.Fatal(err)
	}
	err = os.Setenv(EnvNormalChannel, string(Stderr))
	if err != nil {
		t.Fatal(err)
	}
	err = os.Setenv(EnvCriticalChannel, string(Stdout))
	if err != nil {
		t.Fatal(err)
	}

	err = InitFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	if Level != Debug {
		t.Fatalf("Expected: %v, Got: %v", Debug, Level)
	}
	if NormalChannel != Stderr {
		t.Fatalf("Expected: %v, Got: %v", Stderr, NormalChannel)
	}
	if CriticalChannel != Stdout {
		t.Fatalf("Expected: %v, Got: %v", Stdout, CriticalChannel)
	}
}

func TestLogger(t *testing.T) {
	err := os.Setenv(EnvLevel, string(RWarn))
	if err != nil {
		t.Fatal(err)
	}
	err = os.Setenv(EnvNormalChannel, string(Stderr))
	if err != nil {
		t.Fatal(err)
	}
	err = os.Setenv(EnvCriticalChannel, string(Discard))
	if err != nil {
		t.Fatal(err)
	}

	err = InitFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	l := New("test", log.LstdFlags, nil, nil)

	if l.Error() != DiscardLogger {
		t.Fatalf("Error logger should be DiscardLogger")
	}
	if l.Warn() != DiscardLogger {
		t.Fatalf("Warn logger should be DiscardLogger")
	}
	if l.Info() != DiscardLogger {
		t.Fatalf("Info logger should be DiscardLogger")
	}
	if l.Debug() != DiscardLogger {
		t.Fatalf("Debug logger should be DiscardLogger")
	}
	if l.Verbose() != DiscardLogger {
		t.Fatalf("Verbose logger should be DiscardLogger")
	}
}

func TestLevel(t *testing.T) {
	levels := []LoggerLevel{Verbose, Debug, Info, Warn, Error, 0}

	for _, level := range levels {
		Level = level

		nc := bytes.NewBuffer(nil)
		cc := bytes.NewBuffer(nil)

		l := New("test", log.LstdFlags, nc, cc)

		if level >= Verbose {
			l.Verbose().Print("verbose")
			logged := nc.String()
			if !strings.HasSuffix(logged, "test verbose\n") {
				t.Fatalf("Expected: verbose, Got: %s", logged)
			}
		}
		if level >= Debug {
			l.Debug().Print("debug")
			logged := nc.String()
			if !strings.HasSuffix(logged, "test debug\n") {
				t.Fatalf("Expected: debug, Got: %s", logged)
			}
		}
		if level >= Info {
			l.Info().Print("info")
			logged := nc.String()
			if !strings.HasSuffix(logged, "test info\n") {
				t.Fatalf("Expected: info, Got: %s", logged)
			}
		}
		if level >= Warn {
			l.Warn().Print("warn")
			logged := cc.String()
			if !strings.HasSuffix(logged, "test.warn warn\n") {
				t.Fatalf("Expected: warn, Got: %s", logged)
			}
		}
		if level >= Error {
			l.Error().Print("error")
			logged := cc.String()
			if !strings.HasSuffix(logged, "test.ERROR error\n") {
				t.Fatalf("Expected: error, Got: %s", logged)
			}
		}
	}
}
