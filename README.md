# cmdux

A powerful and elegant command-line user experience library for Go.

## Features

- 🎨 **Rich Theming**: Beautiful color schemes and customizable themes
- 🖼️ **UI Components**: Tables, boxes, menus, trees, and more  
- ✨ **Animations**: Spinners, progress bars, and visual effects
- 🎯 **Interactive**: User input, selections, and forms
- 🔧 **Composable**: Mix and match components as needed
- 🚀 **Zero Config**: Works great out of the box

## Quick Start

```go
package main

import (
    "github.com/bagaking/cmdux"
    "github.com/bagaking/cmdux/ui"
)

func main() {
    app := cmdux.New()
    
    // Create a beautiful header
    app.Render(ui.NewBox().
        Title("Welcome to cmdux").
        Content("Building amazing CLI experiences").
        Style(app.Theme().Primary()))
        
    // Add an interactive menu
    choice, _ := ui.NewMenu().
        Title("Choose an action").
        Options("Start", "Settings", "Exit").
        Show()
        
    // Handle the choice...
}
```

## Architecture

```
cmdux/
├── core/           # Core framework and rendering
├── ui/            # UI components (boxes, tables, menus)
├── ux/            # User experience (animations, effects)
├── input/         # Interactive components (prompts, forms)
└── style/         # Styling system (colors, layouts)
```

## Inspiration

cmdux combines the best ideas from:
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Elm architecture
- [tview](https://github.com/rivo/tview) - Rich widgets
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Styling

## License

MIT License