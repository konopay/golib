package sign

import (
	"reflect"
	"testing"

	"github.com/konopay/golib/errs"
)

func TestSignRequest(t *testing.T) {
	type args struct {
		payload   string
		secretKey string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 errs.Error
	}{
		{
			name: "test_sign",
			args: args{
				payload:   `{"hello":"world"}`,
				secretKey: "1234",
			},
			want:  "+vBg6DOUKfxsuv9VeHTvJs9d4lTjifYGImMWtvCsU7I=",
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SignRequest(tt.args.payload, tt.args.secretKey)
			if got != tt.want {
				t.Errorf("SignRequest() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SignRequest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
