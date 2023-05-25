package reflectx

import (
	"fmt"
	"testing"
)

func TestFnName(t *testing.T) {
	type args struct {
		fn any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{
				fn: fmt.Println,
			},
			want: "Prin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FnName(tt.args.fn); got != tt.want {
				t.Errorf("FnName() = %v, want %v", got, tt.want)
			}
		})
	}
}
