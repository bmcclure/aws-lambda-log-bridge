package request_handler

import (
	"testing"
	"context"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
)

var handleRequestTests = []struct {
	event LambdaEvent
	expected LambdaResponse
}{
	{LambdaEvent{Name: "Test"}, LambdaResponse{Message: "Hello Test!"}},
	{LambdaEvent{Name: "Ben"}, LambdaResponse{Message: "Hello Ben!"}},
	{LambdaEvent{Name: "John Smith"}, LambdaResponse{Message: "Hello John Smith!"}},
}

func TestLambdaRequestHandler_HandleRequest(t *testing.T) {
	Convey("Given a new LambdaRequestHandler", t, func() {
		handler := LambdaRequestHandler{}

		Convey("Given an empty context", func () {
			ctx := context.Background()

			for _, test := range handleRequestTests {
				Convey(fmt.Sprintf("Given a LambdaEvent with name \"%s\"", test.event.Name), func() {
					Convey("When I call the HandleRequest function", func() {
						response, err := handler.HandleRequest(ctx, test.event)

						Convey(fmt.Sprintf("The LambdaResponse message should be \"%s\"", test.expected.Message), func() {
							So(response.Message, ShouldEqual, test.expected.Message)
						})

						Convey("The err result should equal nil", func() {
							So(err, ShouldBeNil)
						})
					})
				})
			}
		})
	})
}
