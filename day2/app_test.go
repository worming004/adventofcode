package main

import (
	"reflect"
	"testing"
)

func Test_report_isAlwaysIncreasing(t *testing.T) {
	tests := []struct {
		name string
		r    report
		want bool
	}{
		{"test1", report{1, 2, 3, 4, 5}, true},
		{"test2", report{1, 2, 3, 5, 4}, false},
		{"test3", report{2, 1, 3, 4, 5}, false},
		{"test4", report{1, 2, 4, 3, 5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.isAlwaysIncreasing(); got != tt.want {
				t.Errorf("report.isAlwaysIncreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_report_isAlwaysDecreasing(t *testing.T) {
	tests := []struct {
		name string
		r    report
		want bool
	}{
		{"test1", report{5, 4, 3, 2, 1}, true},
		{"test2", report{4, 5, 3, 2, 1}, false},
		{"test3", report{5, 4, 3, 1, 2}, false},
		{"test4", report{5, 3, 4, 2, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.isAlwaysDecreasing(); got != tt.want {
				t.Errorf("report.isAlwaysDecreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_report_isDifferSafe(t *testing.T) {
	tests := []struct {
		name string
		r    report
		want bool
	}{
		{"test1", report{1, 2, 3, 4, 5}, true},
		{"test2", report{1, 3, 4, 6}, true},
		{"test3", report{1, 4, 5, 6}, true},
		{"test4", report{1, 5, 6}, false},
		{"test5", report{5, 4, 3, 2, 1}, true},
		{"test6", report{5, 4, 3, 1}, true},
		{"test7", report{6, 5, 4, 1}, true},
		{"test8", report{6, 5, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.isDifferSafe(); got != tt.want {
				t.Errorf("report.isDifferSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_report_isSafe(t *testing.T) {
	tests := []struct {
		name string
		r    report
		want bool
	}{
		{"test1", report{7, 6, 4, 2, 1}, true},
		{"test2", report{1, 2, 7, 8, 9}, false},
		{"test3", report{9, 7, 6, 2, 1}, false},
		{"test4", report{1, 3, 2, 4, 5}, false},
		{"test5", report{8, 6, 4, 4, 1}, false},
		{"test6", report{1, 3, 6, 7, 9}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.isSafe(); got != tt.want {
				t.Errorf("report.isSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_report_isDampenerSafe(t *testing.T) {
	tests := []struct {
		name string
		r    report
		want bool
	}{
		{"test1", report{7, 6, 4, 2, 1}, true},
		{"test2", report{1, 2, 7, 8, 9}, false},
		{"test3", report{9, 7, 6, 2, 1}, false},
		{"test4", report{1, 3, 2, 4, 5}, true},
		{"test5", report{8, 6, 4, 4, 1}, true},
		{"test6", report{1, 3, 6, 7, 9}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.isDampenerSafe(); got != tt.want {
				t.Errorf("report.isDampenerSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_report_generateAllSubReport(t *testing.T) {
	tests := []struct {
		name string
		r    report
		want []report
	}{
		{"test1", report{7, 6, 4, 2, 1}, []report{
			{6, 4, 2, 1},
			{7, 4, 2, 1},
			{7, 6, 2, 1},
			{7, 6, 4, 1},
			{7, 6, 4, 2},
		}},
		{"test2", report{1, 2, 3}, []report{
			{2, 3},
			{1, 3},
			{1, 2},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.generateAllSubReport(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("report.generateAllSubReport() = %v, want %v", got, tt.want)
			}
		})
	}
}
