package mathfunc

import "math"

func Round(x, unit float64) float64 {
	// Compile parses a regular expression and returns, if successful,
	// a Regexp that can be used to match against text.
	return math.Round(x/unit) * unit
}
