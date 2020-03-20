package messagingdriver

import (
	"context"
	"encoding/json"
	"net/http"

	"emperror.dev/errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	appkithttp "github.com/sagikazarmark/appkit/transport/http"
	kitxhttp "github.com/sagikazarmark/kitx/transport/http"

	api "github.com/prasetyowira/message/.gen/api/openapi/messaging/go"
)

// RegisterHTTPHandlers mounts all of the service endpoints into a router.
func RegisterHTTPHandlers(endpoints Endpoints, router *mux.Router, options ...kithttp.ServerOption) {
	errorEncoder := kitxhttp.NewJSONProblemErrorResponseEncoder(appkithttp.NewDefaultProblemConverter())

	router.Methods(http.MethodPost).Path("").Handler(kithttp.NewServer(
		endpoints.CreateMessage,
		decodeCreateMessageHTTPRequest,
		kitxhttp.ErrorResponseEncoder(encodeCreateMessageHTTPResponse, errorEncoder),
		options...,
	))

	router.Methods(http.MethodGet).Path("").Handler(kithttp.NewServer(
		endpoints.ListMessages,
		kithttp.NopRequestDecoder,
		kitxhttp.ErrorResponseEncoder(encodeListMessagesHTTPResponse, errorEncoder),
		options...,
	))

	router.Methods(http.MethodGet).Path("/{id}").Handler(kithttp.NewServer(
		endpoints.GetMessage,
		decodeGerMessageHTTPRequest,
		kitxhttp.ErrorResponseEncoder(encodeGetMessageHTTPResponse, errorEncoder),
		options...,
	))
}

func decodeCreateMessageHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var apiRequest api.CreateMessageRequest

	err := json.NewDecoder(r.Body).Decode(&apiRequest)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode request")
	}

	return CreateMessageRequest{
		Text: apiRequest.Text,
	}, nil
}

func encodeCreateMessageHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(CreateMessageResponse)

	apiResponse := api.CreateMessageResponse{
		Id: resp.Id,
	}

	return kitxhttp.JSONResponseEncoder(ctx, w, kitxhttp.WithStatusCode(apiResponse, http.StatusCreated))
}

func encodeListMessagesHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(ListMessagesResponse)

	apiResponse := api.MessageList{}

	for _, todo := range resp.Messages {
		apiResponse.Todos = append(apiResponse.Todos, api.Message{
			Id:   todo.ID,
			Text: todo.Text,
		})
	}

	return kitxhttp.JSONResponseEncoder(ctx, w, apiResponse)
}

func decodeGerMessageHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok || id == "" {
		return nil, errors.NewWithDetails("missing parameter from the URL", "param", "id")
	}

	return GetMessageRequest{
		Id: id,
	}, nil
}

func encodeGetMessageHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(GetMessageResponse)

	apiResponse := api.Message{
		Id: resp.Message.ID,
		Text: resp.Message.Text,
	}

	return kitxhttp.JSONResponseEncoder(ctx, w, apiResponse)
}
