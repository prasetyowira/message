package messagingadapter


import (
	"context"

	"emperror.dev/errors"

	"github.com/prasetyowira/message/internal/app/chat_app/messaging"
	"github.com/prasetyowira/message/internal/app/chat_app/messaging/messagingadapter/ent"
	messagep "github.com/prasetyowira/message/internal/app/chat_app/messaging/messagingadapter/ent/message"
)

type entStore struct {
	client *ent.Client
}

// NewEntStore returns a new todo store backed by Ent ORM.
func NewEntStore(client *ent.Client) messaging.Store {
	return entStore{
		client: client,
	}
}

func (s entStore) Store(ctx context.Context, message messaging.Message) error {
	existing, err := s.client.Message.Query().Where(messagep.UID(message.ID)).First(ctx)
	if ent.IsNotFound(err) {
		_, err := s.client.Message.Create().
			SetUID(message.ID).
			SetText(message.Text).
			Save(ctx)
		if err != nil {
			return err
		}

		return nil
	}
	if err != nil {
		return err
	}

	_, err = s.client.Message.UpdateOneID(existing.ID).
		SetText(message.Text).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s entStore) All(ctx context.Context) ([]messaging.Message, error) {
	messageModels, err := s.client.Message.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	messages := make([]messaging.Message, 0, len(messageModels))

	for _, messageModel := range messageModels {
		messages = append(messages, messaging.Message{
			ID:   messageModel.UID,
			Text: messageModel.Text,
		})
	}

	return messages, nil
}

func (s entStore) Get(ctx context.Context, id string) (messaging.Message, error) {
	messageModel, err := s.client.Message.Query().Where(messagep.UID(id)).First(ctx)
	if ent.IsNotFound(err) {
		return messaging.Message{}, errors.WithStack(messaging.NotFoundError{ID: id})
	}

	return messaging.Message{
		ID:   messageModel.UID,
		Text: messageModel.Text,
	}, nil
}

