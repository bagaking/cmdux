// Package main demonstrates cmdux usage examples.
package main

import (
	"strings"
	"time"

	"github.com/bagaking/cmdux"
	"github.com/bagaking/cmdux/input"
	"github.com/bagaking/cmdux/style"
	"github.com/bagaking/cmdux/ui"
	"github.com/bagaking/cmdux/ux"
)

func main() {
	// Create a new cmdux application
	app := cmdux.New()

	// Example 1: Simple box
	app.Println("=== Box Example ===", app.Theme().Header)
	box := ui.NewBox().
		Title("Welcome to cmdux").
		Content("This is a beautiful terminal UI library for Go!\n\nFeatures:\n• Rich theming\n• Beautiful components\n• Smooth animations").
		Width(50).
		TitleStyle(app.Theme().Primary).
		ContentStyle(app.Theme().Secondary)
	app.Render(box)
	app.Println("")

	// Example 2: Table
	app.Println("=== Table Example ===", app.Theme().Header)
	table := ui.NewTable().
		Headers("Name", "Role", "Experience").
		AddRow("Alice", "Developer", "5 years").
		AddRow("Bob", "Designer", "3 years").
		AddRow("Charlie", "Manager", "8 years").
		HeaderStyle(app.Theme().Header).
		RowStyle(app.Theme().Primary).
		AltRowStyle(app.Theme().Secondary)
	app.Render(table)
	app.Println("")

	// Example 3: Menu
	app.Println("=== Menu Example ===", app.Theme().Header)
	menu := ui.NewMenu().
		Title("Choose an action:").
		Options("Create new project", "Open existing", "Settings", "Exit").
		TitleStyle(app.Theme().Header).
		SelectedStyle(app.Theme().Selected)
	app.Render(menu)
	app.Println("")

	// Example 4: Spinner animation
	app.Println("=== Spinner Example ===", app.Theme().Header)
	spinner := ux.NewSpinner(ux.SpinnerDots).Color(app.Theme().Primary)
	spinner.Start("Loading awesome features...")
	time.Sleep(2 * time.Second)
	spinner.Success("All features loaded!")
	app.Println("")

	// Example 5: Progress bar
	app.Println("=== Progress Bar Example ===", app.Theme().Header)
	progress := ux.NewProgressBar(30).
		SetTotal(100).
		SetPrefix("Installing").
		Color(app.Theme().Success).
		BgColor(app.Theme().Muted)

	for i := 0; i <= 100; i += 10 {
		progress.Update(i)
		time.Sleep(200 * time.Millisecond)
	}
	progress.Complete("Installation complete!")
	app.Println("")

	// Example 6: Visual effects
	app.Println("=== Effects Example ===", app.Theme().Header)
	ux.TypewriterEffect("This text appears character by character...", 50*time.Millisecond, app.Theme().Primary)
	ux.RainbowEffect("🌈 This text has rainbow colors! 🌈")
	app.Println("")

	// Example 7: Interactive input
	app.Println("=== Input Example ===", app.Theme().Header)
	
	// Simple prompt
	prompt := input.NewPrompt("What's your name?").
		Default("Anonymous").
		Required(true)
	name, _ := prompt.Run()
	app.Println("Hello, "+name+"!", app.Theme().Success)

	// Confirmation
	confirmed, _ := input.Confirm("Do you like cmdux?", true)
	if confirmed {
		app.Println("Great! We're glad you like it! 🎉", app.Theme().Success)
	} else {
		app.Println("We'll keep improving! 💪", app.Theme().Warning)
	}

	// Selection
	_, choice, _ := input.Select("What's your favorite theme?", []string{
		"Default", "Dark", "Light", "Cyberpunk", "Monochrome",
	})
	app.Println("You chose: "+choice, app.Theme().Primary)

	// Example 8: Form
	app.Println("")
	app.Println("=== Form Example ===", app.Theme().Header)
	form := input.NewForm("User Registration").
		TextField("username", "Username", true).
		TextField("email", "Email", true).
		NumberField("age", "Age", false, 25).
		BooleanField("newsletter", "Subscribe to newsletter?", true).
		SelectField("role", "Role", []string{"Developer", "Designer", "Manager"}, true)

	results, _ := form.Run()
	
	app.Println("")
	app.Println("Registration complete! Here's your info:", app.Theme().Success)
	infoBox := ui.NewBox().
		Title("User Information").
		Content(formatUserInfo(results)).
		BorderStyle(app.Theme().Success)
	app.Render(infoBox)

	// Example 9: Theme showcase
	app.Println("")
	app.Println("=== Theme Showcase ===", app.Theme().Header)
	
	themes := map[string]*style.Theme{
		"Default":    style.DefaultTheme(),
		"Dark":       style.DarkTheme(),
		"Cyberpunk":  style.CyberpunkTheme(),
		"Monochrome": style.MonochromeTheme(),
	}
	
	for name, theme := range themes {
		themeApp := cmdux.New(cmdux.WithTheme(theme))
		themeBox := ui.NewBox().
			Title(name + " Theme").
			Content("This is how the " + name + " theme looks!").
			Width(40)
		themeApp.Render(themeBox)
		app.Println("")
	}

	// Final message
	app.Println("🎉 cmdux demo complete! Build amazing CLI apps! 🚀", app.Theme().Success)
}

func formatUserInfo(results map[string]interface{}) string {
	var info []string
	
	if username, ok := results["username"].(string); ok {
		info = append(info, "Username: "+username)
	}
	if email, ok := results["email"].(string); ok {
		info = append(info, "Email: "+email)
	}
	if age, ok := results["age"].(int); ok && age > 0 {
		info = append(info, "Age: "+string(rune(age+'0')))
	}
	if newsletter, ok := results["newsletter"].(bool); ok {
		if newsletter {
			info = append(info, "Newsletter: Yes")
		} else {
			info = append(info, "Newsletter: No")
		}
	}
	if role, ok := results["role"].(string); ok {
		info = append(info, "Role: "+role)
	}
	
	return strings.Join(info, "\n")
}