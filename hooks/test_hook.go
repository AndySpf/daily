package hooks

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type TestHook struct {
}

func (m *TestHook) Levels() []log.Level {
	return log.AllLevels
}

func (*TestHook) Fire(entry *log.Entry) error {
	if err := os.Rename("./logs/log.1", "./logs/log.2"); err != nil {
		panic(err.Error())
	}

	file, err := os.OpenFile("./logs/log.1", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		panic(err.Error())
	}

	entry.Logger.Out = file
	return nil
}
