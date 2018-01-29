package lambda_handler

import (
	"fmt"
	"context"
)

type LambdaRequestHandler struct {}

func (LambdaRequestHandler) HandleRequest(ctx context.Context, event LambdaEvent) (LambdaResponse, error) {
	return LambdaResponse{Message: fmt.Sprintf("Hello %s!", event.Name)}, nil
}
