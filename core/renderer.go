// Package core provides rendering utilities.
package core

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
)

// Renderer provides utilities for rendering text with proper alignment and sizing.
type Renderer struct {
	width  int
	height int
}

// NewRenderer creates a new renderer with the specified dimensions.
func NewRenderer(width, height int) *Renderer {
	return &Renderer{
		width:  width,
		height: height,
	}
}

// PadText pads text to the specified width with proper unicode handling.
func (r *Renderer) PadText(text string, width int, align Alignment) string {
	if width <= 0 {
		return text
	}
	
	textWidth := runewidth.StringWidth(text)
	if textWidth >= width {
		return runewidth.Truncate(text, width, "…")
	}
	
	padding := width - textWidth
	switch align {
	case AlignLeft:
		return text + strings.Repeat(" ", padding)
	case AlignRight:
		return strings.Repeat(" ", padding) + text
	case AlignCenter:
		leftPad := padding / 2
		rightPad := padding - leftPad
		return strings.Repeat(" ", leftPad) + text + strings.Repeat(" ", rightPad)
	default:
		return text + strings.Repeat(" ", padding)
	}
}

// WrapText wraps text to fit within the specified width.
func (r *Renderer) WrapText(text string, width int) []string {
	if width <= 0 {
		return []string{text}
	}
	
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{""}
	}
	
	var lines []string
	var currentLine string
	
	for _, word := range words {
		testLine := currentLine
		if testLine != "" {
			testLine += " "
		}
		testLine += word
		
		if runewidth.StringWidth(testLine) <= width {
			currentLine = testLine
		} else {
			if currentLine != "" {
				lines = append(lines, currentLine)
				currentLine = word
			} else {
				// Word is longer than width, truncate it
				lines = append(lines, runewidth.Truncate(word, width, "…"))
				currentLine = ""
			}
		}
	}
	
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	
	return lines
}

// TruncateText truncates text to fit within the specified width.
func (r *Renderer) TruncateText(text string, width int) string {
	if width <= 0 {
		return ""
	}
	return runewidth.Truncate(text, width, "…")
}

// CenterText centers text within the specified width.
func (r *Renderer) CenterText(text string, width int) string {
	return r.PadText(text, width, AlignCenter)
}

// RepeatChar repeats a character to create a string of specified width.
func (r *Renderer) RepeatChar(char rune, width int) string {
	if width <= 0 {
		return ""
	}
	charWidth := runewidth.RuneWidth(char)
	if charWidth == 0 {
		charWidth = 1
	}
	count := width / charWidth
	return strings.Repeat(string(char), count)
}

// JoinVertical joins multiple strings vertically with proper alignment.
func (r *Renderer) JoinVertical(strs []string, width int, align Alignment) string {
	var result []string
	for _, str := range strs {
		result = append(result, r.PadText(str, width, align))
	}
	return strings.Join(result, "\n")
}

// JoinHorizontal joins multiple strings horizontally.
func (r *Renderer) JoinHorizontal(strs []string, separator string) string {
	return strings.Join(strs, separator)
}

// Alignment represents text alignment options.
type Alignment int

const (
	// AlignLeft aligns text to the left.
	AlignLeft Alignment = iota
	// AlignCenter centers text.
	AlignCenter
	// AlignRight aligns text to the right.
	AlignRight
)

// Box draws a box around text with the specified characters.
func (r *Renderer) Box(content string, width, height int, chars BoxChars) string {
	if width < 3 || height < 3 {
		return content
	}
	
	contentWidth := width - 2  // Account for left and right borders
	contentHeight := height - 2 // Account for top and bottom borders
	
	lines := strings.Split(content, "\n")
	
	// Prepare content lines
	var contentLines []string
	for i := 0; i < contentHeight; i++ {
		if i < len(lines) {
			contentLines = append(contentLines, r.PadText(lines[i], contentWidth, AlignLeft))
		} else {
			contentLines = append(contentLines, strings.Repeat(" ", contentWidth))
		}
	}
	
	// Build the box
	var result []string
	
	// Top border
	topBorder := string(chars.TopLeft) + strings.Repeat(string(chars.Horizontal), width-2) + string(chars.TopRight)
	result = append(result, topBorder)
	
	// Content lines
	for _, line := range contentLines {
		contentLine := string(chars.Vertical) + line + string(chars.Vertical)
		result = append(result, contentLine)
	}
	
	// Bottom border
	bottomBorder := string(chars.BottomLeft) + strings.Repeat(string(chars.Horizontal), width-2) + string(chars.BottomRight)
	result = append(result, bottomBorder)
	
	return strings.Join(result, "\n")
}

// BoxChars defines the characters used for drawing boxes.
type BoxChars struct {
	TopLeft     rune
	TopRight    rune
	BottomLeft  rune
	BottomRight rune
	Horizontal  rune
	Vertical    rune
}

// DefaultBoxChars returns the default box drawing characters.
func DefaultBoxChars() BoxChars {
	return BoxChars{
		TopLeft:     '╭',
		TopRight:    '╮',
		BottomLeft:  '╰',
		BottomRight: '╯',
		Horizontal:  '─',
		Vertical:    '│',
	}
}

// ClassicBoxChars returns classic ASCII box drawing characters.
func ClassicBoxChars() BoxChars {
	return BoxChars{
		TopLeft:     '+',
		TopRight:    '+',
		BottomLeft:  '+',
		BottomRight: '+',
		Horizontal:  '-',
		Vertical:    '|',
	}
}

// GetTerminalSize attempts to get the terminal size. Returns default values if unable to detect.
func GetTerminalSize() (width, height int) {
	// This is a simplified implementation
	// In a real implementation, you would use terminal size detection
	return 80, 24
}

// StripANSI removes ANSI escape codes from a string for width calculation.
func StripANSI(str string) string {
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

// MeasureText measures the display width of text, handling ANSI codes and unicode.
func MeasureText(text string) int {
	return runewidth.StringWidth(StripANSI(text))
}

// FormatTable formats a table with proper column alignment and spacing.
func (r *Renderer) FormatTable(headers []string, rows [][]string, columnWidths []int) string {
	if len(headers) == 0 {
		return ""
	}
	
	var result []string
	
	// Format header
	headerRow := r.formatTableRow(headers, columnWidths)
	result = append(result, headerRow)
	
	// Add separator
	separator := r.formatTableSeparator(columnWidths)
	result = append(result, separator)
	
	// Format rows
	for _, row := range rows {
		formattedRow := r.formatTableRow(row, columnWidths)
		result = append(result, formattedRow)
	}
	
	return strings.Join(result, "\n")
}

func (r *Renderer) formatTableRow(cells []string, columnWidths []int) string {
	var parts []string
	parts = append(parts, "│")
	
	for i, cell := range cells {
		width := 10 // default width
		if i < len(columnWidths) {
			width = columnWidths[i]
		}
		
		paddedCell := r.PadText(cell, width, AlignLeft)
		parts = append(parts, fmt.Sprintf(" %s ", paddedCell))
		parts = append(parts, "│")
	}
	
	return strings.Join(parts, "")
}

func (r *Renderer) formatTableSeparator(columnWidths []int) string {
	var parts []string
	parts = append(parts, "├")
	
	for i, width := range columnWidths {
		if i > 0 {
			parts = append(parts, "┼")
		}
		parts = append(parts, strings.Repeat("─", width+2)) // +2 for padding
	}
	
	parts = append(parts, "┤")
	return strings.Join(parts, "")
}