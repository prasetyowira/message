package chat_app

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	entsql "github.com/facebookincubator/ent/dialect/sql"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/goph/idgen/ulidgen"
	"github.com/gorilla/mux"
	appkitendpoint "github.com/sagikazarmark/appkit/endpoint"
	appkithttp "github.com/sagikazarmark/appkit/transport/http"
	"github.com/sagikazarmark/kitx/correlation"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"
	kitxtransport "github.com/sagikazarmark/kitx/transport"
	kitxgrpc "github.com/sagikazarmark/kitx/transport/grpc"
	kitxhttp "github.com/sagikazarmark/kitx/transport/http"
	"google.golang.org/grpc"
	watermilllog "logur.dev/integration/watermill"

	messagingv1 "github.com/prasetyowira/message/.gen/api/proto/messaging/v1"
	"github.com/prasetyowira/message/internal/app/chat_app/httpbin"
	"github.com/prasetyowira/message/internal/app/chat_app/landing/landingdriver"
	"github.com/prasetyowira/message/internal/app/chat_app/landing/playground"
	"github.com/prasetyowira/message/internal/app/chat_app/messaging"
	"github.com/prasetyowira/message/internal/app/chat_app/messaging/messagingadapter"
	"github.com/prasetyowira/message/internal/app/chat_app/messaging/messagingadapter/ent"
	"github.com/prasetyowira/message/internal/app/chat_app/messaging/messagingadapter/ent/migrate"
	"github.com/prasetyowira/message/internal/app/chat_app/messaging/messagingdriver"
	"github.com/prasetyowira/message/internal/app/chat_app/messaging/messaginggen"
)

const messagingTopic = "messaging"

// InitializeApp initializes a new HTTP and a new gRPC application.
func InitializeApp(
	httpRouter *mux.Router,
	grpcServer *grpc.Server,
	publisher message.Publisher,
	subscriber message.Subscriber,
	storage string,
	db *sql.DB,
	logger Logger,
	errorHandler ErrorHandler,
) {
	endpointMiddleware := []endpoint.Middleware{
		correlation.Middleware(),
		opencensus.TraceEndpoint("", opencensus.WithSpanName(func(ctx context.Context, _ string) string {
			name, _ := kitxendpoint.OperationName(ctx)

			return name
		})),
		appkitendpoint.LoggingMiddleware(logger),
	}

	transportErrorHandler := kitxtransport.NewErrorHandler(errorHandler)

	httpServerOptions := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transportErrorHandler),
		kithttp.ServerErrorEncoder(kitxhttp.NewJSONProblemErrorEncoder(appkithttp.NewDefaultProblemConverter())),
		kithttp.ServerBefore(correlation.HTTPToContext()),
	}

	grpcServerOptions := []kitgrpc.ServerOption{
		kitgrpc.ServerErrorHandler(transportErrorHandler),
		kitgrpc.ServerBefore(correlation.GRPCToContext()),
	}

	{
		eventBus, _ := cqrs.NewEventBus(
			publisher,
			func(eventName string) string { return messagingTopic },
			cqrs.JSONMarshaler{GenerateName: cqrs.StructName},
		)

		var store messaging.Store = messaging.NewInMemoryStore()
		if storage == "database" {
			client := ent.NewClient(ent.Driver(entsql.OpenDB("mysql", db)))
			err := client.Schema.Create(
				context.Background(),
				migrate.WithDropIndex(true),
				migrate.WithDropColumn(true),
			)
			if err != nil {
				panic(err)
			}

			store = messagingadapter.NewEntStore(client)
		}

		service := messaging.NewService(
			ulidgen.NewGenerator(),
			store,
			messaginggen.NewEventDispatcher(eventBus),
		)
		service = messagingdriver.LoggingMiddleware(logger)(service)

		endpoints := messagingdriver.MakeEndpoints(
			service,
			kitxendpoint.Combine(endpointMiddleware...),
		)

		messagingdriver.RegisterHTTPHandlers(
			endpoints,
			httpRouter.PathPrefix("/message").Subrouter(),
			kitxhttp.ServerOptions(httpServerOptions),
		)

		messagingv1.RegisterMessageListServer(
			grpcServer,
			messagingdriver.MakeGRPCServer(
				endpoints,
				kitxgrpc.ServerOptions(grpcServerOptions),
			),
		)

		httpRouter.PathPrefix("/graphql").Handler(messagingdriver.MakeGraphQLHandler(endpoints, errorHandler))
	}

	landingdriver.RegisterHTTPHandlers(httpRouter)
	playground.RegisterHTTPHandlers(httpRouter)
	messagingdriver.RegisterWebSocketHandlers(httpRouter, subscriber)
	httpRouter.PathPrefix("/httpbin").Handler(http.StripPrefix(
		"/httpbin",
		httpbin.MakeHTTPHandler(logger.WithFields(map[string]interface{}{"module": "httpbin"})),
	))
}

// RegisterEventHandlers registers event handlers in a message router.
func RegisterEventHandlers(router *message.Router, subscriber message.Subscriber, logger Logger) error {
	messageEventProcessor, _ := cqrs.NewEventProcessor(
		[]cqrs.EventHandler{
			messaginggen.NewMessageSentEventHandler(messaging.NewLogEventHandler(logger), "message_sent"),
		},
		func(eventName string) string { return messagingTopic },
		func(handlerName string) (message.Subscriber, error) { return subscriber, nil },
		cqrs.JSONMarshaler{GenerateName: cqrs.StructName},
		watermilllog.New(logger.WithFields(map[string]interface{}{"component": "watermill"})),
	)

	err := messageEventProcessor.AddHandlersToRouter(router)
	if err != nil {
		return err
	}

	return nil
}
