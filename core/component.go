// Package core provides the fundamental interfaces and types for cmdux.
package core

import "github.com/bagaking/cmdux/style"

// Renderable represents any component that can be rendered to the terminal.
type Renderable interface {
	// Render returns the string representation of the component using the given theme.
	Render(theme *style.Theme) string
}

// Component represents a basic UI component with common properties.
type Component struct {
	width  int
	height int
	hidden bool
	style  *style.Style
}

// NewComponent creates a new base component.
func NewComponent() *Component {
	return &Component{
		width:  -1, // Auto-size
		height: -1, // Auto-size
		hidden: false,
	}
}

// Width sets the component width.
func (c *Component) Width(w int) *Component {
	c.width = w
	return c
}

// Height sets the component height.
func (c *Component) Height(h int) *Component {
	c.height = h
	return c
}

// Hide hides the component.
func (c *Component) Hide() *Component {
	c.hidden = true
	return c
}

// Show shows the component.
func (c *Component) Show() *Component {
	c.hidden = false
	return c
}

// GetWidth returns the component width.
func (c *Component) GetWidth() int {
	return c.width
}

// GetHeight returns the component height.
func (c *Component) GetHeight() int {
	return c.height
}

// IsHidden returns whether the component is hidden.
func (c *Component) IsHidden() bool {
	return c.hidden
}

// SetStyle sets the component style.
func (c *Component) SetStyle(s *style.Style) *Component {
	c.style = s
	return c
}

// GetStyle returns the component style.
func (c *Component) GetStyle() *style.Style {
	return c.style
}