package gordle

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in English": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 characters in Arabic": {
			input: "مرحبا",
			want:  []rune("مرحبا"),
		},
		"5 characters in Japanese": {
			input: "こんにちは",
			want:  []rune("こんにちは"),
		},
		"3 characters in Japanese": {
			input: "こんに\nこんにちは",
			want:  []rune("こんにちは"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(strings.NewReader(tc.input), string(tc.want), 0)

			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

func TestGameValidateGuess(t *testing.T) {
	tt := map[string]struct {
		word []rune
		want error
	}{
		"nominal": {
			word: []rune("GUESS"),
			want: nil,
		},
		"too short": {
			word: []rune("HI"),
			want: errInvalidWordLength,
		},
		"too long": {
			word: []rune("SHOULDFAIL"),
			want: errInvalidWordLength,
		},
		"empty string": {
			word: []rune(""),
			want: errInvalidWordLength,
		},
		"empty slice": {
			word: []rune{},
			want: errInvalidWordLength,
		},
		"nil": {
			word: nil,
			want: errInvalidWordLength,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			// only the validateGuess is being tested so New take nil/0 args
			g := New(nil, "XXXXX", 0)

			err := g.validateGuess(tc.word)
			if !errors.Is(err, tc.want) {
				t.Errorf("%c, expected %q, got %q", tc.word, tc.want, err)
			}
		})
	}
}

func TestGameSplitToUppercaseCharacters(t *testing.T) {
	tt := map[string]struct {
		word string
		want []rune
	}{
		"lower": {
			word: "lower",
			want: []rune{'L', 'O', 'W', 'E', 'R'},
		},
		"Title": {
			word: "Title",
			want: []rune{'T', 'I', 'T', 'L', 'E'},
		},
		"mIxEd": {
			word: "mIxEd",
			want: []rune{'M', 'I', 'X', 'E', 'D'},
		},
		"CAPITALS": {
			word: "CAPITALS",
			want: []rune{'C', 'A', 'P', 'I', 'T', 'A', 'L', 'S'},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := splitToUppercaseCharacters(tc.word)

			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}
