package models

import "time"

type LogLevel string

const (
	LogLevelInfo    LogLevel = "INFO"
	LogLevelWarning LogLevel = "WARNING"
	LogLevelError   LogLevel = "ERROR"
	LogLevelDebug   LogLevel = "DEBUG"
)

type LogEntry struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Level     LogLevel  `json:"level"`
	Service   string    `json:"service"`
	Message   string                 `json:"message"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Analysis  interface{}            `json:"analysis,omitempty"` // Loose typing to avoid cycle imports if defined in ai package, or better: struct in model
}
