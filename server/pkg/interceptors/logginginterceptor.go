package interceptors

import (
	"context"
	"log"
	"os"

	"connectrpc.com/connect"
)

func NewUnaryLoggingInteceptor() connect.UnaryInterceptorFunc {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	return connect.UnaryInterceptorFunc(
		func(next connect.UnaryFunc) connect.UnaryFunc {
			return connect.UnaryFunc(func(
				ctx context.Context,
				request connect.AnyRequest,
			) (connect.AnyResponse, error) {
				logger.Println("calling:", request.Spec().Procedure)
				logger.Println("request:", request.Any())
				return next(ctx, request)
			})
		},
	)
}
