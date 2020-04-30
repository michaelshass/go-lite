package main

import "testing"

func Test_handleMetaCommand(t *testing.T) {
	type args struct {
		command MetaCommand
	}
	tests := []struct {
		name       string
		args       args
		wantResult MetaCommandResultCode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := handleMetaCommand(tt.args.command); gotResult != tt.wantResult {
				t.Errorf("handleMetaCommand() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
