package cli

import (
	"errors"
	"strings"
)

func IsQuote(c rune) bool {
	return c == '"' || c == '\''
}

func BuildCommand(txt string) ([]string, error) {
	if txt == "" {
		return nil, errors.New("err: Invalid comand string")
	}

	res := make([]string, 0, 3)

	bld := new(strings.Builder)
	inQuote := false
	var openQuote rune
	for _, t := range txt {
		if IsQuote(t) {
			// open quote
			if !inQuote {
				openQuote = t
				inQuote = true
				continue
			}

			// close quote
			if inQuote && openQuote == t {
				inQuote = false
				continue
			}
		}

		if t == ' ' && !inQuote {
			res = append(res, bld.String())

			bld.Reset()
			continue
		}

		bld.WriteRune(t)
	}

	// put the rest, if any
	if bld.Len() > 0 {
		res = append(res, bld.String())
	}

	return res, nil
}
