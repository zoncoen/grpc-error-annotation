// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/proto/user.proto

package example

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/zoncoen/grpc-error-annotation/gen/annotation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var mapUserService = map[string]struct{}{
	"/zoncoen.example.UserService/CreateUser/3": struct{}{},
	"/zoncoen.example.UserService/CreateUser/6": struct{}{},
	"/zoncoen.example.UserService/GetUser/3":    struct{}{},
}

// NewUserServiceCheckErrorUnaryServerInterceptor returns an interceptor to check status code.
// The interceptor calls function f if handler returns the status code which is not included in zoncoen.api.grpc.error annotation.
func NewUserServiceCheckErrorUnaryServerInterceptor(f func(context.Context, *grpc.UnaryServerInfo, *status.Status)) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err == nil {
			return resp, nil
		}
		s := status.Convert(err)
		if code := s.Code(); code != codes.OK {
			// ignore the status codes which may are generated by the library
			// ref. https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
			switch code {
			case codes.Canceled, codes.Unknown, codes.DeadlineExceeded, codes.PermissionDenied, codes.ResourceExhausted, codes.Unimplemented, codes.Internal, codes.Unavailable, codes.Unauthenticated:
				return resp, err
			}
			key := fmt.Sprintf("%s/%d", info.FullMethod, code)
			if _, ok := mapUserService[key]; !ok {
				f(ctx, info, s)
			}
		}
		return resp, err
	}
}

// NewUserServiceRetryUnaryClientInterceptor
func NewUserServiceRetryUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return nil
	}
}
