package main

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_rtoi(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"test1", args{'0'}, 0},
		{"test1", args{'5'}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rtoi(tt.args.r); got != tt.want {
				t.Errorf("rtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDiskMap(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want DiskMap
	}{
		{"12345", args{"12345"}, DiskMap{Files: []File{{Position: 0, ID: 0, Length: 1}, {Position: 3, ID: 1, Length: 3}, {Position: 10, ID: 2, Length: 5}}}},
		{"2333133121414131402", args{"2333133121414131402"}, DiskMap{Files: []File{ // str representation : 00...111...2...333.44.5555.6666.777.888899
			{Position: 0, ID: 0, Length: 2},
			{Position: 5, ID: 1, Length: 3},
			{Position: 11, ID: 2, Length: 1},
			{Position: 15, ID: 3, Length: 3},
			{Position: 19, ID: 4, Length: 2},
			{Position: 22, ID: 5, Length: 4},
			{Position: 27, ID: 6, Length: 4},
			{Position: 32, ID: 7, Length: 3},
			{Position: 36, ID: 8, Length: 4},
			{Position: 40, ID: 9, Length: 2},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToDiskMap(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDiskMap() :\n%v, want :\n%v", got, tt.want)
			}
		})
	}
}

func TestCompact(t *testing.T) {
	type args struct {
		dms string
	}
	tests := []struct {
		name string
		args args
		want []Block
	}{
		{"12345", args{"12345"}, Blocks{
			//022111222......
			{Position: 0, ID: 0},
			{Position: 1, ID: 2},
			{Position: 2, ID: 2},
			{Position: 3, ID: 1},
			{Position: 4, ID: 1},
			{Position: 5, ID: 1},
			{Position: 6, ID: 2},
			{Position: 7, ID: 2},
			{Position: 8, ID: 2},
			{Position: 9, ID: 0, IsEmpty: true},
			{Position: 10, ID: 0, IsEmpty: true},
			{Position: 11, ID: 0, IsEmpty: true},
			{Position: 12, ID: 0, IsEmpty: true},
			{Position: 13, ID: 0, IsEmpty: true},
			{Position: 14, ID: 0, IsEmpty: true},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dm := ToDiskMap(tt.args.dms)
			if got := Compact(dm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBlock(t *testing.T) {
	input := ToDiskMap("12345")
	expected := Blocks{}.
		AddWithId(0).
		AddEmpty().
		AddEmpty().
		AddWithId(1).
		AddWithId(1).
		AddWithId(1).
		AddEmpty().
		AddEmpty().
		AddEmpty().
		AddEmpty().
		AddWithId(2).
		AddWithId(2).
		AddWithId(2).
		AddWithId(2).
		AddWithId(2)

	result := ToBlocks(input)
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Errorf("Slices do not match (-expected +actual):\n%s", diff)
	}
	// if !reflect.DeepEqual(result, expected) {
	// 	t.Errorf("ToBlock():\n%v.\nwant:\n%v", result, expected)
	// }
}

func (bs Blocks) AddWithId(id uint) Blocks {
	newBlock := Block{Position: uint(len(bs)), ID: id, IsEmpty: false}
	return append(bs, newBlock)
}
func (bs Blocks) AddEmpty() Blocks {
	newBlock := Block{Position: uint(len(bs)), ID: 0, IsEmpty: true}
	return append(bs, newBlock)
}

func TestChecksum(t *testing.T) {
	type args struct {
		dms string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2333133121414131402", args{"2333133121414131402"}, 1928},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dm := ToDiskMap(tt.args.dms)
			compacted := Compact(dm)
			res := Checksum(compacted)
			if res != tt.want {
				t.Errorf("Checksum() = %v, want %v", res, tt.want)
			}
		})
	}
}
