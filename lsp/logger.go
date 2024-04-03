package lsp

import (
	"github.com/TBD54566975/ftl/internal/log"
)

type LogSink struct {
	server *Server
}

var _ log.Sink = (*LogSink)(nil)

func NewLogSink(server *Server) *LogSink {
	return &LogSink{server: server}
}

func (l *LogSink) Log(entry log.Entry) error {
	if entry.Level == log.Error {
		l.server.post(entry.Error)
	}
	return nil
}
