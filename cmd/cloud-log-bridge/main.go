package main

import "github.com/bmcclure/cloud-log-bridge/internal/app/lambda-handler"

func main() {
	handler := lambda_handler.LambdaRequestHandler{}
	runner := lambda_handler.LambdaRunner{}

	runner.Run(handler)
}
