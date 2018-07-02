package caseutil

import (
	"strings"
	"unicode"
)

// Words holds a list of words.
type Words [][]rune

// SplitFunc tests a character and returns split=true if the character
// represents a word boundary.  If consume=true, the character will not be
// included in the output, otherwise it will be the first character in the next
// word.
type SplitFunc func(chars []rune, i int) (split bool, consume bool)

// FromDelimFunc splits a string into words using a function which identifies word
// boundaries.
func FromDelimFunc(text string, split SplitFunc) Words {
	chars := []rune(text)
	var words [][]rune
	start := 0

	for i := 0; i < len(chars); i++ {
		ok, consume := split(chars, i)
		if ok {
			words = append(words, chars[start:i])
			start = i
			if consume {
				start++
			}
		}
		chars[i] = unicode.ToLower(chars[i])
	}

	if start < len(chars) {
		words = append(words, chars[start:])
	}

	return Words(words)
}

// FromDelim splits text into words around the given delimiter.
func FromDelim(text string, delim rune) Words {
	return FromDelimFunc(text, func(chars []rune, i int) (bool, bool) {
		return chars[i] == delim, true
	})
}

// FromInitial creates a Words instance from text in initial or camel
// case.
func FromInitial(text string) Words {
	return FromDelimFunc(text, func(chars []rune, i int) (bool, bool) {
		return i != 0 && unicode.IsUpper(chars[i]), false
	})
}

// FromSnake creates a Words instance from text in snake case.
func FromSnake(text string) Words {
	return FromDelim(text, '_')
}

// FromKebab creates a Words instance from text in kebab case.
func FromKebab(text string) Words {
	return FromDelim(text, '-')
}

// ToInitial formats the words as initial case, optionally capitalising the
// first letter.
func (w Words) ToInitial(firstLetter bool) string {
	var text string

	for i, word := range w {
		if len(word) == 0 {
			continue
		}
		if firstLetter || i > 0 {
			if len(word) > 1 {
				text += string(unicode.ToUpper(word[0])) + string(word[1:])
			} else {
				text += string(unicode.ToUpper(word[0]))
			}
		} else {
			text += string(word)
		}
	}

	return text
}

// ToDelim joins the words with the given delimeter, optionally uppercasing the
// words.
func (w Words) ToDelim(delim rune, uppercase bool) string {
	var text string

	for i, word := range w {
		wordStr := string(word)

		if uppercase {
			wordStr = strings.ToUpper(wordStr)
		}

		if i != 0 {
			text += string(delim)
		}
		text += wordStr
	}

	return text
}

// ToSnake makes the words into snake case, optionally screaming.
func (w Words) ToSnake(scream bool) string {
	return w.ToDelim('_', scream)
}

// ToKebab makes the words into kebab case, optionally fat.
func (w Words) ToKebab(fat bool) string {
	return w.ToDelim('-', fat)
}
