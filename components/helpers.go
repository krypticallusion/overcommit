package components

import "github.com/muesli/termenv"

var (
	ACCENT       = term.Color("#8AA8F9")
	SetTextStyle = termenv.String().Bold().Foreground(ACCENT).Styled
)
