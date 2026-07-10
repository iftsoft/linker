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
func (c *ReaderCallbackClient) CardPosition(ctx context.Context, value *model.CardPosition) error {
	c.log.Debug("CallbackClient.CardPosition - grpc",
		slog.String("device", value.Device), slog.Int("position", int(value.Position)))

	input := &device.CardPositionRequest{
		Notify: convertDeviceNotifyToProto(&value.DeviceNotify),
		Data:   convertPositionNotifyToProto(&value.PositionNotify),
	}
	_, err := c.device.CardPosition(ctx, input)
	if err != nil {
		return fmt.Errorf("callback CardPosition for %s.%d failed: %w", value.Device, value.Position, err)
	}

	return nil
}

// CardDescription sends notification about card information
func (c *ReaderCallbackClient) CardDescription(ctx context.Context, value *model.CardDescription) error {
	c.log.Debug("CallbackClient.CardDescription - grpc",
		slog.String("device", value.Device), slog.String("PAN", value.CardPan.String()))

	input := &device.CardDescriptionRequest{
		Notify: convertDeviceNotifyToProto(&value.DeviceNotify),
		Data:   convertCardContentToProto(&value.CardContent),
	}
	_, err := c.device.CardDescription(ctx, input)
	if err != nil {
		return fmt.Errorf("callback CardDescription for %s.%s failed: %w", value.Device, value.CardPan, err)
	}

	return nil
}

func convertPositionNotifyToProto(value *model.PositionNotify) *device.PositionNotify {
	if value == nil {
		return nil
	}
	data := &device.PositionNotify{
		Position: value.Position,
	}
	return data
}

func convertCardContentToProto(value *model.CardContent) *device.CardContent {
	if value == nil {
		return nil
	}
	data := &device.CardContent{
		CardPan: string(value.CardPan),
		ExpDate: value.ExpDate,
		Holder:  value.Holder,
		Track1:  string(value.Track1),
		Track2:  string(value.Track2),
		Track3:  string(value.Track3),
	}
	return data
}
