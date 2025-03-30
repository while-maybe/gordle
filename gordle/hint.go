package gordle

import "strings"

type hint byte

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

// String implements the Stringer interface.
func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "⬜️" // grey square
	case wrongPosition:
		return "🟡" // yellow circle
	case correctPosition:
		return "💚" // green heart
	default:
		// we shouldn't be here
		return "💔" // red broken heart
	}
}

// feedback is a list of hints, one per character of the word.
type feedback []hint

// StringConcat is a naive implementation to build feedback as a string.
// It is used only to benchmark it against the strings.Builder version.
func (fb feedback) StringConcat() string {
	var output string
	for _, hint := range fb {
		output += hint.String()
	}
	return output
}

// String implements the Stringer interface for a slice of hints.
func (fb feedback) String() string {
	sb := strings.Builder{}
	for _, hint := range fb {
		sb.WriteString(hint.String())
	}
	return sb.String()
}
