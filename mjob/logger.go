package mjob

import (
	"context"
	"encoding/json"
	"log"
)

type Logger interface {
	Log(ctx context.Context, e *LogEvent)
}

type LogEvent struct {
	Level   string      `json:"level"`
	Message string      `json:"msg"`
	JobID   string      `json:"job_id,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   error       `json:"error,omitempty"`
}

type DefaultLogger struct{}

func (l *DefaultLogger) Log(ctx context.Context, e *LogEvent) {
	var msg string
	if e != nil {
		msg = e.Message
		e.Message = ""
	}
	b, err := json.Marshal(e)
	if err != nil || e == nil {
		log.Print("error marshalling log event!", e, err)
	} else {
		log.Printf("%s  | %s", msg, string(b))
	}
}
