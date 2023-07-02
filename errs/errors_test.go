package errs

import (
	"reflect"
	"testing"
)

func TestIsError(t *testing.T) {
	type args struct {
		a Error
		b Error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "not_equal",
			args: args{
				a: Success,
				b: New(2, "hello"),
			},
			want: false,
		},
		{
			name: "equal",
			args: args{
				a: Success,
				b: New(0, "hello"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IsError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKonoError_JoinErrMsg(t *testing.T) {
	type fields struct {
		ErrCode int32
		ErrMsg  string
	}
	type args struct {
		err Error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *KonoError
	}{
		{
			name: "join_two",
			fields: fields{
				ErrCode: 1,
				ErrMsg:  "outer_error",
			},
			args: args{
				err: DbError,
			},
			want: New(1, "outer_error,caused by 1:db error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &KonoError{
				ErrCode: tt.fields.ErrCode,
				ErrMsg:  tt.fields.ErrMsg,
			}
			if got := e.WithCause(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JoinErrMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
