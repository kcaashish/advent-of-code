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

func TestGetRibbonLength(t *testing.T) {
	var tests = []struct {
		name string
		in   string
		want int
	}{
		{
			name: "a",
			in:   "2x3x4\n1x1x10",
			want: 48},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := getRibbonLength(tt.in)
			if out != tt.want {
				t.Errorf("getRibbonLength(%s): got %d, want %d.", tt.in, out, tt.want)
			}
		})
	}

}

func TestGetLength(t *testing.T) {
	var tests = []struct {
		name string
		in   []string
		want int
	}{
		{name: "a", in: []string{"2", "4", "3"}, want: 34},
		{name: "b", in: []string{"1", "1", "10"}, want: 14},
		{name: "c", in: []string{"4", "3", "2"}, want: 34},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := getLength(tt.in)

			if out != tt.want {
				t.Errorf("getLength(%s): got %d, want %d.", tt.in, out, tt.want)
			}
		})
	}
}
