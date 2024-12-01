package gogger

import (
	"github.com/allape/goenv"
)

const (
	EnvLevel           = "GOGGER_LEVEL"
	EnvNormalChannel   = "GOGGER_NORMAL_CHANNEL"
	EnvCriticalChannel = "GOGGER_CRITICAL_CHANNEL"
)

func InitFromEnv() error {
	readableLevel, err := goenv.GetSafeEnv(EnvLevel, RInfo)
	if err != nil {
		return err
	}

	level, err := readableLevel.ToLevel()
	if err != nil {
		return err
	}

	Level = level

	nc, err := goenv.GetSafeEnv(EnvNormalChannel, Stdout)
	if err != nil {
		return err
	}
	NormalChannel = nc

	cc, err := goenv.GetSafeEnv(EnvCriticalChannel, Stderr)
	if err != nil {
		return err
	}
	CriticalChannel = cc

	return nil
}
