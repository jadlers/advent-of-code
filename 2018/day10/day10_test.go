package main

import "testing"

func TestDay10(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name     string
		args     args
		wantDone bool
		wantP2   int
	}{
		{"example 1", args{[]string{
			"position=< 9,  1> velocity=< 0,  2>",
			"position=< 7,  0> velocity=<-1,  0>",
			"position=< 3, -2> velocity=<-1,  1>",
			"position=< 6, 10> velocity=<-2, -1>",
			"position=< 2, -4> velocity=< 2,  2>",
			"position=<-6, 10> velocity=< 2, -2>",
			"position=< 1,  8> velocity=< 1, -1>",
			"position=< 1,  7> velocity=< 1,  0>",
			"position=<-3, 11> velocity=< 1, -2>",
			"position=< 7,  6> velocity=<-1, -1>",
			"position=<-2,  3> velocity=< 1,  0>",
			"position=<-4,  3> velocity=< 2,  0>",
			"position=<10, -3> velocity=<-1,  1>",
			"position=< 5, 11> velocity=< 1, -2>",
			"position=< 4,  7> velocity=< 0, -1>",
			"position=< 8, -2> velocity=< 0,  1>",
			"position=<15,  0> velocity=<-2,  0>",
			"position=< 1,  6> velocity=< 1,  0>",
			"position=< 8,  9> velocity=< 0, -1>",
			"position=< 3,  3> velocity=<-1,  1>",
			"position=< 0,  5> velocity=< 0, -1>",
			"position=<-2,  2> velocity=< 2,  0>",
			"position=< 5, -2> velocity=< 1,  2>",
			"position=< 1,  4> velocity=< 2,  1>",
			"position=<-2,  7> velocity=< 2, -2>",
			"position=< 3,  6> velocity=<-1, -1>",
			"position=< 5,  0> velocity=< 1,  0>",
			"position=<-6,  0> velocity=< 2,  0>",
			"position=< 5,  9> velocity=< 1, -2>",
			"position=<14,  7> velocity=<-2,  0>",
			"position=<-3,  6> velocity=< 2, -1>",
		}}, true, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDone, gotP2 := Day10(tt.args.lines)
			if gotDone != tt.wantDone {
				t.Errorf("Day10() gotDone = %v, want %v", gotDone, tt.wantDone)
			}
			if gotP2 != tt.wantP2 {
				t.Errorf("Day10() gotP2 = %v, want %v", gotP2, tt.wantP2)
			}
		})
	}
}
