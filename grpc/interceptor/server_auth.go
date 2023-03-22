package interceptor

import (
	"context"
	"log"

	"github.com/yonisaka/user-service/grpc/contract"
	"github.com/yonisaka/user-service/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// UnaryAuthServerInterceptor is a function to handle authorization
func UnaryAuthServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// check if method is not protected
		if !contract.ProtectedMethods()[info.FullMethod] {
			return handler(ctx, req)
		}

		if errAuthorize := serverAuthorize(ctx); errAuthorize != nil {
			return nil, errAuthorize
		}

		return handler(ctx, req)
	}
}

// StreamAuthServerInterceptor is a function to handle stream authorization
func StreamAuthServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// check if method is not protected
		if !contract.ProtectedMethods()[info.FullMethod] {
			return handler(srv, stream)
		}

		if errAuthorize := serverAuthorize(stream.Context()); errAuthorize != nil {
			return errAuthorize
		}

		return handler(srv, stream)
	}
}

// serverAuthorize is a private function to handle authorization
func serverAuthorize(ctx context.Context) error {
	m, valid := metadata.FromIncomingContext(ctx)
	if !valid {
		return status.Error(codes.Unauthenticated, "no metadata provided")
	}

	tokenAuth, exist := m["authorization"]
	if !exist {
		return status.Error(codes.Unauthenticated, "no token provided")
	}

	decoded, err := utils.DecodeBasicAuth(tokenAuth[0])
	if err != nil {
		return status.Error(codes.Unauthenticated, "invalid token provided")
	}

	log.Printf("DECODED: %s\n", decoded)
	log.Printf("AUTH TOKEN: %s\n", tokenAuth)

	return nil
}
