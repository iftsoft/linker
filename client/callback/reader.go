package callback

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"

	device "github.com/iftsoft/linker/gen/go/linker/device/v1"
	"github.com/iftsoft/linker/model"
)

type ReaderCallbackClient struct {
	log    *slog.Logger
	device device.ReaderCallbackServiceClient
}

func NewReaderCallbackClient(log *slog.Logger, conn *grpc.ClientConn) *ReaderCallbackClient {
	return &ReaderCallbackClient{
		log:    log,
		device: device.NewReaderCallbackServiceClient(conn),
	}
}

// CardPosition sends notification about new card position
func (c *ReaderCallbackClient) CardPosition(ctx context.Context, reply *model.CardPosition) error {
	c.log.Debug("CallbackClient.CardPosition - grpc",
		slog.String("device", reply.Device), slog.Int("position", int(reply.Position)))

	input := &device.CardPositionRequest{
		Data: convertCardPosition(reply),
	}
	_, err := c.device.CardPosition(ctx, input)
	if err != nil {
		return fmt.Errorf("callback CardPosition for %s.%d failed: %w", reply.Device, reply.Position, err)
	}

	return nil
}

// CardDescription sends notification about card information
func (c *ReaderCallbackClient) CardDescription(ctx context.Context, reply *model.CardDescription) error {
	c.log.Debug("CallbackClient.CardDescription - grpc",
		slog.String("device", reply.Device), slog.String("PAN", reply.CardPan.String()))

	input := &device.CardDescriptionRequest{
		Data: convertCardDescription(reply),
	}
	_, err := c.device.CardDescription(ctx, input)
	if err != nil {
		return fmt.Errorf("callback CardDescription for %s.%s failed: %w", reply.Device, reply.CardPan, err)
	}

	return nil
}

func convertCardPosition(value *model.CardPosition) *device.CardPosition {
	if value == nil {
		return nil
	}
	data := &device.CardPosition{
		Device:   value.Device,
		Position: value.Position,
	}
	return data
}

func convertCardDescription(value *model.CardDescription) *device.CardDescription {
	if value == nil {
		return nil
	}
	data := &device.CardDescription{
		Device:  value.Device,
		CardPan: string(value.CardPan),
		ExpDate: value.ExpDate,
		Holder:  value.Holder,
		Track1:  string(value.Track1),
		Track2:  string(value.Track2),
		Track3:  string(value.Track3),
	}
	return data
}
