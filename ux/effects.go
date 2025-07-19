// Package ux provides visual effects and animations.
package ux

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/bagaking/cmdux/style"
)

// TypewriterEffect displays text character by character with a typewriter effect.
func TypewriterEffect(text string, delay time.Duration, color ...*style.Color) {
	textColor := style.Primary
	if len(color) > 0 {
		textColor = color[0]
	}
	
	for _, char := range text {
		fmt.Print(textColor.Sprint(string(char)))
		time.Sleep(delay)
	}
	fmt.Println()
}

// MatrixEffect creates a matrix-style rain effect.
func MatrixEffect(duration time.Duration) {
	width, height := 80, 15
	chars := "アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲン0123456789"

	drops := make([]struct{ x, y, speed int }, width)
	for i := range drops {
		drops[i] = struct{ x, y, speed int }{
			x:     i,
			y:     rand.Intn(height),
			speed: 1 + rand.Intn(3),
		}
	}

	startTime := time.Now()
	for time.Since(startTime) < duration {
		frame := make([][]rune, height)
		for i := range frame {
			frame[i] = []rune(strings.Repeat(" ", width))
		}

		// Update and draw drops
		for i, drop := range drops {
			drop.y += drop.speed
			if drop.y >= height {
				drop.y = 0
				drop.x = rand.Intn(width)
			}
			drops[i] = drop

			for y := 0; y < height; y++ {
				if y >= drop.y-5 && y <= drop.y {
					charIndex := rand.Intn(len(chars))
					char := rune(chars[charIndex])
					if drop.x < width && y >= 0 {
						frame[y][drop.x] = char
					}
				}
			}
		}

		fmt.Print("\033[2J\033[H") // Clear screen
		for y, line := range frame {
			for x, char := range line {
				if char != ' ' {
					// Color based on position for trail effect
					if y > drops[x%len(drops)].y-2 {
						style.Success.Print(string(char))
					} else {
						style.Muted.Print(string(char))
					}
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Print("\033[2J\033[H") // Clear screen
}

// WaveEffect creates a wave animation with text.
func WaveEffect(text string, duration time.Duration, color ...*style.Color) {
	textColor := style.Primary
	if len(color) > 0 {
		textColor = color[0]
	}
	
	width := 80
	height := 5
	startTime := time.Now()

	for time.Since(startTime) < duration {
		frame := make([]string, height)
		for i := range frame {
			frame[i] = strings.Repeat(" ", width)
		}

		// Create wave pattern
		for x := 0; x < len(text) && x < width; x++ {
			y := int(2 + 1.5*math.Sin(float64(x)*0.5+float64(time.Since(startTime).Milliseconds())*0.01))
			if y >= 0 && y < height {
				row := []rune(frame[y])
				if x < len(row) {
					row[x] = rune(text[x%len(text)])
					frame[y] = string(row)
				}
			}
		}

		// Clear screen and print frame
		fmt.Print("\033[2J\033[H")
		for _, line := range frame {
			textColor.Println(line)
		}
		time.Sleep(50 * time.Millisecond)
	}

	// Reset cursor position
	fmt.Print("\033[H")
}

// GlitchEffect creates a glitch-style text effect.
func GlitchEffect(text string, duration time.Duration, color ...*style.Color) {
	glitchColor := style.Error
	normalColor := style.Primary
	if len(color) > 0 {
		normalColor = color[0]
	}
	
	glitchChars := "$#@!%^*&*()_+-=[]{}|;:,.<>?"
	startTime := time.Now()

	for time.Since(startTime) < duration {
		fmt.Print("\033[2K\r") // Clear line

		glitched := ""
		for _, char := range text {
			if rand.Float32() < 0.1 {
				glitched += string(glitchChars[rand.Intn(len(glitchChars))])
			} else {
				glitched += string(char)
			}
		}

		if rand.Float32() < 0.3 {
			glitchColor.Printf("%s", glitched)
		} else {
			normalColor.Printf("%s", glitched)
		}

		time.Sleep(100 * time.Millisecond)
	}
	
	// Show final clean text
	fmt.Print("\033[2K\r")
	normalColor.Println(text)
}

// PulseEffect creates a pulsing color effect.
func PulseEffect(text string, duration time.Duration, colors ...*style.Color) {
	if len(colors) == 0 {
		colors = []*style.Color{style.Primary, style.Secondary, style.Accent1}
	}
	
	startTime := time.Now()
	i := 0
	
	for time.Since(startTime) < duration {
		fmt.Print("\033[2K\r") // Clear line
		color := colors[i%len(colors)]
		color.Print(text)
		time.Sleep(200 * time.Millisecond)
		i++
	}
	
	fmt.Print("\033[2K\r")
	style.Primary.Println(text)
}

// FadeInEffect creates a fade-in effect by gradually increasing brightness.
func FadeInEffect(text string, steps int, stepDelay time.Duration) {
	colors := []*style.Color{
		style.Faint,
		style.Muted,
		style.Secondary,
		style.Primary,
	}
	
	for i := 0; i < steps && i < len(colors); i++ {
		fmt.Print("\033[2K\r") // Clear line
		colors[i].Print(text)
		time.Sleep(stepDelay)
	}
	fmt.Println()
}

// RainbowEffect displays text with rainbow colors.
func RainbowEffect(text string) {
	colors := []*style.Color{
		style.Error,    // Red
		style.Warning,  // Yellow
		style.Success,  // Green
		style.Primary,  // Cyan
		style.Secondary, // Blue
		style.Accent1,  // Magenta
	}

	for i, char := range text {
		if char != ' ' {
			colors[i%len(colors)].Print(string(char))
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

// BreathingEffect creates a breathing pulse effect.
func BreathingEffect(text string, duration time.Duration, color ...*style.Color) {
	textColor := style.Success
	if len(color) > 0 {
		textColor = color[0]
	}
	
	startTime := time.Now()
	for time.Since(startTime) < duration {
		// Create breathing effect
		fmt.Print("\033[2K\r")
		textColor.Printf("%s", text)
		time.Sleep(500 * time.Millisecond)
		
		fmt.Print("\033[2K\r")
		style.Muted.Printf("%s", text)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Print("\033[2K\r")
	textColor.Println(text)
}

// LoadingDots creates animated loading dots.
func LoadingDots(text string, duration time.Duration, color ...*style.Color) {
	textColor := style.Primary
	if len(color) > 0 {
		textColor = color[0]
	}
	
	dots := []string{"", ".", "..", "..."}
	startTime := time.Now()
	i := 0
	
	for time.Since(startTime) < duration {
		fmt.Print("\033[2K\r") // Clear line
		textColor.Printf("%s%s", text, dots[i%len(dots)])
		time.Sleep(300 * time.Millisecond)
		i++
	}
	
	fmt.Print("\033[2K\r")
	textColor.Println(text)
}