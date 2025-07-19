// Package cmdux provides a powerful and elegant command-line user experience library.
//
// cmdux aims to make building beautiful CLI applications as simple and enjoyable as possible.
// It provides a comprehensive set of UI components, animations, and interactive elements
// that can be easily composed to create stunning terminal applications.
//
// # Quick Start
//
//	app := cmdux.New()
//	box := ui.NewBox().Title("Hello").Content("Welcome to cmdux!")
//	app.Render(box)
//
// # Architecture
//
// cmdux is built around several core concepts:
//   - App: The main application context and renderer
//   - Theme: Color schemes and styling
//   - Components: Reusable UI elements
//   - Effects: Visual animations and transitions
//
// Each component follows a fluent API design pattern, making it easy to chain
// method calls for configuration.
package cmdux

import (
	"fmt"
	"io"
	"os"

	"github.com/bagaking/cmdux/core"
	"github.com/bagaking/cmdux/style"
)

// App represents the main cmdux application context.
// It manages themes, rendering, and global state.
type App struct {
	theme  *style.Theme
	writer io.Writer
	config *Config
}

// Config holds configuration options for the cmdux application.
type Config struct {
	// Writer specifies where output should be written. Defaults to os.Stdout.
	Writer io.Writer
	
	// Theme specifies the color theme to use. Defaults to DefaultTheme.
	Theme *style.Theme
	
	// Width specifies the terminal width. If 0, will auto-detect.
	Width int
	
	// EnableColors enables or disables color output. Auto-detected by default.
	EnableColors *bool
}

// New creates a new cmdux application with default settings.
func New(options ...func(*Config)) *App {
	config := &Config{
		Writer: os.Stdout,
		Theme:  style.DefaultTheme(),
	}
	
	for _, option := range options {
		option(config)
	}
	
	return &App{
		theme:  config.Theme,
		writer: config.Writer,
		config: config,
	}
}

// WithTheme sets a custom theme for the application.
func WithTheme(theme *style.Theme) func(*Config) {
	return func(c *Config) {
		c.Theme = theme
	}
}

// WithWriter sets a custom writer for output.
func WithWriter(w io.Writer) func(*Config) {
	return func(c *Config) {
		c.Writer = w
	}
}

// Theme returns the current theme being used by the application.
func (a *App) Theme() *style.Theme {
	return a.theme
}

// Render renders any component that implements the Renderable interface.
func (a *App) Render(component core.Renderable) error {
	output := component.Render(a.theme)
	_, err := fmt.Fprint(a.writer, output)
	return err
}

// Print is a convenience method for printing strings with theme colors.
func (a *App) Print(text string, colorFunc ...*style.Color) {
	if len(colorFunc) > 0 {
		fmt.Fprint(a.writer, colorFunc[0].Sprint(text))
	} else {
		fmt.Fprint(a.writer, text)
	}
}

// Println is like Print but adds a newline.
func (a *App) Println(text string, colorFunc ...*style.Color) {
	a.Print(text+"\n", colorFunc...)
}

// Clear clears the terminal screen.
func (a *App) Clear() {
	fmt.Fprint(a.writer, "\033[2J\033[H")
}

// MoveCursor moves the cursor to the specified position.
func (a *App) MoveCursor(x, y int) {
	fmt.Fprintf(a.writer, "\033[%d;%dH", y, x)
}

// Version returns the current version of cmdux.
func Version() string {
	return "1.0.0"
}