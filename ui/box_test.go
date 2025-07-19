package ui

import (
	"strings"
	"testing"

	"github.com/bagaking/cmdux/style"
)

func TestBoxTitleAlignment(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		width    int
		expected string
	}{
		{
			name:     "Short title centered",
			title:    "Test",
			width:    20,
			expected: "╭─────[ Test ]─────╮",
		},
		{
			name:     "Medium title centered",
			title:    "Medium Title",
			width:    30,
			expected: "╭──────[ Medium Title ]──────╮",
		},
		{
			name:     "Long title truncated",
			title:    "Very Long Title That Should Be Truncated",
			width:    25,
			expected: "╭[ Very Long Title Th… ]╮",
		},
		{
			name:     "Emoji title",
			title:    "🚀 Title",
			width:    25,
			expected: "╭─────[ 🚀 Title ]──────╮",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			box := NewBox().
				Title(tt.title).
				Content("Test content").
				Width(tt.width)

			result := box.Render(style.DefaultTheme())
			lines := strings.Split(result, "\n")

			if len(lines) == 0 {
				t.Fatal("No output generated")
			}

			topLine := lines[0]
			// Remove color codes for comparison
			cleanLine := stripANSI(topLine)

			if cleanLine != tt.expected {
				t.Errorf("Expected: %q, Got: %q", tt.expected, cleanLine)
			}
		})
	}
}

func TestBoxWithoutTitle(t *testing.T) {
	box := NewBox().
		Content("Test content").
		Width(20)

	result := box.Render(style.DefaultTheme())
	lines := strings.Split(result, "\n")

	if len(lines) == 0 {
		t.Fatal("No output generated")
	}

	topLine := lines[0]
	cleanLine := stripANSI(topLine)
	expected := "╭──────────────────╮"

	if cleanLine != expected {
		t.Errorf("Expected: %q, Got: %q", expected, cleanLine)
	}
}

func TestBoxContentAlignment(t *testing.T) {
	box := NewBox().
		Title("Test").
		Content("Left aligned content").
		Width(30)

	result := box.Render(style.DefaultTheme())
	lines := strings.Split(result, "\n")

	if len(lines) < 3 {
		t.Fatal("Not enough lines generated")
	}

	// Check content line (should be left-aligned)
	contentLine := lines[1] // Skip title, no padding now
	cleanLine := stripANSI(contentLine)

	// Should start with border and padding, then content
	if !strings.HasPrefix(cleanLine, "│ Left aligned content") {
		t.Errorf("Content not left-aligned: %q", cleanLine)
	}
}

func TestBoxNoExtraEmptyLines(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected int // Expected number of content lines (excluding borders and padding)
	}{
		{
			name:     "Single line content",
			content:  "Single line",
			expected: 1,
		},
		{
			name:     "Two line content",
			content:  "Line 1\nLine 2",
			expected: 2,
		},
		{
			name:     "Empty content",
			content:  "",
			expected: 1, // Empty content should still render one line
		},
		{
			name:     "Multiple lines with empty lines",
			content:  "Line 1\n\nLine 3",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			box := NewBox().
				Title("Test").
				Content(tt.content).
				Width(30)

			result := box.Render(style.DefaultTheme())
			lines := strings.Split(result, "\n")

			// Count content lines (excluding borders and padding)
			contentLines := 0
			startedContent := false
			for _, line := range lines {
				cleanLine := stripANSI(line)
				if strings.HasPrefix(cleanLine, "│") && !strings.HasPrefix(cleanLine, "╭") && !strings.HasPrefix(cleanLine, "╰") {
					// This is a content line (has vertical border but not corner)
					if !startedContent {
						startedContent = true
					}
					contentLines++
				}
			}

			if contentLines != tt.expected {
				t.Errorf("Expected %d content lines, got %d", tt.expected, contentLines)
				t.Logf("Full output:\n%s", result)
			}
		})
	}
}

// stripANSI removes ANSI color codes from a string
func stripANSI(str string) string {
	var result strings.Builder
	inEscape := false

	for _, r := range str {
		if r == '\x1b' {
			inEscape = true
			continue
		}
		if inEscape {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				inEscape = false
			}
			continue
		}
		result.WriteRune(r)
	}

	return result.String()
}
