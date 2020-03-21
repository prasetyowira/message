package messaging

import (
	"context"
	"sort"

	"emperror.dev/errors"
)

// InMemoryStore keeps messages in the memory.
// Use it in tests or for development/demo purposes.
type InMemoryStore struct {
	messages map[string]Message
}

// NewInMemoryStore returns a new inmemory message store.
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		messages: make(map[string]Message),
	}
}

// Store stores a message.
func (s *InMemoryStore) Store(ctx context.Context, message Message) error {
	s.messages[message.ID] = message

	return nil
}

// All returns all messages.
func (s *InMemoryStore) All(ctx context.Context) ([]Message, error) {
	messages := make([]Message, len(s.messages))

	// This makes sure todos are always returned in the same, sorted order
	keys := make([]string, 0, len(s.messages))
	for k := range s.messages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, key := range keys {
		messages[i] = s.messages[key]
	}

	return messages, nil
}

// Get returns a single message by its ID.
func (s *InMemoryStore) Get(ctx context.Context, id string) (Message, error) {
	message, ok := s.messages[id]
	if !ok {
		return message, NotFoundError{ID: id}
	}

	return message, nil
}

// ReadOnlyStore cannot be modified.
type ReadOnlyStore struct {
	store Store
}

// NewReadOnlyStore returns a new read-only message store instance.
func NewReadOnlyStore(store Store) *ReadOnlyStore {
	return &ReadOnlyStore{
		store: store,
	}
}

// Store stores a mesage.
func (*ReadOnlyStore) Store(ctx context.Context, message Message) error {
	return errors.NewWithDetails(
		"read-only message store cannot be modified",
		"message_id", message.ID)
}

// All returns all messages.
func (s *ReadOnlyStore) All(ctx context.Context) ([]Message, error) {
	return s.store.All(ctx)
}

// Get returns a single message by its ID.
func (s *ReadOnlyStore) Get(ctx context.Context, id string) (Message, error) {
	return s.store.Get(ctx, id)
}
