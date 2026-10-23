package reposerver

import (
	"context"

	"google.golang.org/grpc"
)

func cancelOnShutdownUnary(shutdownCtx context.Context) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		ctx, stop := contextWithShutdown(ctx, shutdownCtx)
		defer stop()
		return handler(ctx, req)
	}
}

func cancelOnShutdownStream(shutdownCtx context.Context) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx, stop := contextWithShutdown(ss.Context(), shutdownCtx)
		defer stop()
		return handler(srv, &shutdownServerStream{ServerStream: ss, ctx: ctx})
	}
}

func contextWithShutdown(ctx, shutdownCtx context.Context) (context.Context, func()) {
	ctx, cancel := context.WithCancelCause(ctx)
	stop := context.AfterFunc(shutdownCtx, func() {
		cancel(context.Cause(shutdownCtx))
	})
	return ctx, func() {
		stop()
		cancel(context.Canceled)
	}
}

type shutdownServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (s *shutdownServerStream) Context() context.Context {
	return s.ctx
}
