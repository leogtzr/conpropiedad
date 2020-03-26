package main

import "testing"

func Test_validateOptions(t *testing.T) {
	type args struct {
		options []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateOptions(tt.args.options...); got != tt.want {
				t.Errorf("validateOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
