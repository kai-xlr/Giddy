package buffer

import (
	"strings"
	"testing"
)

// TestNewRopeBasic tests creating a rope from a simple string
func TestNewRopeBasic(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"empty string", ""},
		{"single character", "a"},
		{"simple word", "hello"},
		{"sentence", "The quick brown fox jumps over the lazy dog"},
		{"with newlines", "line1\nline2\nline3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRope(tt.input)
			if r == nil {
				t.Fatal("NewRope returned nil")
			}
			
			got := r.String()
			if got != tt.input {
				t.Errorf("NewRope().String() = %q, want %q", got, tt.input)
			}
		})
	}
}

// TestNewLeaf tests the leaf node creation
func TestNewLeaf(t *testing.T) {
	text := []byte("hello world")
	leaf := NewLeaf(text)
	
	if leaf == nil {
		t.Fatal("NewLeaf returned nil")
	}
	
	if leaf.Weight != len(text) {
		t.Errorf("NewLeaf().Weight = %d, want %d", leaf.Weight, len(text))
	}
	
	if string(leaf.Value) != string(text) {
		t.Errorf("NewLeaf().Value = %q, want %q", string(leaf.Value), string(text))
	}
	
	// Test newline counting
	textWithNewlines := []byte("line1\nline2\nline3")
	leaf2 := NewLeaf(textWithNewlines)
	expectedNewlines := strings.Count(string(textWithNewlines), "\n")
	
	if leaf2.Newlines != expectedNewlines {
		t.Errorf("NewLeaf().Newlines = %d, want %d", leaf2.Newlines, expectedNewlines)
	}
}

// TestRopeBytes tests the Bytes() method
func TestRopeBytes(t *testing.T) {
	input := "hello world"
	r := NewRope(input)
	
	got := r.Bytes()
	if string(got) != input {
		t.Errorf("Bytes() = %q, want %q", string(got), input)
	}
}

// TestTotalWeight tests the TotalWeight() method on nodes
func TestTotalWeight(t *testing.T) {
	text := []byte("test")
	leaf := NewLeaf(text)
	
	if leaf.TotalWeight() != len(text) {
		t.Errorf("TotalWeight() = %d, want %d", leaf.TotalWeight(), len(text))
	}
}

// TestRopeInsert tests the Insert operation
func TestRopeInsert(t *testing.T) {
	tests := []struct {
		name     string
		initial  string
		pos      int
		insert   string
		expected string
	}{
		{"insert at start", "world", 0, "hello ", "hello world"},
		{"insert at end", "hello", 5, " world", "hello world"},
		{"insert in middle", "helo", 3, "l", "hello"},
		{"insert into empty", "", 0, "hello", "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRope(tt.initial)
			r2 := r.Insert(tt.pos, tt.insert)
			
			got := r2.String()
			if got != tt.expected {
				t.Errorf("Insert() = %q, want %q", got, tt.expected)
			}
			
			// Test immutability: original rope should be unchanged
			if r.String() != tt.initial {
				t.Errorf("Original rope was mutated: got %q, want %q", r.String(), tt.initial)
			}
		})
	}
}

// TestRopeDelete tests the Delete operation
func TestRopeDelete(t *testing.T) {
	tests := []struct {
		name     string
		initial  string
		start    int
		end      int
		expected string
	}{
		{"delete from start", "hello world", 0, 6, "world"},
		{"delete from end", "hello world", 5, 11, "hello"},
		{"delete middle", "hello world", 5, 6, "helloworld"},
		{"delete all", "hello", 0, 5, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRope(tt.initial)
			r2 := r.Delete(tt.start, tt.end)
			
			got := r2.String()
			if got != tt.expected {
				t.Errorf("Delete() = %q, want %q", got, tt.expected)
			}
			
			// Test immutability
			if r.String() != tt.initial {
				t.Errorf("Original rope was mutated: got %q, want %q", r.String(), tt.initial)
			}
		})
	}
}

// TestRopeLongString tests handling of strings longer than 1024 bytes
func TestRopeLongString(t *testing.T) {
	// Create a string longer than 1024 bytes
	longString := strings.Repeat("abcdefghij", 150) // 1500 bytes
	
	r := NewRope(longString)
	got := r.String()
	
	if got != longString {
		t.Errorf("Long string roundtrip failed, got length %d, want %d", len(got), len(longString))
	}
}

// TestRopeConcat tests the concat helper function
func TestRopeConcat(t *testing.T) {
	left := NewLeaf([]byte("hello "))
	right := NewLeaf([]byte("world"))
	
	result := concat(left, right)
	
	if result == nil {
		t.Fatal("concat returned nil")
	}
	
	if result.Left != left {
		t.Error("concat result.Left != left")
	}
	
	if result.Right != right {
		t.Error("concat result.Right != right")
	}
	
	expectedWeight := left.TotalWeight()
	if result.Weight != expectedWeight {
		t.Errorf("concat result.Weight = %d, want %d", result.Weight, expectedWeight)
	}
	
	expectedNewlines := left.Newlines + right.Newlines
	if result.Newlines != expectedNewlines {
		t.Errorf("concat result.Newlines = %d, want %d", result.Newlines, expectedNewlines)
	}
}

// TestRopeSplit tests the split helper function
func TestRopeSplit(t *testing.T) {
	text := []byte("hello world")
	leaf := NewLeaf(text)
	
	left, right := split(leaf, 6)
	
	if left == nil || right == nil {
		t.Fatal("split returned nil")
	}
	
	leftStr := string(left.Value)
	rightStr := string(right.Value)
	
	if leftStr != "hello " {
		t.Errorf("split left = %q, want %q", leftStr, "hello ")
	}
	
	if rightStr != "world" {
		t.Errorf("split right = %q, want %q", rightStr, "world")
	}
}

// TestRopeStressRandom performs random insertions and compares against a standard Go string
// This is the critical fuzz-like test mentioned in the spec
func TestRopeStressRandom(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping stress test in short mode")
	}
	
	// TODO: Implement 1,000 random inserts/deletes
	// Compare Rope output against standard Go string operations
	t.Skip("TODO: Implement random stress test with 1000 operations")
}
