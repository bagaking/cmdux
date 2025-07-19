// Package ui provides table components.
package ui

import (
	"fmt"
	"strings"

	"github.com/bagaking/cmdux/core"
	"github.com/bagaking/cmdux/style"
	"github.com/mattn/go-runewidth"
)

// Table represents a data table component.
type Table struct {
	*core.Component
	headers     []string
	rows        [][]string
	columnWidths []int
	border      bool
	borderStyle *style.Color
	headerStyle *style.Color
	rowStyle    *style.Color
	altRowStyle *style.Color
	alignment   []core.Alignment
}

// NewTable creates a new table component.
func NewTable() *Table {
	return &Table{
		Component: core.NewComponent(),
		border:    true,
		alignment: []core.Alignment{core.AlignLeft}, // Default alignment
	}
}

// Headers sets the table headers.
func (t *Table) Headers(headers ...string) *Table {
	t.headers = headers
	// Auto-initialize column widths and alignment if not set
	if len(t.columnWidths) == 0 {
		t.columnWidths = make([]int, len(headers))
		for i, header := range headers {
			t.columnWidths[i] = runewidth.StringWidth(header)
		}
	}
	if len(t.alignment) < len(headers) {
		for len(t.alignment) < len(headers) {
			t.alignment = append(t.alignment, core.AlignLeft)
		}
	}
	return t
}

// Rows sets the table data rows.
func (t *Table) Rows(rows ...[]string) *Table {
	t.rows = rows
	t.calculateColumnWidths()
	return t
}

// AddRow adds a single row to the table.
func (t *Table) AddRow(row ...string) *Table {
	t.rows = append(t.rows, row)
	t.updateColumnWidthsForRow(row)
	return t
}

// ColumnWidths sets explicit column widths.
func (t *Table) ColumnWidths(widths ...int) *Table {
	t.columnWidths = widths
	return t
}

// Border enables or disables table borders.
func (t *Table) Border(enabled bool) *Table {
	t.border = enabled
	return t
}

// BorderStyle sets the border color.
func (t *Table) BorderStyle(color *style.Color) *Table {
	t.borderStyle = color
	return t
}

// HeaderStyle sets the header row color.
func (t *Table) HeaderStyle(color *style.Color) *Table {
	t.headerStyle = color
	return t
}

// RowStyle sets the regular row color.
func (t *Table) RowStyle(color *style.Color) *Table {
	t.rowStyle = color
	return t
}

// AltRowStyle sets the alternating row color.
func (t *Table) AltRowStyle(color *style.Color) *Table {
	t.altRowStyle = color
	return t
}

// Alignment sets column alignments.
func (t *Table) Alignment(alignments ...core.Alignment) *Table {
	t.alignment = alignments
	return t
}

// Render renders the table using the given theme.
func (t *Table) Render(theme *style.Theme) string {
	if t.IsHidden() || len(t.headers) == 0 {
		return ""
	}

	borderColor := t.borderStyle
	if borderColor == nil {
		borderColor = theme.Border
	}

	headerColor := t.headerStyle
	if headerColor == nil {
		headerColor = theme.Header
	}

	rowColor := t.rowStyle
	if rowColor == nil {
		rowColor = theme.Primary
	}

	altRowColor := t.altRowStyle
	if altRowColor == nil {
		altRowColor = theme.Secondary
	}

	var result []string

	if t.border {
		// Top border
		result = append(result, t.renderTopBorder(borderColor))
		
		// Header row
		result = append(result, t.renderRow(t.headers, headerColor, borderColor, true))
		
		// Header separator
		result = append(result, t.renderSeparator(borderColor))
		
		// Data rows
		for i, row := range t.rows {
			var color *style.Color
			if i%2 == 0 {
				color = rowColor
			} else {
				color = altRowColor
			}
			result = append(result, t.renderRow(row, color, borderColor, false))
		}
		
		// Bottom border
		result = append(result, t.renderBottomBorder(borderColor))
	} else {
		// No border version
		result = append(result, t.renderRowNoBorder(t.headers, headerColor))
		result = append(result, strings.Repeat("-", t.getTotalWidth()))
		
		for i, row := range t.rows {
			var color *style.Color
			if i%2 == 0 {
				color = rowColor
			} else {
				color = altRowColor
			}
			result = append(result, t.renderRowNoBorder(row, color))
		}
	}

	return strings.Join(result, "\n")
}

func (t *Table) calculateColumnWidths() {
	if len(t.columnWidths) == 0 {
		t.columnWidths = make([]int, len(t.headers))
	}

	// Initialize with header widths
	for i, header := range t.headers {
		if i < len(t.columnWidths) {
			t.columnWidths[i] = runewidth.StringWidth(header)
		}
	}

	// Update with row data
	for _, row := range t.rows {
		t.updateColumnWidthsForRow(row)
	}
}

func (t *Table) updateColumnWidthsForRow(row []string) {
	for i, cell := range row {
		if i < len(t.columnWidths) {
			cellWidth := runewidth.StringWidth(cell)
			if cellWidth > t.columnWidths[i] {
				t.columnWidths[i] = cellWidth
			}
		}
	}
}

func (t *Table) getAlignment(colIndex int) core.Alignment {
	if colIndex < len(t.alignment) {
		return t.alignment[colIndex]
	}
	return core.AlignLeft
}

func (t *Table) renderTopBorder(borderColor *style.Color) string {
	var parts []string
	parts = append(parts, borderColor.Sprint(style.BoxTopLeft))
	
	for i, width := range t.columnWidths {
		if i > 0 {
			parts = append(parts, borderColor.Sprint(style.BoxTeeTop))
		}
		parts = append(parts, borderColor.Sprint(strings.Repeat(style.BoxHorizontal, width+2))) // +2 for padding
	}
	
	parts = append(parts, borderColor.Sprint(style.BoxTopRight))
	return strings.Join(parts, "")
}

func (t *Table) renderBottomBorder(borderColor *style.Color) string {
	var parts []string
	parts = append(parts, borderColor.Sprint(style.BoxBottomLeft))
	
	for i, width := range t.columnWidths {
		if i > 0 {
			parts = append(parts, borderColor.Sprint(style.BoxTeeBottom))
		}
		parts = append(parts, borderColor.Sprint(strings.Repeat(style.BoxHorizontal, width+2))) // +2 for padding
	}
	
	parts = append(parts, borderColor.Sprint(style.BoxBottomRight))
	return strings.Join(parts, "")
}

func (t *Table) renderSeparator(borderColor *style.Color) string {
	var parts []string
	parts = append(parts, borderColor.Sprint(style.BoxTee))
	
	for i, width := range t.columnWidths {
		if i > 0 {
			parts = append(parts, borderColor.Sprint(style.BoxCross))
		}
		parts = append(parts, borderColor.Sprint(strings.Repeat(style.BoxHorizontal, width+2))) // +2 for padding
	}
	
	parts = append(parts, borderColor.Sprint(style.BoxTeeRight))
	return strings.Join(parts, "")
}

func (t *Table) renderRow(cells []string, cellColor, borderColor *style.Color, isHeader bool) string {
	var parts []string
	parts = append(parts, borderColor.Sprint(style.BoxVertical))
	
	for i, width := range t.columnWidths {
		var cell string
		if i < len(cells) {
			cell = cells[i]
		}
		
		// Truncate if too long
		if runewidth.StringWidth(cell) > width {
			cell = runewidth.Truncate(cell, width, "…")
		}
		
		// Apply alignment
		alignment := t.getAlignment(i)
		renderer := core.NewRenderer(width, 1)
		paddedCell := renderer.PadText(cell, width, alignment)
		
		styledCell := cellColor.Sprint(paddedCell)
		parts = append(parts, fmt.Sprintf(" %s ", styledCell))
		parts = append(parts, borderColor.Sprint(style.BoxVertical))
	}
	
	return strings.Join(parts, "")
}

func (t *Table) renderRowNoBorder(cells []string, cellColor *style.Color) string {
	var parts []string
	
	for i, width := range t.columnWidths {
		var cell string
		if i < len(cells) {
			cell = cells[i]
		}
		
		// Truncate if too long
		if runewidth.StringWidth(cell) > width {
			cell = runewidth.Truncate(cell, width, "…")
		}
		
		// Apply alignment
		alignment := t.getAlignment(i)
		renderer := core.NewRenderer(width, 1)
		paddedCell := renderer.PadText(cell, width, alignment)
		
		styledCell := cellColor.Sprint(paddedCell)
		parts = append(parts, styledCell)
	}
	
	return strings.Join(parts, " ")
}

func (t *Table) getTotalWidth() int {
	total := 0
	for _, width := range t.columnWidths {
		total += width
	}
	// Add separators
	if len(t.columnWidths) > 1 {
		total += len(t.columnWidths) - 1
	}
	return total
}