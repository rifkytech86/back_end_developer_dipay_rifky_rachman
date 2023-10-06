package internal

import "testing"

func TestErrorMobile_String(t *testing.T) {
	tests := []struct {
		name string
		e    ErrorMobile
		want string
	}{
		{
			name: "tester",
			e:    ErrorMobile("1234"),
			want: "1234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorMobile_GetCode(t *testing.T) {
	tests := []struct {
		name string
		e    ErrorMobile
		want int64
	}{
		{
			name: "tester",
			e:    ErrorMobile("invalid request"),
			want: 400000,
		},
		{
			name: "not exist",
			e:    ErrorMobile("xxx"),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GetCode(); got != tt.want {
				t.Errorf("GetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCodeByString(t *testing.T) {
	type args struct {
		messageError string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "tester",
			args: args{
				messageError: "internal server error",
			},
			want: 500000,
		},
		{
			name: "not exist",
			args: args{
				messageError: "xxx",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCodeByString(tt.args.messageError); got != tt.want {
				t.Errorf("GetCodeByString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMsg(t *testing.T) {
	type args struct {
		code int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tester",
			args: args{
				code: 500000,
			},
			want: "internal server error",
		},
		{
			name: "not exist",
			args: args{
				code: 5003333,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMsg(tt.args.code); got != tt.want {
				t.Errorf("GetMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
