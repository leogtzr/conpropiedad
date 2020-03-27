package main

import (
	"regexp"
	"testing"
)

func Test_validateOptions(t *testing.T) {
	type test struct {
		options []string
		want    bool
	}

	tests := []test{
		{
			options: []string{"", "t"},
			want:    false,
		},
		{
			options: []string{"tags", "t"},
			want:    true,
		},
		{
			options: []string{},
			want:    false,
		},
	}

	for _, tt := range tests {
		got := validateOptions(tt.options...)
		if got != tt.want {
			t.Errorf("got=[%t], want=[%t]", got, tt.want)
		}

	}
}

func Test_shouldIgnoreLine(t *testing.T) {
	type test struct {
		line string
		want bool
	}

	tests := []test{
		{line: "hola", want: false},
		{line: "# comment", want: true},
		{line: "", want: true},
	}

	for _, tt := range tests {
		got := shouldIgnoreLine(tt.line)
		if got != tt.want {
			t.Errorf("got=[%t], want=[%t]", got, tt.want)
		}
	}
}

func Test_isWordFormatValid(t *testing.T) {
	rgx := regexp.MustCompile(`^.+;.+;.+$`)
	type test struct {
		input string
		match bool
	}

	tests := []test{
		{input: "Putón;Mujer promiscua.;volada,wila,promiscua", match: true},
		{input: "Querida;Amante.;amante,canco", match: true},
		{input: "Hola;Abc;", match: false},
		{input: "Hola;", match: false},
		{input: "", match: false},
	}
	for _, tt := range tests {
		match := isWordFormatValid(tt.input, rgx)
		if match != tt.match {
			t.Errorf("got=[%t], want=[%t] for '%s'", match, tt.match, tt.input)
		}
	}
}

func Test_extractWordFromText(t *testing.T) {
	type test struct {
		wordText string
		want     Word
	}

	tests := []test{
		{
			wordText: "a;b;c",
			want:     Word{word: "a", meaning: "b", tags: []string{"c"}},
		},
		{
			wordText: "a;b;c,d",
			want:     Word{word: "a", meaning: "b", tags: []string{"c", "d"}},
		},
	}

	for _, tt := range tests {
		got := extractWordFromText(tt.wordText)
		if !equalWords(got, tt.want) {
			t.Errorf("got=[%s], want=[%s]", got, tt.want)
		}
	}

}

func TestEqualTags(t *testing.T) {
	type test struct {
		a      []string
		b      []string
		result bool
	}

	tests := []test{
		{a: []string{"java", "linux"}, b: []string{"linux", "java"}, result: false},
		{a: []string{"java", "linux"}, b: []string{"java", "linux"}, result: true},
		{a: []string{"java"}, b: []string{"java", "linux"}, result: false},
	}

	for _, tt := range tests {
		got := equalTags(tt.a, tt.b)
		if got != tt.result {
			t.Errorf("got=[%t], should be [%t]", got, tt.result)
		}
	}
}
