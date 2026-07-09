package manager

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"

	device "github.com/iftsoft/linker/gen/go/linker/device/v1"
	"github.com/iftsoft/linker/model"
)

type ValidatorManagerClient struct {
	log    *slog.Logger
	device device.ValidatorManagerServiceClient
}

func NewValidatorManagerClient(log *slog.Logger, conn *grpc.ClientConn) *ValidatorManagerClient {
	return &ValidatorManagerClient{
		log:    log,
		device: device.NewValidatorManagerServiceClient(conn),
	}
}

// InitValidator does primary initialization of the validator
func (c *ValidatorManagerClient) InitValidator(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	c.log.Debug("ManagerClient.InitValidator - grpc", slog.String("device", query.Device))

	input := &device.InitValidatorRequest{
		Query: convertValidatorQueryToProto(query),
	}
	resp, err := c.device.InitValidator(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("command InitValidator for %s failed: %w", query.Device, err)
	}

	reply := convertDeviceReplyToModel(resp.GetReply())
	return &reply, nil
}

// DoValidate starts accepting cash from user
func (c *ValidatorManagerClient) DoValidate(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	c.log.Debug("ManagerClient.DoValidate - grpc", slog.String("device", query.Device))

	input := &device.DoValidateRequest{
		Query: convertValidatorQueryToProto(query),
	}
	resp, err := c.device.DoValidate(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("command DoValidate for %s failed: %w", query.Device, err)
	}

	reply := convertDeviceReplyToModel(resp.GetReply())
	return &reply, nil
}

// AcceptNote puts the validated note to the cassette
func (c *ValidatorManagerClient) AcceptNote(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	c.log.Debug("ManagerClient.AcceptNote - grpc", slog.String("device", query.Device))

	input := &device.AcceptNoteRequest{
		Query: convertValidatorQueryToProto(query),
	}
	resp, err := c.device.AcceptNote(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("command AcceptNote for %s failed: %w", query.Device, err)
	}

	reply := convertDeviceReplyToModel(resp.GetReply())
	return &reply, nil
}

// ReturnNote returns the validated note to the user
func (c *ValidatorManagerClient) ReturnNote(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	c.log.Debug("ManagerClient.ReturnNote - grpc", slog.String("device", query.Device))

	input := &device.ReturnNoteRequest{
		Query: convertValidatorQueryToProto(query),
	}
	resp, err := c.device.ReturnNote(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("command ReturnNote for %s failed: %w", query.Device, err)
	}

	reply := convertDeviceReplyToModel(resp.GetReply())
	return &reply, nil
}

// StopValidate disables accepting new notes by validator
func (c *ValidatorManagerClient) StopValidate(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	c.log.Debug("ManagerClient.StopValidate - grpc", slog.String("device", query.Device))

	input := &device.StopValidateRequest{
		Query: convertValidatorQueryToProto(query),
	}
	resp, err := c.device.StopValidate(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("command StopValidate for %s failed: %w", query.Device, err)
	}

	reply := convertDeviceReplyToModel(resp.GetReply())
	return &reply, nil
}

// CheckValidator returns current cassette state
func (c *ValidatorManagerClient) CheckValidator(ctx context.Context, query *model.ValidatorQuery) (*model.ValidatorStore, error) {
	c.log.Debug("ManagerClient.CheckValidator - grpc", slog.String("device", query.Device))

	input := &device.CheckValidatorRequest{
		Query: convertValidatorQueryToProto(query),
	}
	resp, err := c.device.CheckValidator(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("command CheckValidator for %s failed: %w", query.Device, err)
	}

	reply := &model.ValidatorStore{
		DeviceReply:    convertDeviceReplyToModel(resp.GetReply()),
		ValidatorBatch: convertValidatorBatchToModel(resp.GetBatch()),
	}
	return reply, nil
}

// ClearValidator clears all cassette data (settlement or reconciliation)
func (c *ValidatorManagerClient) ClearValidator(ctx context.Context, query *model.ValidatorQuery) (*model.ValidatorStore, error) {
	c.log.Debug("ManagerClient.ClearValidator - grpc", slog.String("device", query.Device))

	input := &device.ClearValidatorRequest{
		Query: convertValidatorQueryToProto(query),
	}
	resp, err := c.device.ClearValidator(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("command ClearValidator for %s failed: %w", query.Device, err)
	}

	reply := &model.ValidatorStore{
		DeviceReply:    convertDeviceReplyToModel(resp.GetReply()),
		ValidatorBatch: convertValidatorBatchToModel(resp.GetBatch()),
	}
	return reply, nil
}

func convertValidatorQueryToProto(value *model.ValidatorQuery) *device.ValidatorQuery {
	if value == nil {
		return nil
	}
	data := &device.ValidatorQuery{
		Device:    value.Device,
		Currency:  uint32(value.Currency),
		Operation: value.Operation,
	}
	return data
}

func convertValidatorBatchToModel(value *device.ValidatorBatch) model.ValidatorBatch {
	if value == nil {
		return model.ValidatorBatch{}
	}
	data := model.ValidatorBatch{
		Device:  value.GetDevice(),
		BatchId: value.GetBatchId(),
		State:   model.BatchState(value.GetState()),
		Details: value.GetDetails(),
	}
	for _, note := range value.GetNotes() {
		cash := model.ValidatorNote{
			Currency: model.Currency(note.GetCurrency()),
			Nominal:  model.Amount(note.GetNominal()),
			Count:    model.Counter(note.GetCount()),
			Amount:   model.Amount(note.GetAmount()),
		}
		data.Notes = append(data.Notes, cash)
	}
	return data
}
