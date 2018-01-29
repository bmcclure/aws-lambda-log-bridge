package lambda_handler

import "context"

type RequestHandler interface {
	HandleRequest(ctx context.Context, event LambdaEvent) (LambdaResponse, error)
}
