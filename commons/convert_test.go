package commons

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	mockS := "tester"
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "convert to pointer",
			args: args{
				s: mockS,
			},
			want: &mockS,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	mockI := 1
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		{
			name: "convert to pointer",
			args: args{
				i: mockI,
			},
			want: &mockI,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolean(t *testing.T) {
	mockB := true
	type args struct {
		i bool
	}
	tests := []struct {
		name string
		args args
		want *bool
	}{
		{
			name: "convert to pointer",
			args: args{
				i: mockB,
			},
			want: &mockB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Boolean(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Boolean() = %v, want %v", got, tt.want)
			}
		})
	}
}
