package util

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type EventResponse = *events.APIGatewayV2HTTPResponse
type EventRequest = *events.APIGatewayV2HTTPRequest

type Handler func(context.Context, EventRequest) (ResponseBuilder, ApiError)
type lambdaHandler func(context.Context, EventRequest) (EventResponse, error)

func getLambdaHandler(handler Handler) lambdaHandler {
	return func(ctx context.Context, event EventRequest) (EventResponse, error) {
		eventJSON, _ := json.Marshal(event)
		log.Printf("event:%s", string(eventJSON))

		var response EventResponse

		resBuilder, err := handler(ctx, event)
		if err != nil {
			log.Printf("request_id:%s, error:%s", event.RequestContext.RequestID, err.getMessage())
			response = err.build()
		} else {
			response = resBuilder.build()
		}

		log.Printf("request_id:%s, status:%d", event.RequestContext.RequestID, response.StatusCode)

		return response, nil
	}
}

func Start(handler Handler) {
	lambda.Start(getLambdaHandler(handler))
}
