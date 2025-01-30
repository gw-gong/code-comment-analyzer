package util

import (
	"fmt"
	"testing"
)

func TestFormatUserIDStr(t *testing.T) {
	tests := []struct {
		userID   uint64
		expected string
	}{
		{12345, "12345"},
		{0, "0"},
		{9876543210, "9876543210"},
		{18446744073709551615, "18446744073709551615"}, // 最大的 uint64 值
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("userID=%d", test.userID), func(t *testing.T) {
			result := FormatUserIDStr(test.userID)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}
