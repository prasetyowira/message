package messaging

import (
	"context"
	"testing"

	"github.com/goph/idgen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type messageEventsStub struct {
	messageSent MessageSent
}

func (s *messageEventsStub) NewMessageSent(ctx context.Context, event MessageSent) error {
	s.messageSent = event

	return nil
}

func TestList_CreatesAMessage(t *testing.T) {
	messageStore := NewInMemoryStore()

	const expectedID = "id"
	const text = "My first message lalala"
	events := &messageEventsStub{}

	messageList := NewService(idgen.NewConstantGenerator(expectedID), messageStore, events)

	id, err := messageList.CreateMessage(context.Background(), text)
	require.NoError(t, err)

	assert.Equal(t, expectedID, id)

	expectedMessage := Message{
		ID:   expectedID,
		Text: text,
	}

	message, err := messageStore.Get(context.Background(), id)
	require.NoError(t, err)

	assert.Equal(t, expectedMessage, message)

	expectedEvent := MessageSent{
		ID: "id",
		Text: text,
	}

	assert.Equal(t, expectedEvent, events.messageSent)
}

func TestList_CannotCreateAMessage(t *testing.T) {
	messageList := NewService(idgen.NewConstantGenerator("id"), NewReadOnlyStore(NewInMemoryStore()), nil)

	_, err := messageList.CreateMessage(context.Background(), "My first todo")
	require.Error(t, err)
}

func TestList_ListMessages(t *testing.T) {
	messageStore := NewInMemoryStore()

	message := Message{
		ID:   "id",
		Text: "Make the listing work",
	}
	require.NoError(t, messageStore.Store(context.Background(), message))

	todoList := NewService(idgen.NewConstantGenerator("id"), messageStore, nil)

	messages, err := todoList.ListMessages(context.Background())
	require.NoError(t, err)

	expectedMessages := []Message{message}

	assert.Equal(t, expectedMessages, messages)
}

func TestList_GetMessage(t *testing.T) {
	messageStore := NewInMemoryStore()

	const id = "id"

	message := Message{
		ID:   "id",
		Text: "getting message",
	}
	require.NoError(t, messageStore.Store(context.Background(), message))

	todoList := NewService(idgen.NewConstantGenerator("id"), messageStore, nil)

	message, err := todoList.GetMessage(context.Background(), id)
	require.NoError(t, err)

	expectedMessage := message

	assert.Equal(t, expectedMessage, message)
}
