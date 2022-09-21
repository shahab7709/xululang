package xulu

import (
	"errors"
	"strings"
)

var xuluVerbs = []string{"abcd", "bcde", "dede"}

type xuluVerb string

type xuluLangProvider struct {
	words []string
	index int
}

func NewXuluLang(input string) xuluLangProvider {
	words := strings.Split(input, xuluSeprator)
	return xuluLangProvider{words: words}
}

func (xl *xuluLangProvider) create() (ixuluSentence, error) {

	isVerbSentence := false
	if !xl.isVerb(xl.words[xl.index]) {
		return nil, errors.New("sentence dosnt had verb")
	}

	if len(xl.words) == xl.index+1 {
		return nil, errors.New("sentence dosnt had name")
	}

	if xl.isVerb(xl.words[xl.index+1]) {
		isVerbSentence = true
	}

	result := NewXuluSentence(xuluVerb(xl.words[xl.index]))
	xl.index++
	for xl.index < len(xl.words) {
		if xl.isVerb(xl.words[xl.index]) && isVerbSentence {
			sentence, err := xl.create()
			if err != nil {
				return nil, errors.New("not verb")
			}
			result.sentences = append(result.sentences, sentence)
			continue
		}

		if xl.isVerb(xl.words[xl.index]) && !isVerbSentence {
			if len(result.sentences) > 0 {
				return &result, nil
			}
			return nil, errors.New("wrong order")
		}

		if !isVerbSentence {
			sentence, err := NewXuluName(xl.words[xl.index])
			if err != nil {
				return nil, err
			}
			result.sentences = append(result.sentences, sentence)
		}
		xl.index++

	}
	return &result, nil
}

func (xl *xuluLangProvider) isVerb(input string) bool {
	for _, item := range xuluVerbs {
		if input == item {
			return true
		}
	}
	return false
}

func (xl *xuluLangProvider) isName(input string) error {
	for _, item := range xuluVerbs {
		if input == item {
			return nil
		}
	}
	return errors.New("not verb")
}
