package auth

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

const (
	RegisterMethodPattern = "Register"
	LoginMethodPattern    = "Login"
)

type AuthInterceptor struct {
	jwtResolver JWTResolver
}

func NewAuthInterceptor(resolver JWTResolver) AuthInterceptor {
	return AuthInterceptor{
		jwtResolver: resolver,
	}
}

func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		if !strings.Contains(info.FullMethod, RegisterMethodPattern) && !strings.Contains(info.FullMethod, LoginMethodPattern) {
			// currently, should ignore access management
			_, err = i.authenticate(ctx)
			if err != nil {
				return nil, err
			}
		}
		return handler(ctx, req)
	}
}

func (i *AuthInterceptor) authenticate(ctx context.Context) (*JWTClaim, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "")
	}

	values := md["authorization"]
	if md["authorization"] == nil {
		values = md["Authorization"]
	}

	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "")
	}

	accessToken := values[0]
	claims, err := i.jwtResolver.VerifyJWTAccessToken(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "")
	}

	return claims, nil
}
