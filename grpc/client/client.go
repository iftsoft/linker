package client

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"

	model "github.com/iftsoft/linker/model"
)

const (
	defaultConnectTimeout = 5 * time.Second
	defaultCallTimeout    = 5 * time.Second
)

// DeviceCallbackClient is the client API for DeviceCallbackService service.
type DeviceCallbackClient interface {
	// DeviceReply sends notification about device reply
	DeviceReply(ctx context.Context, reply *model.DeviceReply) error
	// ExecuteError sends notification about execute error
	ExecuteError(ctx context.Context, value *model.DeviceReply) error
	// StateChanged sends notification about device state changing
	StateChanged(ctx context.Context, value *model.DeviceState) error
	// ActionPrompt sends notification about action prompt for user
	ActionPrompt(ctx context.Context, value *model.DevicePrompt) error
	// ReaderReturn sends notification about device reading result
	ReaderReturn(ctx context.Context, value *model.DeviceInform) error
}

type Client struct {
	conn *grpc.ClientConn
}

func NewClient(ctx context.Context, address string) (*Client, error) {
	opts := []grpc.DialOption{}
	opts = append(opts,
		grpc.WithUserAgent("linker"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(TimeoutInterceptor(defaultCallTimeout)),
	)
	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		return nil, fmt.Errorf("new callback client: %w", err)
	}
	// Ensure connection is up within timeout before returning client
	cctx, cancel := context.WithTimeout(ctx, defaultConnectTimeout)
	defer cancel()

	conn.Connect()
	for s := conn.GetState(); s != connectivity.Ready; {
		if !conn.WaitForStateChange(cctx, s) {
			return nil, cctx.Err() // deadline or cancel
		}
		s = conn.GetState()
	}

	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Connection() *grpc.ClientConn {
	return c.conn
}

type TimeoutCallOption struct {
	grpc.EmptyCallOption
	forcedTimeout time.Duration
}

func WithForcedTimeout(forceTimeout time.Duration) TimeoutCallOption {
	return TimeoutCallOption{forcedTimeout: forceTimeout}
}

func getForcedTimeout(callOptions []grpc.CallOption) (time.Duration, bool) {
	for _, opt := range callOptions {
		if co, ok := opt.(TimeoutCallOption); ok {
			return co.forcedTimeout, true
		}
	}

	return 0, false
}

func TimeoutInterceptor(t time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		timeout := t
		if v, ok := getForcedTimeout(opts); ok {
			timeout = v
		}

		if timeout <= 0 {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
