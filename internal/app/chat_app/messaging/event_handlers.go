package messaging

import (
	"context"
)

// LogEventHandler handles message events and logs them.
type LogEventHandler struct {
	logger Logger
}

// NewLogEventHandler returns a new LogEventHandler instance.
func NewLogEventHandler(logger Logger) LogEventHandler {
	return LogEventHandler{
		logger: logger,
	}
}

// NewMessageSent logs a MessageSent event.
func (h LogEventHandler) NewMessageSent(ctx context.Context, event MessageSent) error {
	logger := h.logger.WithContext(ctx)

	logger.Info("message sent", map[string]interface{}{
		"event":   "MessageSent",
		"message_id": event.ID,
	})

	return nil
}
