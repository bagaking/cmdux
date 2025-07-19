// Package ux provides progress bar components.
package ux

import (
	"fmt"
	"strings"

	"github.com/bagaking/cmdux/core"
	"github.com/bagaking/cmdux/style"
)

// ProgressBar represents a progress indicator.
type ProgressBar struct {
	*core.Component
	current   int
	total     int
	width     int
	prefix    string
	suffix    string
	completed bool
	showPercent bool
	showNumbers bool
	fillChar    string
	emptyChar   string
	leftCap     string
	rightCap    string
	color       *style.Color
	bgColor     *style.Color
}

// NewProgressBar creates a new progress bar.
func NewProgressBar(width int) *ProgressBar {
	return &ProgressBar{
		Component:   core.NewComponent(),
		width:       width,
		prefix:      "Progress",
		fillChar:    "█",
		emptyChar:   "░",
		leftCap:     "[",
		rightCap:    "]",
		showPercent: true,
		showNumbers: true,
		color:       style.Primary,
		bgColor:     style.Muted,
	}
}

// SetTotal sets the total value for the progress bar.
func (pb *ProgressBar) SetTotal(total int) *ProgressBar {
	pb.total = total
	return pb
}

// SetPrefix sets the prefix text shown before the progress bar.
func (pb *ProgressBar) SetPrefix(prefix string) *ProgressBar {
	pb.prefix = prefix
	return pb
}

// SetSuffix sets the suffix text shown after the progress bar.
func (pb *ProgressBar) SetSuffix(suffix string) *ProgressBar {
	pb.suffix = suffix
	return pb
}

// ShowPercent enables or disables percentage display.
func (pb *ProgressBar) ShowPercent(show bool) *ProgressBar {
	pb.showPercent = show
	return pb
}

// ShowNumbers enables or disables current/total number display.
func (pb *ProgressBar) ShowNumbers(show bool) *ProgressBar {
	pb.showNumbers = show
	return pb
}

// SetChars sets the characters used for the progress bar.
func (pb *ProgressBar) SetChars(fill, empty, leftCap, rightCap string) *ProgressBar {
	pb.fillChar = fill
	pb.emptyChar = empty
	pb.leftCap = leftCap
	pb.rightCap = rightCap
	return pb
}

// Color sets the progress bar color.
func (pb *ProgressBar) Color(color *style.Color) *ProgressBar {
	pb.color = color
	return pb
}

// BgColor sets the background color.
func (pb *ProgressBar) BgColor(color *style.Color) *ProgressBar {
	pb.bgColor = color
	return pb
}

// Update updates the current progress value.
func (pb *ProgressBar) Update(current int) {
	pb.current = current
	fmt.Print("\r" + pb.Render())
}

// Complete marks the progress as complete and shows a completion message.
func (pb *ProgressBar) Complete(message string) {
	pb.current = pb.total
	pb.completed = true
	fmt.Print("\r" + pb.Render())
	if message != "" {
		fmt.Printf("\n%s %s\n", style.Success.Sprint("✓"), message)
	} else {
		fmt.Println()
	}
}

// Render renders the progress bar as a string.
func (pb *ProgressBar) Render() string {
	if pb.total == 0 {
		return pb.prefix + " [indeterminate]"
	}

	percentage := float64(pb.current) / float64(pb.total) * 100
	filledWidth := int(float64(pb.width) * float64(pb.current) / float64(pb.total))
	emptyWidth := pb.width - filledWidth

	// Build the progress bar
	var bar strings.Builder
	
	// Add filled portion
	if filledWidth > 0 {
		bar.WriteString(pb.color.Sprint(strings.Repeat(pb.fillChar, filledWidth)))
	}
	
	// Add empty portion
	if emptyWidth > 0 {
		bar.WriteString(pb.bgColor.Sprint(strings.Repeat(pb.emptyChar, emptyWidth)))
	}

	// Build the complete display
	var result strings.Builder
	
	// Prefix
	if pb.prefix != "" {
		result.WriteString(pb.prefix + " ")
	}
	
	// Progress bar with caps
	result.WriteString(pb.leftCap)
	result.WriteString(bar.String())
	result.WriteString(pb.rightCap)
	
	// Percentage
	if pb.showPercent {
		result.WriteString(fmt.Sprintf(" %.1f%%", percentage))
	}
	
	// Numbers
	if pb.showNumbers {
		result.WriteString(fmt.Sprintf(" (%d/%d)", pb.current, pb.total))
	}
	
	// Suffix
	if pb.suffix != "" {
		result.WriteString(" " + pb.suffix)
	}

	return result.String()
}

// SetCurrent sets the current progress value.
func (pb *ProgressBar) SetCurrent(current int) *ProgressBar {
	pb.current = current
	return pb
}

// GetCurrent returns the current progress value.
func (pb *ProgressBar) GetCurrent() int {
	return pb.current
}

// GetTotal returns the total progress value.
func (pb *ProgressBar) GetTotal() int {
	return pb.total
}

// GetPercentage returns the current percentage (0-100).
func (pb *ProgressBar) GetPercentage() float64 {
	if pb.total == 0 {
		return 0
	}
	return float64(pb.current) / float64(pb.total) * 100
}

// IsComplete returns true if the progress is complete.
func (pb *ProgressBar) IsComplete() bool {
	return pb.completed || pb.current >= pb.total
}

// Increment increments the current progress by 1.
func (pb *ProgressBar) Increment() {
	pb.current++
	if pb.current > pb.total {
		pb.current = pb.total
	}
}

// IncrementBy increments the current progress by the specified amount.
func (pb *ProgressBar) IncrementBy(amount int) {
	pb.current += amount
	if pb.current > pb.total {
		pb.current = pb.total
	}
}