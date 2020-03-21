package messaging

import (
	"context"

	"emperror.dev/errors"
)

// Message describing a message and its content
type Message struct {
	ID   string
	Text string
}

// +kit:endpoint:errorStrategy=service

// Service manages a list of messages.
type Service interface {
	// CreateMessage adds a new message to the message list.
	CreateMessage(ctx context.Context, text string) (id string, err error)

	// ListMessages returns the list of messages.
	ListMessages(ctx context.Context) (messages []Message, err error)

	// GetMessage returns the message by id.
	GetMessage(ctx context.Context, id string) (message Message, err error)
}

type service struct {
	idgenerator IDGenerator
	store       Store
	events      Events
}

// IDGenerator generates a new ID.
type IDGenerator interface {
	// Generate generates a new ID.
	Generate() (string, error)
}

// Store provides message persistence.
type Store interface {
	// Store stores a message.
	Store(ctx context.Context, message Message) error

	// All returns all messages.
	All(ctx context.Context) ([]Message, error)

	// Get returns a single message by its ID.
	Get(ctx context.Context, id string) (Message, error)
}

// NotFoundError is returned if a message cannot be found.
type NotFoundError struct {
	ID string
}

// Error implements the error interface.
func (NotFoundError) Error() string {
	return "message not found"
}

// Details returns error details.
func (e NotFoundError) Details() []interface{} {
	return []interface{}{"message_id", e.ID}
}

// NotFound tells a client that this error is related to a resource being not found.
// Can be used to translate the error to eg. status code.
func (NotFoundError) NotFound() bool {
	return true
}

// ServiceError tells the transport layer whether this error should be translated into the transport format
// or an internal error should be returned instead.
func (NotFoundError) ServiceError() bool {
	return true
}

// +mga:event:dispatcher

// Events dispatches message events.
type Events interface {
	// NewMessageSent dispatches a MessageSent event.
	NewMessageSent(ctx context.Context, event MessageSent) error
}

// +mga:event:handler

// MessageSent event is triggered when a new message created.
type MessageSent struct {
	ID 		string
	Text 	string
}

// NewService returns a new Service.
func NewService(idgenerator IDGenerator, store Store, events Events) Service {
	return &service{
		idgenerator: idgenerator,
		store:       store,
		events:      events,
	}
}

type validationError struct {
	violations map[string][]string
}

func (validationError) Error() string {
	return "invalid message"
}

func (e validationError) Violations() map[string][]string {
	return e.violations
}

// Validation tells a client that this error is related to a resource being invalid.
// Can be used to translate the error to eg. status code.
func (validationError) Validation() bool {
	return true
}

// ServiceError tells the transport layer whether this error should be translated into the transport format
// or an internal error should be returned instead.
func (validationError) ServiceError() bool {
	return true
}

func (s service) CreateMessage(ctx context.Context, text string) (string, error) {
	id, err := s.idgenerator.Generate()
	if err != nil {
		return "", err
	}

	if text == "" {
		return "", errors.WithStack(validationError{violations: map[string][]string{
			"text": {
				"text cannot be empty",
			},
		}})
	}

	message := Message{
		ID:   id,
		Text: text,
	}

	err = s.store.Store(ctx, message)
	if err != nil {
		return "", err
	}

	//nolint
	event := MessageSent{
		ID: message.ID,
		Text: message.Text,
	}

	err = s.events.NewMessageSent(ctx, event)

	return id, err
}

func (s service) ListMessages(ctx context.Context) ([]Message, error) {
	return s.store.All(ctx)
}

func (s service) GetMessage(ctx context.Context, id string) (Message, error) {
	message, err := s.store.Get(ctx, id)
	if err != nil {
		return Message{}, errors.WithMessage(err, "message not found")
	}

	return message, nil
}
