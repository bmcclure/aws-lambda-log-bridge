package lambda_handler

import "github.com/aws/aws-lambda-go/lambda"

type LambdaRunner struct {}

func (LambdaRunner) Run(handler LambdaRequestHandler) {
	lambda.Start(handler.HandleRequest)
}
