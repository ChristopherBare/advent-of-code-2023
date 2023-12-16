package main

import (
	"testing"
)

func Test_checkPossibleGames(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			filename string
			red      int
			green    int
			blue     int
		}
		want    int
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: struct {
				filename string
				red      int
				green    int
				blue     int
			}{
				filename: "test.txt",
				red:      12,
				green:    13,
				blue:     14,
			},
			want:    2286,
			wantErr: false,
		},
		// Add more test cases if needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkPossibleGames(tt.args.filename, tt.args.red, tt.args.green, tt.args.blue)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkPossibleGames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkPossibleGames() got = %v, want %v", got, tt.want)
			}
		})
	}
}
