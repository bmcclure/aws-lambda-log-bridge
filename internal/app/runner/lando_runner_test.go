package runner

import (
	"testing"
	"context"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/bmcclure/aws-lambda-loggly-bridge/internal/app/request-handler"
)

type testRequestHandler struct{}
func (testRequestHandler) HandleRequest(ctx context.Context, event request_handler.LambdaEvent) (request_handler.LambdaResponse, error) {
	return request_handler.LambdaResponse{Message: fmt.Sprintf("Hello %s!", event.Name)}, nil
}

func TestLambdaRunner_Run(t *testing.T) {
	Convey("Given a new LambdaRunner", t, func() {
		lambdaRunner := LambdaRunner{}

		Convey("Given a test RequestHandler", func() {
			requestHandler := testRequestHandler{}

			Convey("When I call the Run function", func() {
				lambdaRunner.Run(requestHandler)
			})
		})


	})

}
