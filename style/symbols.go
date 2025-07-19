// Package style provides symbol definitions for drawing UI elements.
package style

// Box drawing characters for modern terminals
const (
	BoxTopLeft     = "╭"
	BoxTopRight    = "╮"
	BoxBottomLeft  = "╰"
	BoxBottomRight = "╯"
	BoxHorizontal  = "─"
	BoxVertical    = "│"
	BoxTee         = "├"
	BoxCross       = "┼"
	BoxElbow       = "└"
	BoxTeeRight    = "┤"
	BoxTeeTop      = "┬"
	BoxTeeBottom   = "┴"
)

// Modern bullets and separators
const (
	Bullet    = "●"
	Arrow     = "▸"
	CheckMark = "✓"
	CrossMark = "✗"
	Lightning = "⚡"
	Gear      = "⚙"
	Rocket    = "🚀"
	Diamond   = "◆"
	Circle    = "●"
	Star      = "★"
	Heart     = "♥"
	Fire      = "🔥"
	Target    = "🎯"
	Trophy    = "🏆"
)

// Progress and loading symbols
const (
	ProgressFull  = "█"
	ProgressEmpty = "░"
	ProgressLeft  = "▌"
	ProgressRight = "▐"
)

// Spacing constants
const (
	Indent       = "  "
	DoubleIndent = "    "
	TripleIndent = "      "
)

// Classic ASCII alternatives
const (
	ClassicBoxTopLeft     = "+"
	ClassicBoxTopRight    = "+"
	ClassicBoxBottomLeft  = "+"
	ClassicBoxBottomRight = "+"
	ClassicBoxHorizontal  = "-"
	ClassicBoxVertical    = "|"
	ClassicBullet         = "*"
	ClassicArrow          = ">"
	ClassicCheckMark      = "v"
	ClassicCrossMark      = "x"
)

// SymbolSet represents a set of symbols for drawing UI elements.
type SymbolSet struct {
	// Box drawing
	BoxTopLeft     string
	BoxTopRight    string
	BoxBottomLeft  string
	BoxBottomRight string
	BoxHorizontal  string
	BoxVertical    string
	BoxTee         string
	BoxCross       string
	
	// UI elements
	Bullet     string
	Arrow      string
	CheckMark  string
	CrossMark  string
	Selected   string
	Unselected string
}

// DefaultSymbols returns the default Unicode symbol set.
func DefaultSymbols() SymbolSet {
	return SymbolSet{
		BoxTopLeft:     BoxTopLeft,
		BoxTopRight:    BoxTopRight,
		BoxBottomLeft:  BoxBottomLeft,
		BoxBottomRight: BoxBottomRight,
		BoxHorizontal:  BoxHorizontal,
		BoxVertical:    BoxVertical,
		BoxTee:         BoxTee,
		BoxCross:       BoxCross,
		
		Bullet:     Bullet,
		Arrow:      Arrow,
		CheckMark:  CheckMark,
		CrossMark:  CrossMark,
		Selected:   "▶",
		Unselected: " ",
	}
}

// ASCIISymbols returns ASCII-only symbols for compatibility.
func ASCIISymbols() SymbolSet {
	return SymbolSet{
		BoxTopLeft:     ClassicBoxTopLeft,
		BoxTopRight:    ClassicBoxTopRight,
		BoxBottomLeft:  ClassicBoxBottomLeft,
		BoxBottomRight: ClassicBoxBottomRight,
		BoxHorizontal:  ClassicBoxHorizontal,
		BoxVertical:    ClassicBoxVertical,
		BoxTee:         "+",
		BoxCross:       "+",
		
		Bullet:     ClassicBullet,
		Arrow:      ClassicArrow,
		CheckMark:  ClassicCheckMark,
		CrossMark:  ClassicCrossMark,
		Selected:   ">",
		Unselected: " ",
	}
}