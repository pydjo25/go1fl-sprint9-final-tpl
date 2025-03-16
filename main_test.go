package main

import (
	"testing"
)

func Test_generateRandomElements(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		// TODO: Add test cases.
		{
			name: "generate 5 elements",
			args: args{size: 5},
			want: 5,
		},
		{
			name: "generate 0 elements",
			args: args{size: 0},
			want: 0,
		},
		{
			name: "generate 10 000 000 elements",
			args: args{size: 10_000_000},
			want: 10_000_000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateRandomElements(tt.args.size); len(got) != tt.want {
				t.Errorf("generateRandomElements() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func Test_maximum(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "find maximum in empty array",
			args: args{data: []int{}},
			want: 0,
		},
		{
			name: "find maximum in array with one element",
			args: args{data: []int{1}},
			want: 1,
		},
		{
			name: "find maximum in array with two elements",
			args: args{data: []int{1, 2}},
			want: 2,
		},
		{
			name: "find maximum in array with many elements",
			args: args{data: []int{1, 2, 3, 4, 5}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.args.data); got != tt.want {
				t.Errorf("maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}
