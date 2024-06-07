package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"1 + 2\n", "3"},
		{"VI / III\n", "II"},
		{"I - II\n", "panic: Invalid result: negative numbers are not allowed in the Roman numeral system."},
		{"I + 1\n", "panic: Invalid input: mixed numeric systems are not allowed."},
		{"1\n", "panic: Invalid input: the mathematical expression must contain two operands and one operator (+, -, /, *)."},
		{"1 + 2 + 3\n", "panic: Invalid input: the mathematical expression must contain two operands and one operator (+, -, /, *)."},
	}

	for _, c := range cases {
		func() {
			defer func() {
				if r := recover(); r != nil {
					// Convert any panic to a string and compare
					panicMessage := "panic: " + r.(string)
					if !strings.Contains(panicMessage, c.expected) {
						t.Errorf("Expected %v, but got %v", c.expected, panicMessage)
					}
				}
			}()

			// Redirect stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Set stdin to the input
			oldStdin := os.Stdin
			os.Stdin = stringToStdin(c.input)

			// Run main
			main()

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout
			var buf bytes.Buffer
			buf.ReadFrom(r)
			output := buf.String()

			// Restore stdin
			os.Stdin = oldStdin

			// Remove prompt from output
			if idx := strings.Index(output, "Enter an expression"); idx != -1 {
				output = output[idx+len("Enter an expression (e.g., 5 + 3 or VII - II): "):]
			}

			// Trim whitespace from output and expected values
			output = strings.TrimSpace(output)
			expected := strings.TrimSpace(c.expected)

			// Compare output
			if output != expected {
				t.Errorf("Expected %v, but got %v", expected, output)
			}
		}()
	}
}

func stringToStdin(s string) *os.File {
	r, w, _ := os.Pipe()
	w.WriteString(s)
	w.Close()
	return r
}
