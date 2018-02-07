// Package infrastructures consist of component used to connect to outside microservice (eg : database, log, etc.) //
package infrastructures

import (
	l4g "github.com/alecthomas/log4go"
)

// Logger - struct of logger
type Logger struct{}

// InitLogger - initialize logger
// InitLogger, initialize logger //
func (Logger) InitLogger(logPath string) {
	l4g.LoadConfiguration(logPath)
}
