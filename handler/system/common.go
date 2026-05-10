package system

import (
	"encoding/json"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	StrMissingRequest  = "missing request"
	StrConversionError = "conversion error"
	StrServiceFault    = "service fault"
)

func MakeErrorWithDetails(code codes.Code, msg string, e error) error {
	details := &errdetails.ErrorInfo{
		Reason: e.Error(),
		Domain: "linker.v1",
	}

	sts, err := status.New(code, msg).WithDetails(details)
	if err != nil {
		sts = status.New(codes.Internal, err.Error())
	}

	return sts.Err()
}

func Serialize(value any) string {
	dump, err := json.Marshal(value)
	if err != nil {
		return "{}"
	}

	return string(dump)
}
