package messaging

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInMemoryStore_StoresAMessage(t *testing.T) {
	store := NewInMemoryStore()

	message := Message{
		ID:   "id",
		Text: "Store me!",
	}

	err := store.Store(context.Background(), message)
	require.NoError(t, err)

	assert.Equal(t, message, store.messages[message.ID])
}

func TestInMemoryStore_OverwritesAnExistingMessage(t *testing.T) {
	store := NewInMemoryStore()

	message := Message{
		ID:   "id",
		Text: "Store me first!",
	}

	err := store.Store(context.Background(), message)
	require.NoError(t, err)

	message = Message{
		ID:   "id",
		Text: "Store me!",
	}

	err = store.Store(context.Background(), message)
	require.NoError(t, err)

	assert.Equal(t, message, store.messages[message.ID])
}

func TestInMemoryStore_ListsAllMessages(t *testing.T) {
	store := NewInMemoryStore()

	store.messages["id"] = Message{
		ID:   "id",
		Text: "Store me first!",
	}

	store.messages["id2"] = Message{
		ID:   "id2",
		Text: "Store me second!",
	}

	messages, err := store.All(context.Background())
	require.NoError(t, err)

	expectedMessages := []Message{store.messages["id"], store.messages["id2"]}

	assert.Equal(t, expectedMessages, messages)
}

func TestInMemoryStore_GetsAMessage(t *testing.T) {
	store := NewInMemoryStore()

	id := "id"

	store.messages[id] = Message{
		ID:   id,
		Text: "Store me!",
	}

	message, err := store.Get(context.Background(), id)
	require.NoError(t, err)

	assert.Equal(t, store.messages[id], message)
}

func TestInMemoryStore_CannotReturnANonExistingMessage(t *testing.T) {
	store := NewInMemoryStore()

	_, err := store.Get(context.Background(), "id")
	require.Error(t, err)

	require.IsType(t, NotFoundError{}, err)

	e := err.(NotFoundError)
	assert.Equal(t, "id", e.ID)
}

func TestReadOnlyStore_IsReadOnly(t *testing.T) {
	message := Message{
		ID:   "id",
		Text: "Store me!",
	}

	store := NewReadOnlyStore(NewInMemoryStore())

	err := store.Store(context.Background(), message)
	require.Error(t, err)
}

func TestReadOnlyStore_ListsAllMessages(t *testing.T) {
	inmemStore := NewInMemoryStore()
	store := NewReadOnlyStore(inmemStore)

	inmemStore.messages["id"] = Message{
		ID:   "id",
		Text: "Store me first!",
	}

	inmemStore.messages["id2"] = Message{
		ID:   "id2",
		Text: "Store me second!",
	}

	messages, err := store.All(context.Background())
	require.NoError(t, err)

	expectedTodos := []Message{inmemStore.messages["id"], inmemStore.messages["id2"]}

	assert.Equal(t, expectedTodos, messages)
}

func TestReadOnlyStore_GetsAMessage(t *testing.T) {
	inmemStore := NewInMemoryStore()
	store := NewReadOnlyStore(inmemStore)

	id := "id"

	inmemStore.messages[id] = Message{
		ID:   id,
		Text: "Store me!",
	}

	message, err := store.Get(context.Background(), id)
	require.NoError(t, err)

	assert.Equal(t, inmemStore.messages[id], message)
}
