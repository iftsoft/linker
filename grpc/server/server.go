package server

import (
	"context"
	"log/slog"
	"net"
	"runtime/debug"
	"time"

	pv "buf.build/go/protovalidate"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/testing/testpb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	gracefulStopTimeout = 10 * time.Second
)

type (
	Server struct {
		log  *slog.Logger
		conf Config
		conn *grpc.Server
	}

	ServiceRegistrar interface {
		Register(sr grpc.ServiceRegistrar)
	}
)

func New(log *slog.Logger, cfg Config, handlers ...ServiceRegistrar) *Server {
	srv := &Server{
		log:  log,
		conf: cfg,
	}

	// Init proto validator
	validator, err := pv.New()
	if err != nil {
		srv.log.Error("error create validator", slog.Any("error", err))

		return nil
	}

	serverOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandlerContext(srv.grpcPanicRecoveryHandler)),
			protovalidate.UnaryServerInterceptor(validator),
		),
		grpc.ChainStreamInterceptor(
			recovery.StreamServerInterceptor(recovery.WithRecoveryHandlerContext(srv.grpcPanicRecoveryHandler)),
			protovalidate.StreamServerInterceptor(validator),
		),
	}
	grpcSrv := grpc.NewServer(serverOptions...)

	// add ping-pong service
	t := &testpb.TestPingService{}
	testpb.RegisterTestServiceServer(grpcSrv, t)

	// This is a pretty cool feature, where it will give introspection capabilities to API (NOT FOR PROD!!!)
	if srv.conf.ReflectionEnable {
		reflection.Register(grpcSrv)
	}

	// register a server services handlers and its implementation
	for _, handler := range handlers {
		handler.Register(grpcSrv)
	}

	srv.conn = grpcSrv

	return srv
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.conf.Address)
	if err != nil {
		s.log.Error("GPRS server failed to listen", slog.Any("error", err))

		return errors.WithStack(err)
	}

	s.log.Info("Starting grpc server", slog.String("address", s.conf.Address))

	err = s.conn.Serve(lis)
	if err != nil {
		s.log.Error("GPRS server failed to start", slog.Any("error", err))

		return errors.WithStack(err)
	}

	return nil
}

func (s *Server) Shutdown() {
	s.log.Info("Stopping grpc server...")

	// Graceful stop according to best practices
	// https://github.com/grpc/grpc-go/tree/master/examples/features/gracefulstop#graceful-stop
	timer := time.AfterFunc(gracefulStopTimeout, func() {
		s.log.Warn("Server couldn't stop gracefully in time. Doing force stop.")
		s.conn.Stop()
	})
	defer timer.Stop()

	s.conn.GracefulStop()
}

// grpcPanicRecoveryHandler - panic recoveries.
func (s *Server) grpcPanicRecoveryHandler(ctx context.Context, p any) error {
	s.log.Error("recovered from panic", slog.Any("stack", debug.Stack()), slog.Any("data:", p))

	return status.Errorf(codes.Internal, "%s", p)
}

//func (s *Server) grpcExemplarFromContext(ctx context.Context) prometheus.Labels {
//	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
//		return prometheus.Labels{"traceID": span.TraceID().String()}
//	}
//
//	return nil
//}
