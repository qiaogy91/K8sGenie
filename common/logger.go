package common

import (
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

var (
	logger *zerolog.Logger
	once   sync.Once
)

func Init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	ins := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	logger = &ins
}

func L() *zerolog.Logger {
	return logger
}

func init() {
	once.Do(Init)
}
