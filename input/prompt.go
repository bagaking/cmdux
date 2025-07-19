// Package input provides interactive input components.
package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bagaking/cmdux/style"
)

// Prompt represents an interactive user prompt.
type Prompt struct {
	message     string
	defaultValue string
	validator   func(string) error
	transformer func(string) string
	required    bool
	hidden      bool // For password input
	prefix      string
	style       *style.Color
	errorStyle  *style.Color
}

// NewPrompt creates a new prompt.
func NewPrompt(message string) *Prompt {
	return &Prompt{
		message:    message,
		prefix:     "? ",
		style:      style.Primary,
		errorStyle: style.Error,
	}
}

// Default sets the default value for the prompt.
func (p *Prompt) Default(value string) *Prompt {
	p.defaultValue = value
	return p
}

// Required makes the prompt require a non-empty input.
func (p *Prompt) Required(required bool) *Prompt {
	p.required = required
	return p
}

// Hidden makes the input hidden (for passwords).
func (p *Prompt) Hidden(hidden bool) *Prompt {
	p.hidden = hidden
	return p
}

// Validator sets a validation function.
func (p *Prompt) Validator(validator func(string) error) *Prompt {
	p.validator = validator
	return p
}

// Transformer sets a transformation function applied to the input.
func (p *Prompt) Transformer(transformer func(string) string) *Prompt {
	p.transformer = transformer
	return p
}

// Prefix sets the prompt prefix.
func (p *Prompt) Prefix(prefix string) *Prompt {
	p.prefix = prefix
	return p
}

// Style sets the prompt color.
func (p *Prompt) Style(color *style.Color) *Prompt {
	p.style = color
	return p
}

// ErrorStyle sets the error message color.
func (p *Prompt) ErrorStyle(color *style.Color) *Prompt {
	p.errorStyle = color
	return p
}

// Run executes the prompt and returns the user input.
func (p *Prompt) Run() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		// Display the prompt
		p.displayPrompt()
		
		// Read input
		var input string
		var err error
		
		if p.hidden {
			// TODO: Implement hidden input (password)
			// For now, use regular input
			input, err = reader.ReadString('\n')
		} else {
			input, err = reader.ReadString('\n')
		}
		
		if err != nil {
			return "", err
		}
		
		// Trim newline
		input = strings.TrimSpace(input)
		
		// Use default if empty
		if input == "" && p.defaultValue != "" {
			input = p.defaultValue
		}
		
		// Check required
		if p.required && input == "" {
			p.errorStyle.Println("✗ This field is required")
			continue
		}
		
		// Apply transformer
		if p.transformer != nil {
			input = p.transformer(input)
		}
		
		// Validate
		if p.validator != nil {
			if err := p.validator(input); err != nil {
				p.errorStyle.Printf("✗ %s\n", err.Error())
				continue
			}
		}
		
		return input, nil
	}
}

func (p *Prompt) displayPrompt() {
	prompt := p.style.Sprint(p.prefix + p.message)
	
	if p.defaultValue != "" {
		prompt += style.Muted.Sprintf(" (%s)", p.defaultValue)
	}
	
	if p.required {
		prompt += style.Error.Sprint(" *")
	}
	
	prompt += ": "
	fmt.Print(prompt)
}

// Confirm creates a yes/no confirmation prompt.
func Confirm(message string, defaultValue ...bool) (bool, error) {
	defaultVal := false
	if len(defaultValue) > 0 {
		defaultVal = defaultValue[0]
	}
	
	prompt := style.Primary.Sprint("? " + message)
	
	if defaultVal {
		prompt += style.Muted.Sprint(" (Y/n)")
	} else {
		prompt += style.Muted.Sprint(" (y/N)")
	}
	
	prompt += ": "
	fmt.Print(prompt)
	
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	
	input = strings.TrimSpace(strings.ToLower(input))
	
	if input == "" {
		return defaultVal, nil
	}
	
	return input == "y" || input == "yes", nil
}

// Select creates a selection prompt from a list of options.
func Select(message string, options []string) (int, string, error) {
	if len(options) == 0 {
		return -1, "", fmt.Errorf("no options provided")
	}
	
	// Display options
	fmt.Println(style.Primary.Sprint("? " + message))
	for i, option := range options {
		fmt.Printf("  %d) %s\n", i+1, option)
	}
	
	// Get selection
	fmt.Print(style.Primary.Sprint("Enter choice (1-" + strconv.Itoa(len(options)) + "): "))
	
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return -1, "", err
	}
	
	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return -1, "", fmt.Errorf("invalid choice: %s", input)
	}
	
	if choice < 1 || choice > len(options) {
		return -1, "", fmt.Errorf("choice must be between 1 and %d", len(options))
	}
	
	return choice - 1, options[choice-1], nil
}

// MultiSelect creates a multi-selection prompt.
func MultiSelect(message string, options []string) ([]int, []string, error) {
	if len(options) == 0 {
		return nil, nil, fmt.Errorf("no options provided")
	}
	
	// Display options
	fmt.Println(style.Primary.Sprint("? " + message + " (comma-separated numbers)"))
	for i, option := range options {
		fmt.Printf("  %d) %s\n", i+1, option)
	}
	
	// Get selections
	fmt.Print(style.Primary.Sprint("Enter choices: "))
	
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, nil, err
	}
	
	input = strings.TrimSpace(input)
	if input == "" {
		return []int{}, []string{}, nil
	}
	
	parts := strings.Split(input, ",")
	var indices []int
	var selected []string
	
	for _, part := range parts {
		choice, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, nil, fmt.Errorf("invalid choice: %s", part)
		}
		
		if choice < 1 || choice > len(options) {
			return nil, nil, fmt.Errorf("choice must be between 1 and %d", len(options))
		}
		
		indices = append(indices, choice-1)
		selected = append(selected, options[choice-1])
	}
	
	return indices, selected, nil
}

// Password creates a hidden password input prompt.
func Password(message string) (string, error) {
	prompt := NewPrompt(message).
		Hidden(true).
		Required(true)
	
	return prompt.Run()
}