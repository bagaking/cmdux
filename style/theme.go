// Package style provides theming support.
package style

import "github.com/fatih/color"

// Theme represents a cohesive color and styling theme.
type Theme struct {
	// Primary colors
	Primary   *Color
	Secondary *Color
	Success   *Color
	Warning   *Color
	Error     *Color
	Muted     *Color
	
	// Accent colors
	Accent1 *Color
	Accent2 *Color
	Accent3 *Color
	
	// Text styles
	Bold      *Color
	Italic    *Color
	Underline *Color
	Faint     *Color
	
	// UI Elements
	Border    *Color
	Header    *Color
	Footer    *Color
	Selected  *Color
	Disabled  *Color
}

// NewTheme creates a new theme with default colors.
func NewTheme() *Theme {
	return &Theme{
		Primary:   color.New(color.FgHiCyan, color.Bold),
		Secondary: color.New(color.FgHiBlue),
		Success:   color.New(color.FgHiGreen, color.Bold),
		Warning:   color.New(color.FgHiYellow),
		Error:     color.New(color.FgHiRed, color.Bold),
		Muted:     color.New(color.FgHiBlack),
		
		Accent1:   color.New(color.FgHiMagenta),
		Accent2:   color.New(color.FgHiCyan),
		Accent3:   color.New(color.FgHiWhite),
		
		Bold:      color.New(color.Bold),
		Italic:    color.New(color.Italic),
		Underline: color.New(color.Underline),
		Faint:     color.New(color.Faint),
		
		Border:    color.New(color.FgHiCyan),
		Header:    color.New(color.FgHiWhite, color.Bold),
		Footer:    color.New(color.FgHiBlack),
		Selected:  color.New(color.FgHiMagenta),
		Disabled:  color.New(color.FgHiBlack),
	}
}

// DefaultTheme returns the default cmdux theme.
func DefaultTheme() *Theme {
	return NewTheme()
}

// DarkTheme returns a dark theme optimized for dark terminals.
func DarkTheme() *Theme {
	theme := NewTheme()
	theme.Primary = color.New(color.FgHiGreen, color.Bold)
	theme.Secondary = color.New(color.FgGreen)
	theme.Accent1 = color.New(color.FgHiYellow)
	theme.Border = color.New(color.FgGreen)
	return theme
}

// LightTheme returns a light theme optimized for light terminals.
func LightTheme() *Theme {
	theme := NewTheme()
	theme.Primary = color.New(color.FgBlue, color.Bold)
	theme.Secondary = color.New(color.FgHiBlue)
	theme.Muted = color.New(color.FgBlack)
	theme.Border = color.New(color.FgBlue)
	return theme
}

// CyberpunkTheme returns a cyberpunk-style theme.
func CyberpunkTheme() *Theme {
	theme := NewTheme()
	theme.Primary = color.New(color.FgHiMagenta, color.Bold)
	theme.Secondary = color.New(color.FgMagenta)
	theme.Success = color.New(color.FgHiGreen, color.Bold)
	theme.Accent1 = color.New(color.FgHiCyan)
	theme.Accent2 = color.New(color.FgHiYellow)
	theme.Border = color.New(color.FgHiMagenta)
	theme.Selected = color.New(color.FgHiYellow, color.Bold)
	return theme
}

// MonochromeTheme returns a monochrome theme.
func MonochromeTheme() *Theme {
	theme := NewTheme()
	theme.Primary = color.New(color.FgHiWhite, color.Bold)
	theme.Secondary = color.New(color.FgWhite)
	theme.Success = color.New(color.FgHiWhite, color.Bold)
	theme.Warning = color.New(color.FgHiWhite, color.Bold)
	theme.Error = color.New(color.FgHiWhite, color.Bold)
	theme.Accent1 = color.New(color.FgWhite)
	theme.Accent2 = color.New(color.FgHiWhite)
	theme.Accent3 = color.New(color.FgWhite)
	theme.Border = color.New(color.FgWhite)
	theme.Selected = color.New(color.FgHiWhite, color.Underline)
	return theme
}