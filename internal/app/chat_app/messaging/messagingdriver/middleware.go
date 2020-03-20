package messagingdriver

import (
	"context"

	"github.com/prasetyowira/message/internal/app/chat_app/messaging"
)

// Middleware describes a service middleware.
type Middleware func(messaging.Service) messaging.Service

// LoggingMiddleware is a service level logging middleware for TodoList.
func LoggingMiddleware(logger messaging.Logger) Middleware {
	return func(next messaging.Service) messaging.Service {
		return loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   messaging.Service
	logger messaging.Logger
}

func (mw loggingMiddleware) CreateMessage(ctx context.Context, text string) (string, error) {
	logger := mw.logger.WithContext(ctx)

	logger.Info("creating message")

	id, err := mw.next.CreateMessage(ctx, text)
	if err != nil {
		return id, err
	}

	logger.Info("created message", map[string]interface{}{
		"id": id,
	})

	return id, err
}

func (mw loggingMiddleware) ListMessages(ctx context.Context) ([]messaging.Message, error) {
	logger := mw.logger.WithContext(ctx)

	logger.Info("listing messages")

	return mw.next.ListMessages(ctx)
}

func (mw loggingMiddleware) GetMessage(ctx context.Context, id string) (messaging.Message, error) {
	logger := mw.logger.WithContext(ctx)

	logger.Info("get one message", map[string]interface{}{
		"id": id,
	})

	return mw.next.GetMessage(ctx, id)
}

