package callback

import (
	"reflect"
	"testing"

	srv "github.com/iftsoft/linker/gen/go/linker/device/v1"
	model "github.com/iftsoft/linker/model"
)

func TestConvertDeviceReply(t *testing.T) {
	tests := []struct {
		name     string
		filters  *model.DeviceReply
		expected *srv.DeviceReply
	}{
		{
			name:     "empty input",
			filters:  &model.DeviceReply{},
			expected: &srv.DeviceReply{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertDeviceReply(tt.filters)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("convertDeviceReply() = %v, want %v", got, tt.expected)
			}
		})
	}
}
