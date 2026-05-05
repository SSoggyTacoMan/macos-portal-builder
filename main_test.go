package main

import "testing"

func TestShellQuote(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "''",
		},
		{
			name:     "Simple string",
			input:    "hello",
			expected: "'hello'",
		},
		{
			name:     "String with spaces",
			input:    "hello world",
			expected: "'hello world'",
		},
		{
			name:     "String with single quote",
			input:    "don't",
			expected: "'don'\\''t'",
		},
		{
			name:     "Multiple single quotes",
			input:    "a'b'c",
			expected: "'a'\\''b'\\''c'",
		},
		{
			name:     "String with already quoted content",
			input:    "'quoted'",
			expected: "''\\''quoted'\\'''",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := shellQuote(tt.input)
			if actual != tt.expected {
				t.Errorf("shellQuote(%q) = %q; want %q", tt.input, actual, tt.expected)
			}
		})
	}
}
