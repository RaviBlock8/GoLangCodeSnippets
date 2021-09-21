package main

import "testing"

func Test_linkedlist_insert(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		ll   *linkedlist
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ll.insert(tt.args.value)
		})
	}
}
