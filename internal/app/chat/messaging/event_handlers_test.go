package messaging_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"logur.dev/logur"
	"logur.dev/logur/logtesting"

	. "github.com/prasetyowira/message/internal/app/chat/messaging"
	"github.com/prasetyowira/message/internal/common/commonadapter"
)

func TestLogEventHandler_NewMessageSent(t *testing.T) {
	logger := &logur.TestLoggerFacade{}

	eventHandler := NewLogEventHandler(commonadapter.NewLogger(logger))

	event := MessageSent{
		ID: "1234",
		Text: "Hello World",
	}

	err := eventHandler.NewMessageSent(context.Background(), event)
	require.NoError(t, err)

	logEvent := logur.LogEvent{
		Level: logur.Info,
		Line:  "message sent",
		Fields: map[string]interface{}{
			"event":   "MessageSent",
			"message_id": "1234",
		},
	}

	logtesting.AssertLogEventsEqual(t, logEvent, *(logger.LastEvent()))
}
