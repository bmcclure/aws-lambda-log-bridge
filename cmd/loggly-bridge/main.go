package main

import (
	"github.com/bmcclure/aws-lambda-loggly-bridge/internal/app/runner"
	"github.com/bmcclure/aws-lambda-loggly-bridge/internal/app/request-handler"
)

func main() {
	lambdaRequestHandler := request_handler.LambdaRequestHandler{}
	lambdaRunner := runner.LambdaRunner{}

	lambdaRunner.Run(lambdaRequestHandler)
}
