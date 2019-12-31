package auth

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ClientInterceptor struct {
	token string
}

func (ci *ClientInterceptor) addToken(md metadata.MD) {
	md.Set("Authorization", ci.token)
}

func (ci *ClientInterceptor) setUserID(ctx context.Context, md metadata.MD) {
	md.Set(UserIDHeader, getUserIDContext(ctx))
}

func (ci *ClientInterceptor) setTargetID(ctx context.Context, md metadata.MD) {
	md.Set(UserIDHeader, getTargetIDContext(ctx))
}

func (ci *ClientInterceptor) setRoles(ctx context.Context, md metadata.MD) {
	b, _ := json.Marshal(getRolesContext(ctx))

	md.Set(RolesHeader, string(b))
}

func (ci *ClientInterceptor) wrapContext(ctx context.Context) context.Context {
	md, ok := metadata.FromOutgoingContext(ctx)

	if !ok {
		md = metadata.New(nil)
	}

	ci.addToken(md)
	ci.setUserID(ctx, md)
	ci.setTargetID(ctx, md)
	ci.setRoles(ctx, md)

	return metadata.NewOutgoingContext(ctx, md)
}

func UnaryClientInterceptor(token string) grpc.UnaryClientInterceptor {
	ci := &ClientInterceptor{
		token: token,
	}

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = ci.wrapContext(ctx)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func StreamClientInterceptor(token string) grpc.StreamClientInterceptor {
	ci := &ClientInterceptor{
		token: token,
	}

	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		ctx = ci.wrapContext(ctx)

		return streamer(ctx, desc, cc, method, opts...)
	}
}