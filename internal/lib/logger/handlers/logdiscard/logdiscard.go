package logdiscard

import (
	"io"

	"github.com/charmbracelet/log"
)

// NewDiscardLogger создает логгер, который игнорирует все записи журнала.
func NewDiscardLogger() *log.Logger {
	return log.New(io.Discard)
}
