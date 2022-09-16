package notifier

import (
	"testing"
)

func Test_chunks(t *testing.T) {
	result := chunks("123", 5)
	if len(result) != 1 {
		t.Fatalf("Expected 1, got %d", len(result))
	}
	if result[0] != "123" {
		t.Fatalf("Expected 123, got %s", result[0])
	}

	result = chunks("1234567890", 5)
	if len(result) != 2 {
		t.Fatalf("Expected 2, got %d", len(result))
	}
	if result[0] != "12345" {
		t.Fatalf("Expected 12345, got %s", result[0])
	}
	if result[1] != "67890" {
		t.Fatalf("Expected 67890, got %s", result[1])
	}
}
