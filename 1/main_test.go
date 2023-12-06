package main

import "testing"

func Test_extractRealDigits(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name           string
		args           args
		wantFirstDigit string
		wantLastDigit  string
	}{
		{
			name:           "Test Case 1",
			args:           args{line: "854rmtrkhjzrx2nine9ldqrqq"},
			wantFirstDigit: "8",
			wantLastDigit:  "9",
		},
		{
			name:           "Test Case 2",
			args:           args{line: "onejmrcmphzksixfbbftwo7"},
			wantFirstDigit: "1",
			wantLastDigit:  "7",
		}, {
			name:           "Test Case 3",
			args:           args{line: "8937"},
			wantFirstDigit: "8",
			wantLastDigit:  "7",
		}, {
			name:           "Test Case 4",
			args:           args{line: "8eight277kts7"},
			wantFirstDigit: "8",
			wantLastDigit:  "7",
		}, {
			name:           "Test Case 5",
			args:           args{line: "49four"},
			wantFirstDigit: "4",
			wantLastDigit:  "4",
		}, {
			name:           "Test Case 6",
			args:           args{line: "one6seven2n"},
			wantFirstDigit: "1",
			wantLastDigit:  "2",
		}, {
			name:           "Test Case 7",
			args:           args{line: "two6gmjd"},
			wantFirstDigit: "2",
			wantLastDigit:  "6",
		}, {
			name:           "Test Case 8",
			args:           args{line: "479"},
			wantFirstDigit: "4",
			wantLastDigit:  "9",
		}, {
			name:           "Test Case 9",
			args:           args{line: "fcfskbfive9eight"},
			wantFirstDigit: "5",
			wantLastDigit:  "8",
		}, {
			name:           "Test Case 10",
			args:           args{line: "jdmgmsglmpl3"},
			wantFirstDigit: "3",
			wantLastDigit:  "3",
		}, {
			name:           "Test Case 11",
			args:           args{line: "221frgs5nineeightwojj"},
			wantFirstDigit: "2",
			wantLastDigit:  "2",
		}, {
			name:           "Test Case 12",
			args:           args{line: "6ninevvrfqntzmjxpc"},
			wantFirstDigit: "6",
			wantLastDigit:  "9",
		}, {
			name:           "Test Case 13",
			args:           args{line: "four9threesixzlknkxz8one1pvxff"},
			wantFirstDigit: "4",
			wantLastDigit:  "1",
		}, {
			name:           "Test Case 14",
			args:           args{line: "seven33219"},
			wantFirstDigit: "7",
			wantLastDigit:  "9",
		}, {
			name:           "Test Case 15",
			args:           args{line: "98nine"},
			wantFirstDigit: "9",
			wantLastDigit:  "9",
		}, {
			name:           "Test Case 16",
			args:           args{line: "27three18"},
			wantFirstDigit: "2",
			wantLastDigit:  "8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirstDigit, gotLastDigit := extractRealDigits(tt.args.line)
			if gotFirstDigit != tt.wantFirstDigit {
				t.Errorf("extractRealDigits() gotFirstDigit = %v, want %v", gotFirstDigit, tt.wantFirstDigit)
			}
			if gotLastDigit != tt.wantLastDigit {
				t.Errorf("extractRealDigits() gotLastDigit = %v, want %v", gotLastDigit, tt.wantLastDigit)
			}
		})
	}
}
