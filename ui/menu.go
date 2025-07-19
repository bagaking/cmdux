// Package ui provides menu components.
package ui

import (
	"strings"

	"github.com/bagaking/cmdux/core"
	"github.com/bagaking/cmdux/style"
	"github.com/mattn/go-runewidth"
)

// Menu represents an interactive menu component.
type Menu struct {
	*core.Component
	title       string
	options     []string
	descriptions []string
	selected    int
	prefix      string
	selectedPrefix string
	titleStyle  *style.Color
	optionStyle *style.Color
	selectedStyle *style.Color
	descStyle   *style.Color
}

// NewMenu creates a new menu component.
func NewMenu() *Menu {
	return &Menu{
		Component:      core.NewComponent(),
		selected:       0,
		prefix:         "  ",
		selectedPrefix: "▶ ",
	}
}

// Title sets the menu title.
func (m *Menu) Title(title string) *Menu {
	m.title = title
	return m
}

// Options sets the menu options.
func (m *Menu) Options(options ...string) *Menu {
	m.options = options
	return m
}

// OptionsWithDesc sets options with descriptions.
func (m *Menu) OptionsWithDesc(optionsDesc map[string]string) *Menu {
	m.options = make([]string, 0, len(optionsDesc))
	m.descriptions = make([]string, 0, len(optionsDesc))
	
	for option, desc := range optionsDesc {
		m.options = append(m.options, option)
		m.descriptions = append(m.descriptions, desc)
	}
	return m
}

// Selected sets the currently selected option index.
func (m *Menu) Selected(index int) *Menu {
	if index >= 0 && index < len(m.options) {
		m.selected = index
	}
	return m
}

// Prefix sets the prefix for unselected options.
func (m *Menu) Prefix(prefix string) *Menu {
	m.prefix = prefix
	return m
}

// SelectedPrefix sets the prefix for the selected option.
func (m *Menu) SelectedPrefix(prefix string) *Menu {
	m.selectedPrefix = prefix
	return m
}

// TitleStyle sets the title color.
func (m *Menu) TitleStyle(color *style.Color) *Menu {
	m.titleStyle = color
	return m
}

// OptionStyle sets the option color.
func (m *Menu) OptionStyle(color *style.Color) *Menu {
	m.optionStyle = color
	return m
}

// SelectedStyle sets the selected option color.
func (m *Menu) SelectedStyle(color *style.Color) *Menu {
	m.selectedStyle = color
	return m
}

// DescStyle sets the description color.
func (m *Menu) DescStyle(color *style.Color) *Menu {
	m.descStyle = color
	return m
}

// Render renders the menu using the given theme.
func (m *Menu) Render(theme *style.Theme) string {
	if m.IsHidden() || len(m.options) == 0 {
		return ""
	}

	titleColor := m.titleStyle
	if titleColor == nil {
		titleColor = theme.Header
	}

	optionColor := m.optionStyle
	if optionColor == nil {
		optionColor = theme.Primary
	}

	selectedColor := m.selectedStyle
	if selectedColor == nil {
		selectedColor = theme.Selected
	}

	descColor := m.descStyle
	if descColor == nil {
		descColor = theme.Muted
	}

	var result []string

	// Add title if present
	if m.title != "" {
		result = append(result, titleColor.Sprint(m.title))
		result = append(result, "") // Empty line
	}

	// Calculate widths for alignment
	maxOptionWidth := 0
	for _, option := range m.options {
		width := runewidth.StringWidth(option)
		if width > maxOptionWidth {
			maxOptionWidth = width
		}
	}

	// Add options
	for i, option := range m.options {
		var line string
		var desc string
		
		if i < len(m.descriptions) {
			desc = m.descriptions[i]
		}

		if i == m.selected {
			// Selected option
			line = selectedColor.Sprint(m.selectedPrefix + option)
			if desc != "" {
				// Pad option to align descriptions
				optionPadding := maxOptionWidth - runewidth.StringWidth(option)
				line += strings.Repeat(" ", optionPadding + 2) // 2 extra spaces
				line += descColor.Sprint(desc)
			}
		} else {
			// Regular option
			line = optionColor.Sprint(m.prefix + option)
			if desc != "" {
				// Pad option to align descriptions
				optionPadding := maxOptionWidth - runewidth.StringWidth(option)
				line += strings.Repeat(" ", optionPadding + 2) // 2 extra spaces
				line += descColor.Sprint(desc)
			}
		}

		result = append(result, line)
	}

	return strings.Join(result, "\n")
}

// GetSelected returns the currently selected option index.
func (m *Menu) GetSelected() int {
	return m.selected
}

// GetSelectedOption returns the currently selected option text.
func (m *Menu) GetSelectedOption() string {
	if m.selected >= 0 && m.selected < len(m.options) {
		return m.options[m.selected]
	}
	return ""
}

// SelectNext moves selection to the next option.
func (m *Menu) SelectNext() *Menu {
	if len(m.options) > 0 {
		m.selected = (m.selected + 1) % len(m.options)
	}
	return m
}

// SelectPrev moves selection to the previous option.
func (m *Menu) SelectPrev() *Menu {
	if len(m.options) > 0 {
		m.selected = (m.selected - 1 + len(m.options)) % len(m.options)
	}
	return m
}

// SelectByIndex sets the selected option by index.
func (m *Menu) SelectByIndex(index int) *Menu {
	if index >= 0 && index < len(m.options) {
		m.selected = index
	}
	return m
}

// SelectByOption sets the selected option by matching the option text.
func (m *Menu) SelectByOption(option string) *Menu {
	for i, opt := range m.options {
		if opt == option {
			m.selected = i
			break
		}
	}
	return m
}