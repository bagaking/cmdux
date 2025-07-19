// Package ui provides UI components for cmdux.
package ui

import (
	"strings"

	"github.com/bagaking/cmdux/core"
	"github.com/bagaking/cmdux/style"
	"github.com/mattn/go-runewidth"
)

// Box represents a rectangular container with optional border and title.
type Box struct {
	*core.Component
	title        string
	content      string
	padding      int
	border       bool
	borderStyle  *style.Color
	titleStyle   *style.Color
	contentStyle *style.Color
}

// NewBox creates a new box component.
func NewBox() *Box {
	return &Box{
		Component: core.NewComponent(),
		padding:   1,
		border:    true,
	}
}

// Title sets the box title.
func (b *Box) Title(title string) *Box {
	b.title = title
	return b
}

// Content sets the box content.
func (b *Box) Content(content string) *Box {
	b.content = content
	return b
}

// Padding sets the internal padding.
func (b *Box) Padding(padding int) *Box {
	b.padding = padding
	return b
}

// Width sets the box width and returns the box for chaining.
func (b *Box) Width(w int) *Box {
	b.Component.Width(w)
	return b
}

// Height sets the box height and returns the box for chaining.
func (b *Box) Height(h int) *Box {
	b.Component.Height(h)
	return b
}

// Border enables or disables the border.
func (b *Box) Border(enabled bool) *Box {
	b.border = enabled
	return b
}

// BorderStyle sets the border color.
func (b *Box) BorderStyle(color *style.Color) *Box {
	b.borderStyle = color
	return b
}

// TitleStyle sets the title color.
func (b *Box) TitleStyle(color *style.Color) *Box {
	b.titleStyle = color
	return b
}

// ContentStyle sets the content color.
func (b *Box) ContentStyle(color *style.Color) *Box {
	b.contentStyle = color
	return b
}

// Render renders the box using the given theme.
func (b *Box) Render(theme *style.Theme) string {
	if b.IsHidden() {
		return ""
	}

	width := b.GetWidth()
	if width <= 0 {
		width = b.calculateWidth()
	}

	height := b.GetHeight()
	if height <= 0 {
		height = b.calculateHeight(width)
	}

	if !b.border {
		return b.renderWithoutBorder(theme, width, height)
	}

	return b.renderWithBorder(theme, width, height)
}

func (b *Box) calculateWidth() int {
	// Calculate width based on content
	maxWidth := runewidth.StringWidth(b.title)

	lines := strings.Split(b.content, "\n")
	for _, line := range lines {
		lineWidth := runewidth.StringWidth(line)
		if lineWidth > maxWidth {
			maxWidth = lineWidth
		}
	}

	// Add padding and border
	return maxWidth + (b.padding * 2) + 2 // 2 for border
}

func (b *Box) calculateHeight(width int) int {
	contentWidth := width - (b.padding * 2) - 2 // Account for padding and border
	if contentWidth <= 0 {
		contentWidth = 1
	}

	// Count wrapped lines
	lines := strings.Split(b.content, "\n")
	totalLines := 0

	for _, line := range lines {
		if line == "" {
			totalLines++
			continue
		}

		lineWidth := runewidth.StringWidth(line)
		wrappedLines := (lineWidth + contentWidth - 1) / contentWidth // Ceiling division
		if wrappedLines == 0 {
			wrappedLines = 1
		}
		totalLines += wrappedLines
	}

	// Add padding, border, and title
	height := totalLines + (b.padding * 2) + 2 // 2 for top and bottom border
	if b.title != "" {
		height++ // Extra line for title
	}

	return height
}

func (b *Box) renderWithBorder(theme *style.Theme, width, height int) string {
	if width < 3 || height < 3 {
		return b.content
	}

	borderColor := b.borderStyle
	if borderColor == nil {
		borderColor = theme.Border
	}

	titleColor := b.titleStyle
	if titleColor == nil {
		titleColor = theme.Header
	}

	contentColor := b.contentStyle
	if contentColor == nil {
		contentColor = theme.Primary
	}

	var result []string

	// Top border with title
	if b.title != "" {
		titleStr := b.title
		titleWidth := runewidth.StringWidth(titleStr)

		// Calculate available space for title (accounting for borders and brackets)
		availableWidth := width - 2         // Account for left and right borders
		maxTitleWidth := availableWidth - 4 // Account for "[ ]" brackets

		if titleWidth > maxTitleWidth {
			titleStr = runewidth.Truncate(titleStr, maxTitleWidth, "…")
			titleWidth = maxTitleWidth
		}

		// Calculate padding to center the title
		totalPadding := availableWidth - titleWidth - 4 // 4 for "[ ]"
		leftPadding := totalPadding / 2
		rightPadding := totalPadding - leftPadding

		topLine := borderColor.Sprint(style.BoxTopLeft) +
			strings.Repeat(borderColor.Sprint(style.BoxHorizontal), leftPadding) +
			borderColor.Sprint("[ ") + titleColor.Sprint(titleStr) + borderColor.Sprint(" ]") +
			strings.Repeat(borderColor.Sprint(style.BoxHorizontal), rightPadding) +
			borderColor.Sprint(style.BoxTopRight)
		result = append(result, topLine)
	} else {
		topLine := borderColor.Sprint(style.BoxTopLeft) +
			strings.Repeat(borderColor.Sprint(style.BoxHorizontal), width-2) +
			borderColor.Sprint(style.BoxTopRight)
		result = append(result, topLine)
	}

	// Content area
	contentWidth := width - 2 - (b.padding * 2)
	if contentWidth <= 0 {
		contentWidth = 1
	}

	// Wrap and pad content
	contentLines := b.wrapContent(contentWidth)
	contentHeight := height - 2 // Remove top and bottom borders

	// Add padding rows at top
	for i := 0; i < b.padding; i++ {
		paddingLine := borderColor.Sprint(style.BoxVertical) +
			strings.Repeat(" ", width-2) +
			borderColor.Sprint(style.BoxVertical)
		result = append(result, paddingLine)
		contentHeight--
	}

	// Add content lines
	for i := 0; i < contentHeight-b.padding; i++ {
		var line string
		if i < len(contentLines) {
			line = contentColor.Sprint(contentLines[i])
		}

		// Pad line to fit width
		lineWidth := runewidth.StringWidth(core.StripANSI(line))
		padding := contentWidth - lineWidth
		if padding > 0 {
			line += strings.Repeat(" ", padding)
		}

		contentLine := borderColor.Sprint(style.BoxVertical) +
			strings.Repeat(" ", b.padding) +
			line +
			strings.Repeat(" ", b.padding) +
			borderColor.Sprint(style.BoxVertical)
		result = append(result, contentLine)
	}

	// Add padding rows at bottom
	for i := 0; i < b.padding; i++ {
		paddingLine := borderColor.Sprint(style.BoxVertical) +
			strings.Repeat(" ", width-2) +
			borderColor.Sprint(style.BoxVertical)
		result = append(result, paddingLine)
	}

	// Bottom border
	bottomLine := borderColor.Sprint(style.BoxBottomLeft) +
		strings.Repeat(borderColor.Sprint(style.BoxHorizontal), width-2) +
		borderColor.Sprint(style.BoxBottomRight)
	result = append(result, bottomLine)

	return strings.Join(result, "\n")
}

func (b *Box) renderWithoutBorder(theme *style.Theme, width, height int) string {
	contentColor := b.contentStyle
	if contentColor == nil {
		contentColor = theme.Primary
	}

	titleColor := b.titleStyle
	if titleColor == nil {
		titleColor = theme.Header
	}

	var result []string

	// Add title if present
	if b.title != "" {
		result = append(result, titleColor.Sprint(b.title))
	}

	// Add content
	contentWidth := width - (b.padding * 2)
	if contentWidth <= 0 {
		contentWidth = width
	}

	contentLines := b.wrapContent(contentWidth)
	for _, line := range contentLines {
		paddedLine := strings.Repeat(" ", b.padding) + contentColor.Sprint(line)
		result = append(result, paddedLine)
	}

	return strings.Join(result, "\n")
}

func (b *Box) wrapContent(width int) []string {
	if width <= 0 {
		return []string{b.content}
	}

	var result []string
	lines := strings.Split(b.content, "\n")

	for _, line := range lines {
		if line == "" {
			result = append(result, "")
			continue
		}

		// Simple word wrapping
		words := strings.Fields(line)
		if len(words) == 0 {
			result = append(result, "")
			continue
		}

		currentLine := ""
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
					result = append(result, currentLine)
					currentLine = word
				} else {
					// Word is longer than width, truncate it
					result = append(result, runewidth.Truncate(word, width, "…"))
					currentLine = ""
				}
			}
		}

		if currentLine != "" {
			result = append(result, currentLine)
		}
	}

	return result
}
