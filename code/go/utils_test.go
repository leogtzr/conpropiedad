package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"testing"
	"time"
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
		if got := validateOptions(tt.options...); got != tt.want {
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
		if got := shouldIgnoreLine(tt.line); got != tt.want {
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
		if match := isWordFormatValid(tt.input, rgx); match != tt.match {
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
		if got := extractWordFromText(tt.wordText); !equalWords(got, tt.want) {
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
		if got := equalTags(tt.a, tt.b); got != tt.result {
			t.Errorf("got=[%t], should be [%t]", got, tt.result)
		}
	}
}

func Test_loadInputFile(t *testing.T) {

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	randFileName := stringWithCharset(5, charset, seededRand)
	tmpFilePath := filepath.Join("/tmp", randFileName)

	wordFormatRgx := *regexp.MustCompile(`^.+;.+;.+$`)

	_, err := loadInputFile(tmpFilePath, &wordFormatRgx)
	if err == nil {
		t.Errorf("'%s' does not exist, it should have failed to load the file.", tmpFilePath)
	}

	file, err := os.Create(tmpFilePath)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	fmt.Fprintln(w, "a;b;c,d,e")
	fmt.Fprintln(w, "# que diceeee")
	fmt.Fprintln(w, "que diceeee;")
	fmt.Fprintln(w, "hola;ok;bye")
	w.Flush()

	type test struct {
		want []Word
	}

	tests := []test{
		{
			want: []Word{
				Word{word: "a", meaning: "b", tags: []string{"c", "d", "e"}},
				Word{word: "hola", meaning: "ok", tags: []string{"bye"}},
			},
		},
	}

	for _, tt := range tests {
		words, err := loadInputFile(tmpFilePath, &wordFormatRgx)
		if err != nil {
			t.Error(err)
		}
		if len(words) != len(tt.want) {
			t.Errorf("got=[%d] words, want=[%d]", len(words), len(tt.want))
		}
		for i, w := range words {
			if !equalWords(w, tt.want[i]) {
				t.Errorf("got=[%s], want=[%s]", w, tt.want[i])
			}
		}
	}

	err = os.RemoveAll(tmpFilePath)
	if err != nil {
		t.Errorf("unexpedted error: [%s]", err)
	}
}

func Test_filterByText(t *testing.T) {
	words := []Word{
		Word{word: "w1", meaning: "m1", tags: []string{"hola", "bebe"}},
		Word{word: "w2", meaning: "m2", tags: []string{"oki", "doki"}},
		Word{word: "w3", meaning: "m3", tags: []string{"oki3", "doki3"}},
		Word{word: "w4", meaning: "m4", tags: []string{"facebook", "twitter"}},
	}

	type test struct {
		words     *[]Word
		tagSearch string
		want      []Word
	}

	tests := []test{
		{
			words:     &words,
			tagSearch: "a",
			want: []Word{
				Word{word: "w1", meaning: "m1", tags: []string{"hola", "bebe"}},
				Word{word: "w4", meaning: "m4", tags: []string{"facebook", "twitter"}},
			},
		},
		{
			words:     &words,
			tagSearch: "o",
			want: []Word{
				Word{word: "w1", meaning: "m1", tags: []string{"hola", "bebe"}},
				Word{word: "w2", meaning: "m2", tags: []string{"oki", "doki"}},
				Word{word: "w3", meaning: "m3", tags: []string{"oki3", "doki3"}},
				Word{word: "w4", meaning: "m4", tags: []string{"facebook", "twitter"}},
			},
		},
	}

	for _, tt := range tests {
		got := filterByText(tt.words, tt.tagSearch)
		if len(got) != len(tt.want) {
			t.Errorf("got=[%d], want=[%d]", len(got), len(tt.want))
		}
		for i, w := range got {
			if !equalWords(w, tt.want[i]) {
				t.Errorf("got=[%s], want=[%s]", w, tt.want[i])
			}
		}
	}
}

func TestWord_String(t *testing.T) {
	type test struct {
		word Word
		want string
	}

	tests := []test{
		{
			word: Word{word: "w", meaning: "m", tags: []string{"t1", "t2"}},
			want: "w: m		(t1, t2)",
		},
		{
			word: Word{word: "w2", meaning: "m2", tags: []string{"tg1", "tg2"}},
			want: "w2: m2		(tg1, tg2)",
		},
	}

	for _, tt := range tests {
		if got := tt.word.String(); got != tt.want {
			t.Errorf("got=[%s], want=[%s]", got, tt.want)
		}
	}
}
