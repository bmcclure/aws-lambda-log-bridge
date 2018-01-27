package runner

import "github.com/bmcclure/aws-lambda-loggly-bridge/internal/app/request-handler"

type Runner interface {
	Run(handler request_handler.RequestHandler)
}
