package server

import (
	"context"
	"net"
	"testing"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/testing/testpb"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

// TestGRPCServer is an example gRPC ping test
func TestGRPCServer(t *testing.T) {
	lis := bufconn.Listen(bufSize)
	baseServer := grpc.NewServer()
	pingSvc := &testpb.TestPingService{}
	testpb.RegisterTestServiceServer(baseServer, pingSvc)

	go func() {
		if err := baseServer.Serve(lis); err != nil {
			t.Fatalf("Server failed to start: %v", err)
		}
	}()

	bufDialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}
	conn, err := grpc.NewClient(
		"passthrough:whatever", // passthrough is a transport that dials the server, port is ignored but required
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	defer conn.Close()

	client := testpb.NewTestServiceClient(conn)
	resp, err := client.Ping(context.Background(), &testpb.PingRequest{Value: "test"})
	require.NoError(t, err)
	require.Equal(t, "test", resp.Value)
}
