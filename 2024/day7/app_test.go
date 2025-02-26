package main

import (
	"reflect"
	"testing"
)

func Test_operationTester_test(t *testing.T) {
	type fields struct {
		result    int
		values    []int
		operators []operator
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"test1", fields{190, []int{10, 19}, []operator{add, mult}}, true},
		{"test2", fields{3267, []int{81, 40, 27}, []operator{add, mult}}, true},
		{"test3", fields{83, []int{14, 5}, []operator{add, mult}}, false},
		{"test4", fields{156, []int{15, 6}, []operator{add, mult}}, false},
		{"test5", fields{292, []int{11, 6, 16, 20}, []operator{add, mult}}, true},
		{"test6", fields{21037, []int{9, 7, 18, 13}, []operator{add, mult}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ot := &operationTester{
				result:    tt.fields.result,
				values:    tt.fields.values,
				operators: tt.fields.operators,
			}
			if got := ot.test(); got != tt.want {
				t.Errorf("operationTester.test() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []line
	}{
		{"test1", args{"10: 19 20 30 40\n"}, []line{{10, []int{19, 20, 30, 40}}}},
		{"test2", args{"21037: 9 7 18 13\n"}, []line{{21037, []int{9, 7, 18, 13}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_concat(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{1, 2}, 12},
		{"test2", args{12, 345}, 12345},
		{"test3", args{193, 35}, 19335},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concat(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("concat() = %v, want %v", got, tt.want)
			}
		})
	}
}
