package slog

import (
	"testing"
)

func TestAddCallerInfo(t *testing.T) {
	manager := New()

	tests := []struct {
		name string
		msg  string
	}{
		{"Test with simple message", "This is a test message"},
		{"Test with empty message", ""},
		{"Test with special characters", "Special characters !@#$%^&*()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := manager.addCallerInfo(tt.msg)
			if result == tt.msg {
				t.Errorf("Expected caller info to be added, but got: %s", result)
			}
		})
	}
}
