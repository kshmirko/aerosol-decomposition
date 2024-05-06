package utlis

import (
	"reflect"
	"testing"
)

func TestLogSpace(t *testing.T) {
	type args struct {
		r0 float64
		r1 float64
		N  int
	}
	tests := []struct {
		name  string
		args  args
		want  []float64
		want1 float64
	}{
		{
			name: "Test1",
			args: args{r0: 0.1,
				r1: 1.0,
				N:  5,
			},
			want:  []float64{0.10000000000000002, 0.17782794100389232, 0.31622776601683794, 0.5623413251903491, 1.0},
			want1: 0.24999999999999997,
		},
		{
			name: "Test2",
			args: args{r0: 0.1,
				r1: 1.0,
				N:  3,
			},
			want:  []float64{0.10000000000000002, 0.31622776601683794, 1},
			want1: 0.49999999999999994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, got1 := LogSpace(tt.args.r0, tt.args.r1, tt.args.N)
			t.Log(tt.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogSpace() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LogSpace() got1 = %v, want %v", got1, tt.want1)
			}

		})
	}
}
