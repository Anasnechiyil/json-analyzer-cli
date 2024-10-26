package cmd

import "testing"

// Test_countKeys tests the countKeys function.
//
// It takes a testing.T object as its parameter.
// No return values.
func Test_countKeys(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{
				filename: "test.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countKeys(tt.args.filename)
		})
	}
}
