package internal

import (
	"reflect"
	"testing"
)

func TestCheckDuplicateZero(t *testing.T) {
	type args struct {
		dataN []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "check duplicate zero",
			args: args{
				dataN: []int32{1, 0, 2, 3, 0, 4, 5, 0},
			},
			want: []int32{
				1,
				0,
				0,
				2,
				3,
				0,
				0,
				4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckDuplicateZero(tt.args.dataN); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckDuplicateZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
