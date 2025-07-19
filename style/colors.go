// Package style provides theming and styling capabilities for cmdux.
package style

import (
	"github.com/fatih/color"
)

// Color wraps the fatih/color package for easy styling.
type Color = color.Color

// Style represents a collection of styling properties.
type Style struct {
	Foreground *Color
	Background *Color
	Bold       bool
	Italic     bool
	Underline  bool
	Faint      bool
}

// NewStyle creates a new style with default properties.
func NewStyle() *Style {
	return &Style{}
}

// Fg sets the foreground color.
func (s *Style) Fg(c *Color) *Style {
	s.Foreground = c
	return s
}

// Bg sets the background color.
func (s *Style) Bg(c *Color) *Style {
	s.Background = c
	return s
}

// SetBold enables or disables bold text.
func (s *Style) SetBold(bold bool) *Style {
	s.Bold = bold
	return s
}

// SetItalic enables or disables italic text.
func (s *Style) SetItalic(italic bool) *Style {
	s.Italic = italic
	return s
}

// SetUnderline enables or disables underlined text.
func (s *Style) SetUnderline(underline bool) *Style {
	s.Underline = underline
	return s
}

// SetFaint enables or disables faint text.
func (s *Style) SetFaint(faint bool) *Style {
	s.Faint = faint
	return s
}

// Render applies the style to the given text.
func (s *Style) Render(text string) string {
	if s == nil {
		return text
	}
	
	attrs := []color.Attribute{}
	
	// Add foreground color if set
	if s.Foreground != nil {
		// We can't directly access the attributes, so we'll create a new color
		// This is a simplified approach for now
	}
	
	if s.Bold {
		attrs = append(attrs, color.Bold)
	}
	if s.Italic {
		attrs = append(attrs, color.Italic)
	}
	if s.Underline {
		attrs = append(attrs, color.Underline)
	}
	if s.Faint {
		attrs = append(attrs, color.Faint)
	}
	
	// Create a new color with the attributes
	if len(attrs) > 0 {
		c := color.New(attrs...)
		return c.Sprint(text)
	}
	
	// If we have a foreground color but no other attributes, use it directly
	if s.Foreground != nil {
		return s.Foreground.Sprint(text)
	}
	
	// No styling, return text as-is
	return text
}

// Sprint applies the style and returns the styled string.
func (s *Style) Sprint(text string) string {
	return s.Render(text)
}

// Quick style constructors
var (
	// Primary colors
	Primary   = color.New(color.FgHiCyan, color.Bold)
	Secondary = color.New(color.FgHiBlue)
	Success   = color.New(color.FgHiGreen, color.Bold)
	Warning   = color.New(color.FgHiYellow)
	Error     = color.New(color.FgHiRed, color.Bold)
	Muted     = color.New(color.FgHiBlack)
	
	// Accent colors
	Accent1 = color.New(color.FgHiMagenta)
	Accent2 = color.New(color.FgHiCyan)
	Accent3 = color.New(color.FgHiWhite)
	
	// Text styles
	Bold      = color.New(color.Bold)
	Italic    = color.New(color.Italic)
	Underline = color.New(color.Underline)
	Faint     = color.New(color.Faint)
)