package main

import (
	"context"
	"testing"

	"github.com/zoncoen/grpc-error-annotation/example/gen/example"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCheckErrorUnaryServerInterceptor(t *testing.T) {
	t.Run("defined status code", func(t *testing.T) {
		interceptor := example.NewUserServiceCheckErrorUnaryServerInterceptor(func(_ context.Context, _ *grpc.UnaryServerInfo, _ *status.Status) {
			t.Error("function called")
		})
		interceptor(nil, nil, &grpc.UnaryServerInfo{FullMethod: "/zoncoen.example.UserService/CreateUser"}, func(_ context.Context, _ interface{}) (interface{}, error) {
			return nil, status.New(codes.InvalidArgument, "invalid argument").Err()
		})
	})
	t.Run("not defined status code", func(t *testing.T) {
		var called bool
		interceptor := example.NewUserServiceCheckErrorUnaryServerInterceptor(func(_ context.Context, _ *grpc.UnaryServerInfo, _ *status.Status) {
			called = true
		})
		interceptor(nil, nil, &grpc.UnaryServerInfo{FullMethod: "/zoncoen.example.UserService/CreateUser"}, func(_ context.Context, _ interface{}) (interface{}, error) {
			return nil, status.New(codes.Unknown, "unknown").Err()
		})
		if !called {
			t.Error("function not called")
		}
	})
}
