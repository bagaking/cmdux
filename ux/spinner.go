// Package ux provides user experience components like spinners and animations.
package ux

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/bagaking/cmdux/style"
)

// Spinner represents an animated loading spinner.
type Spinner struct {
	frames []string
	color  *style.Color
	stop   chan bool
	text   string
	delay  time.Duration
}

// SpinnerStyle represents different spinner animation styles.
type SpinnerStyle string

const (
	SpinnerDots    SpinnerStyle = "dots"
	SpinnerCircle  SpinnerStyle = "circle"
	SpinnerArrows  SpinnerStyle = "arrows"
	SpinnerBounce  SpinnerStyle = "bounce"
	SpinnerPulse   SpinnerStyle = "pulse"
	SpinnerBlocks  SpinnerStyle = "blocks"
	SpinnerWaves   SpinnerStyle = "waves"
	SpinnerMatrix  SpinnerStyle = "matrix"
)

// Animation frames for different spinner styles
var spinnerFrames = map[SpinnerStyle][]string{
	SpinnerDots:    {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	SpinnerCircle:  {"◐", "◓", "◑", "◒"},
	SpinnerArrows:  {"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
	SpinnerBounce:  {"⠁", "⠂", "⠄", "⠂"},
	SpinnerPulse:   {"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃"},
	SpinnerBlocks:  {"▖", "▘", "▝", "▗"},
	SpinnerWaves:   {"▂", "▄", "▅", "▆", "▇", "▆", "▅", "▄"},
	SpinnerMatrix:  {"ｦ", "ｧ", "ｨ", "ｩ", "ｪ", "ｫ", "ｬ", "ｭ", "ｮ", "ｯ"},
}

// NewSpinner creates a new spinner with the specified style.
func NewSpinner(spinnerStyle SpinnerStyle) *Spinner {
	frames, exists := spinnerFrames[spinnerStyle]
	if !exists {
		frames = spinnerFrames[SpinnerDots] // Default fallback
	}

	return &Spinner{
		frames: frames,
		color:  style.Primary,
		stop:   make(chan bool),
		delay:  100 * time.Millisecond,
	}
}

// Color sets the spinner color.
func (s *Spinner) Color(color *style.Color) *Spinner {
	s.color = color
	return s
}

// Delay sets the animation delay between frames.
func (s *Spinner) Delay(delay time.Duration) *Spinner {
	s.delay = delay
	return s
}

// Start starts the spinner animation with the given text.
func (s *Spinner) Start(text string) {
	s.text = text
	go func() {
		i := 0
		for {
			select {
			case <-s.stop:
				return
			default:
				frame := s.frames[i%len(s.frames)]
				fmt.Printf("\r%s %s", s.color.Sprint(frame), s.text)
				time.Sleep(s.delay)
				i++
			}
		}
	}()
}

// Stop stops the spinner animation and clears the line.
func (s *Spinner) Stop() {
	close(s.stop)
	fmt.Print("\r")
	fmt.Print(strings.Repeat(" ", utf8.RuneCountInString(s.text)+3))
	fmt.Print("\r")
}

// Success stops the spinner and shows a success message.
func (s *Spinner) Success(message string) {
	s.Stop()
	fmt.Printf("\r%s %s\n", style.Success.Sprint("✓"), message)
}

// Error stops the spinner and shows an error message.
func (s *Spinner) Error(message string) {
	s.Stop()
	fmt.Printf("\r%s %s\n", style.Error.Sprint("✗"), message)
}

// Warning stops the spinner and shows a warning message.
func (s *Spinner) Warning(message string) {
	s.Stop()
	fmt.Printf("\r%s %s\n", style.Warning.Sprint("⚠"), message)
}

// Info stops the spinner and shows an info message.
func (s *Spinner) Info(message string) {
	s.Stop()
	fmt.Printf("\r%s %s\n", style.Primary.Sprint("ℹ"), message)
}

// Update updates the spinner text without restarting the animation.
func (s *Spinner) Update(text string) {
	s.text = text
}