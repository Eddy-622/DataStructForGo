package main

import (
	"testing"
)

func TestStringUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"测试1",
			args{"abc"},
			"ABC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringUpper(tt.args.s); got != tt.want {
				t.Errorf("StringUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}
