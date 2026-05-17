package manager

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"

	device "github.com/iftsoft/linker/gen/go/linker/device/v1"
	"github.com/iftsoft/linker/model"
)

type ReaderManagerClient struct {
	log    *slog.Logger
	device device.ReaderManagerServiceClient
}

func NewReaderManagerClient(log *slog.Logger, conn *grpc.ClientConn) *ReaderManagerClient {
	return &ReaderManagerClient{
		log:    log,
		device: device.NewReaderManagerServiceClient(conn),
	}
}

// EnterCard trys to accept card in card reader device
func (c *ReaderManagerClient) EnterCard(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	c.log.Debug("ManagerClient.EnterCard - grpc", slog.String("device", query.Device))

	input := &device.EnterCardRequest{
		Query: convertDeviceQuery(query),
	}
	resp, err := c.device.EnterCard(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("command EnterCard for %s failed: %w", query.Device, err)
	}

	reply := convertDeviceReply(resp.GetReply())
	return reply, nil
}

// EjectCard trys to reject card from card reader device
func (c *ReaderManagerClient) EjectCard(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	c.log.Debug("ManagerClient.EjectCard - grpc", slog.String("device", query.Device))

	input := &device.EjectCardRequest{
		Query: convertDeviceQuery(query),
	}
	resp, err := c.device.EjectCard(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("command EjectCard for %s failed: %w", query.Device, err)
	}

	reply := convertDeviceReply(resp.GetReply())
	return reply, nil
}

// CaptureCard trys to capture card inside card reader device
func (c *ReaderManagerClient) CaptureCard(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	c.log.Debug("ManagerClient.CaptureCard - grpc", slog.String("device", query.Device))

	input := &device.CaptureCardRequest{
		Query: convertDeviceQuery(query),
	}
	resp, err := c.device.CaptureCard(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("command CaptureCard for %s failed: %w", query.Device, err)
	}

	reply := convertDeviceReply(resp.GetReply())
	return reply, nil
}

// ReadCard trys to read card information from card
func (c *ReaderManagerClient) ReadCard(ctx context.Context, query *model.DeviceQuery) (*model.ReadCardReply, error) {
	c.log.Debug("ManagerClient.ReadCard - grpc", slog.String("device", query.Device))

	input := &device.ReadCardRequest{
		Query: convertDeviceQuery(query),
	}
	resp, err := c.device.ReadCard(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("command ReadCard for %s failed: %w", query.Device, err)
	}

	reply := convertReadCardReply(resp.GetReply(), resp.GetCard())
	return reply, nil
}

func convertReadCardReply(reply *device.DeviceReply, card *device.CardDescription) *model.ReadCardReply {
	data := &model.ReadCardReply{
		Reply: convertDeviceReply(reply),
		Card:  convertCardDescription(card),
	}
	return data
}

func convertCardDescription(value *device.CardDescription) *model.CardDescription {
	if value == nil {
		return nil
	}
	data := &model.CardDescription{
		Device:  value.Device,
		CardPan: model.CardPAN(value.CardPan),
		ExpDate: value.ExpDate,
		Holder:  value.Holder,
		Track1:  model.Track(value.Track1),
		Track2:  model.Track(value.Track2),
		Track3:  model.Track(value.Track3),
	}
	return data
}
