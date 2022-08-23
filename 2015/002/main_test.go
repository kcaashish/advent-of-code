package main

import "testing"

func TestGetWrappingPaperLength(t *testing.T) {
	var tests = []struct {
		name string
		in   string
		want int
	}{
		{
			name: "a",
			in:   "2x3x4\n1x1x10",
			want: 101},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := getWrappingPaperLength(tt.in)
			if out != tt.want {
				t.Errorf("getWrappingPaperLength(%s): got %d, want %d.", tt.in, out, tt.want)
			}
		})
	}

}

func TestGetArea(t *testing.T) {
	var tests = []struct {
		name string
		in   []string
		want int
	}{
		{name: "a", in: []string{"2", "3", "4"}, want: 58},
		{name: "b", in: []string{"1", "1", "10"}, want: 43},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := getArea(tt.in)

			if out != tt.want {
				t.Errorf("getArea(%s): got %d, want %d.", tt.in, out, tt.want)
			}
		})
	}
}
