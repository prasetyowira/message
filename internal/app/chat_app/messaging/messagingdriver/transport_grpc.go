package messagingdriver

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	appkitgrpc "github.com/sagikazarmark/appkit/transport/grpc"
	kitxgrpc "github.com/sagikazarmark/kitx/transport/grpc"

	messagingv1 "github.com/prasetyowira/message/.gen/api/proto/messaging/v1"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC server.
func MakeGRPCServer(endpoints Endpoints, options ...kitgrpc.ServerOption) messagingv1.MessageListKitServer {
	errorEncoder := kitxgrpc.NewStatusErrorResponseEncoder(appkitgrpc.NewDefaultStatusConverter())

	return messagingv1.MessageListKitServer{
		CreateMessageHandler: kitxgrpc.NewErrorEncoderHandler(kitgrpc.NewServer(
			endpoints.CreateMessage,
			decodeCreateMessageGRPCRequest,
			kitxgrpc.ErrorResponseEncoder(encodeCreateMessageGRPCResponse, errorEncoder),
			options...,
		), errorEncoder),
		ListMessagesHandler: kitxgrpc.NewErrorEncoderHandler(kitgrpc.NewServer(
			endpoints.ListMessages,
			decodeListMessagesGRPCRequest,
			kitxgrpc.ErrorResponseEncoder(encodeListMessagesGRPCResponse, errorEncoder),
			options...,
		), errorEncoder),
		GetMessageHandler: kitxgrpc.NewErrorEncoderHandler(kitgrpc.NewServer(
			endpoints.GetMessage,
			decodeGetMessageGRPCRequest,
			kitxgrpc.ErrorResponseEncoder(encodeGetMessageGRPCResponse, errorEncoder),
			options...,
		), errorEncoder),
	}
}

func decodeCreateMessageGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*messagingv1.CreateMessageRequest)

	return CreateMessageRequest{
		Text: req.GetText(),
	}, nil
}

func encodeCreateMessageGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(CreateMessageResponse)

	return &messagingv1.CreateMessageResponse{
		Id: resp.ID,
	}, nil
}

func decodeListMessagesGRPCRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return nil, nil
}

func encodeListMessagesGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(ListMessagesResponse)

	grpcResp := &messagingv1.ListMessagesResponse{
		Messages: make([]*messagingv1.Message, len(resp.Messages)),
	}

	for i, t := range resp.Messages {
		grpcResp.Messages[i] = &messagingv1.Message{
			Id:   t.ID,
			Text: t.Text,
		}
	}

	return grpcResp, nil
}

func decodeGetMessageGRPCRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*messagingv1.GetMessageRequest)

	return GetMessageRequest{
		Id: req.GetId(),
	}, nil
}

func encodeGetMessageGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(GetMessageResponse)

	return &messagingv1.GetMessageResponse{
		Message: &messagingv1.Message{
			Id:   resp.Message.ID,
			Text: resp.Message.Text,
		},
	}, nil
}
