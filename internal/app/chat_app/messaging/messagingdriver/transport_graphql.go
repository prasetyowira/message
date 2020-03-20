package messagingdriver

import (
	"context"
	"errors"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-kit/kit/endpoint"

	"github.com/prasetyowira/message/.gen/api/graphql"
	"github.com/prasetyowira/message/internal/app/chat_app/messaging"
)

// MakeGraphQLHandler mounts all of the service endpoints into a GraphQL handler.
func MakeGraphQLHandler(endpoints Endpoints, errorHandler messaging.ErrorHandler) http.Handler {
	return handler.New(
		graphql.NewExecutableSchema(graphql.Config{
			Resolvers: &resolver{
				endpoints:    endpoints,
				errorHandler: errorHandler,
			},
		}),
	)
}

type resolver struct {
	endpoints    Endpoints
	errorHandler messaging.ErrorHandler
}

func (r *resolver) Mutation() graphql.MutationResolver {
	return &mutationResolver{r}
}
func (r *resolver) Query() graphql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *resolver }

func (r *mutationResolver) CreateMessage(ctx context.Context, input graphql.NewMessage) (string, error) {
	req := CreateMessageRequest{
		Text: input.Text,
	}

	resp, err := r.endpoints.CreateMessage(ctx, req)
	if err != nil {
		r.errorHandler.HandleContext(ctx, err)

		return "", errors.New("internal server error")
	}

	if f, ok := resp.(endpoint.Failer); ok {
		return "", f.Failed()
	}

	return resp.(CreateMessageResponse).Id, nil
}


type queryResolver struct{ *resolver }

func (r *queryResolver) Messages(ctx context.Context) ([]*graphql.Message, error) {
	resp, err := r.endpoints.ListMessages(ctx, nil)
	if err != nil {
		r.errorHandler.HandleContext(ctx, err)

		return nil, errors.New("internal server error")
	}

	messages := make([]*graphql.Message, len(resp.(ListMessagesResponse).Messages))

	for i, message := range resp.(ListMessagesResponse).Messages {
		messages[i] = &graphql.Message{
			ID: message.ID,
			Text: message.Text,
		}
	}

	return messages, nil
}

func (r *queryResolver) Message(ctx context.Context, input string) (*graphql.Message, error) {
	req := GetMessageRequest{
		Id: input,
	}
	resp, err := r.endpoints.GetMessage(ctx, req)
	if err != nil {
		r.errorHandler.HandleContext(ctx, err)

		return nil, errors.New("internal server error")
	}

	message := &graphql.Message{
		ID: resp.(GetMessageResponse).Message.ID,
		Text: resp.(GetMessageResponse).Message.Text,
	}

	return message, nil
}
