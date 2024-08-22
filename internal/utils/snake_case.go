package utils

import (
	"os"
	"strings"
	"unicode"
)

type SeparatedCaseState int

const (
	SeparatedCaseStateStart SeparatedCaseState = iota
	SeparatedCaseStateUpper
	SeparatedCaseStateLower
	SeparatedCaseStateNewWord
)

func WantsSnakeCase() bool {
	return os.Getenv("OUTFLUX_JSON_SNAKE_CASE") == "true"
}

// Copy of https://github.com/JamesNK/Newtonsoft.Json/blob/4738a64817bb753667d9ed0ea99c1f955d414b33/Src/Newtonsoft.Json/Utilities/StringUtils.cs#L218
func ToSnakeCase(s string) string {
	if s == "" {
		return s
	}

	var sb strings.Builder
	state := SeparatedCaseStateStart
	separator := byte('_')

	for i, c := range s {
		if c == ' ' {
			if state != SeparatedCaseStateStart {
				state = SeparatedCaseStateNewWord
			}
		} else if unicode.IsUpper(c) {
			switch state {
			case SeparatedCaseStateUpper:
				if i > 0 && i+1 < len(s) {
					nextChar := s[i+1]
					if !unicode.IsUpper(rune(nextChar)) && nextChar != byte(separator) {
						sb.WriteByte(separator)
					}
				}
			case SeparatedCaseStateLower, SeparatedCaseStateNewWord:
				sb.WriteByte(separator)
			}

			sb.WriteByte(byte(unicode.ToLower(rune(c))))
			state = SeparatedCaseStateUpper
		} else if byte(c) == separator {
			sb.WriteByte(separator)
			state = SeparatedCaseStateStart
		} else {
			if state == SeparatedCaseStateNewWord {
				sb.WriteByte(separator)
			}

			sb.WriteByte(byte(c))
			state = SeparatedCaseStateLower
		}
	}

	return sb.String()
}
