package interceptors

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/gilwong00/go-discord/pkg/token"
)

const tokenHeader = "Authentication"

func NewAuthInterceptor(maker token.Maker) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			token := req.Header().Get(tokenHeader)
			if token == "" {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no token provided"),
				)
			}
			_, err := maker.ValidateToken(token)
			if err != nil {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					fmt.Errorf("invalid token: %w", err),
				)
			}
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
