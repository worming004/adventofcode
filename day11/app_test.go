package main

import (
	"bytes"
	"log"
	"reflect"
	"testing"
)

func Test_split(t *testing.T) {
	type args struct {
		i          int
		digitCount int
	}
	tests := []struct {
		name string
		args args
		want struct{ a, b int }
	}{
		{"1234", args{1234, 4}, struct{ a, b int }{12, 34}},
		{"12", args{12, 2}, struct{ a, b int }{1, 2}},
		{"123456", args{123456, 6}, struct{ a, b int }{123, 456}},
		{"12345678", args{12345678, 8}, struct{ a, b int }{1234, 5678}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := split(tt.args.i, tt.args.digitCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlink(t *testing.T) {
	type args struct {
		s Stone
	}
	tests := []struct {
		name string
		args args
		want []Stone
	}{
		{"if 0", args{0}, []Stone{1}},
		{"if 20", args{20}, []Stone{2, 0}},
		{"if 2", args{2}, []Stone{4048}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Blink(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGlobal(t *testing.T) {
	tests := []struct {
		input    string
		loop     int
		expected int
	}{
		{"125 17", 6, 22},
		{"125 17", 25, 55312},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {

			var buff bytes.Buffer
			testLogger := log.New(&buff, "", log.LstdFlags)
			input := tt.input
			parsed := Parse(input)
			state := NewState(parsed, WithLogger(*testLogger))
			for i := 0; i < tt.loop; i++ {
				state.Blink()
			}

			if state.Length() != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, state.Length())
			}
			if t.Failed() {
				t.Log("\n" + buff.String())
			}
		})
	}
}
