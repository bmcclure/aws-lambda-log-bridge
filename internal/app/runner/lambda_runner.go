package runner

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bmcclure/aws-lambda-loggly-bridge/internal/app/request-handler"
)

type LambdaRunner struct {}

func (LambdaRunner) Run(handler request_handler.RequestHandler) {
	lambda.Start(handler.HandleRequest)
}
