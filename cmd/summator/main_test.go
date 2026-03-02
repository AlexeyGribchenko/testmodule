package main

import (
	"testing"
)

func TestProcessArgs(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected int
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "valid numbers",
			args:     []string{"1", "2", "3", "4", "5"},
			expected: 15,
			wantErr:  false,
		},
		{
			name:     "negative numbers",
			args:     []string{"-1", "-2", "-3"},
			expected: -6,
			wantErr:  false,
		},
		{
			name:     "mixed numbers",
			args:     []string{"-1", "2", "-3", "4"},
			expected: 2,
			wantErr:  false,
		},
		{
			name:     "single number",
			args:     []string{"42"},
			expected: 42,
			wantErr:  false,
		},
		{
			name:     "zero",
			args:     []string{"0"},
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "empty args",
			args:     []string{},
			expected: 0,
			wantErr:  true,
			errMsg:   "no data to sum",
		},
		{
			name:     "invalid number",
			args:     []string{"1", "abc", "3"},
			expected: 0,
			wantErr:  true,
			errMsg:   "argument №2 is not a number: abc",
		},
		{
			name:     "float number",
			args:     []string{"1.5", "2"},
			expected: 0,
			wantErr:  true,
			errMsg:   "argument №1 is not a number: 1.5",
		},
		{
			name:     "empty string",
			args:     []string{"1", "", "3"},
			expected: 0,
			wantErr:  true,
			errMsg:   "argument №2 is not a number: ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := processArgs(tt.args)

			// Проверяем ошибку
			if tt.wantErr {
				if err == nil {
					t.Errorf("ProcessArgs() expected error, got nil")
				} else if err.Error() != tt.errMsg {
					t.Errorf("ProcessArgs() error = %v, want %v", err, tt.errMsg)
				}
				return
			}

			if err != nil {
				t.Errorf("ProcessArgs() unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("ProcessArgs() = %d, want %d", result, tt.expected)
			}
		})
	}
}
